package promise

import (
	"context"

	"github.com/soverenio/vanilla/atomickit"
	"github.com/soverenio/vanilla/throw"
	"github.com/soverenio/vanilla/zero"
)

var (
	simpleCounter atomickit.Int
)

func SimpleCounter() int {
	return simpleCounter.Load()
}

func NewZeroValueSimpleFunc[T any](inp SimplerFunc[T]) SimpleFunc[T] {
	return func(ctx context.Context) (T, error) { return zero.Zero[T](), inp(ctx) }
}

type SimplerFunc[T any] func(context.Context) error
type SimpleFunc[T any] func(context.Context) (T, error)

type Simple[T any] struct {
	executor SimpleFunc[T]
	await    chan struct{}

	val    T
	err    error
	cancel context.CancelFunc
}

// NewSimpleStarted creates a new promise with the given executor.
// NB: not thread-safe
func NewSimpleStarted[T any](ctx context.Context, executor SimpleFunc[T]) *Simple[T] {
	ctx, cancelFn := context.WithCancel(ctx)
	return (&Simple[T]{
		executor: executor,
		cancel:   cancelFn,
	}).executeAsync(ctx)
}

// ExecuteAsync executes the promise asynchronously.
func (p *Simple[T]) executeAsync(ctx context.Context) *Simple[T] {
	if p.await != nil {
		panic(throw.IllegalState())
	}
	p.await = make(chan struct{})

	simpleCounter.Add(1)

	go func() {
		defer simpleCounter.Add(-1)
		defer close(p.await)

		p.val, p.err = p.executor(ctx)
	}()

	return p
}

func (p *Simple[T]) IsStopped() bool {
	if p.await == nil {
		panic(throw.IllegalState())
	}

	select {
	case <-p.await:
		return true
	default:
		return false
	}
}

func (p *Simple[T]) C() <-chan struct{} {
	return p.await
}

func (p *Simple[T]) Poll() (T, bool, error) {
	if p.await == nil {
		panic(throw.IllegalState())
	} else if !p.IsStopped() {
		return p.val, false, p.err
	}

	return p.val, true, p.err
}

func (p *Simple[T]) Wait(ctx context.Context) (T, bool, error) {
	if p.await == nil {
		panic(throw.IllegalState())
	}

	select {
	case <-ctx.Done():
		return p.val, false, ctx.Err()
	case <-p.await:
		return p.val, true, p.err
	}
}

func (p *Simple[T]) Cancel() {
	if p.await == nil {
		panic(throw.IllegalState())
	}

	p.cancel()
}

func (p *Simple[T]) CancelAndWait(ctx context.Context) (T, bool, error) {
	p.Cancel()
	return p.Wait(ctx)
}
