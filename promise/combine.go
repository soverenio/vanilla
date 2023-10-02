package promise

import (
	"context"
	"sync"

	"github.com/soverenio/vanilla/throw"
)

var (
	ErrNotStarted = throw.New("not started")
)

type contextIDKey struct{}

func WithContextID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, contextIDKey{}, id)
}

func ContextID(ctx context.Context) int {
	return ctx.Value(contextIDKey{}).(int)
}

type SimpleGroup[T any] struct {
	executors []SimpleFunc[T]
	cancelFn  context.CancelFunc

	groupLock sync.Mutex
	group     []*Simple[T]
}

func NewSimpleGroup[T any](_ context.Context) *SimpleGroup[T] {
	return &SimpleGroup[T]{}
}

func NewSimpleGroupFixed[T any](ctx context.Context, executors ...SimpleFunc[T]) *SimpleGroup[T] {
	g := NewSimpleGroup[T](ctx)
	g.executors = executors
	return g
}

func NewSimpleGroupFixedRepeated[T any](ctx context.Context, executor SimpleFunc[T], count int) *SimpleGroup[T] {
	g := NewSimpleGroup[T](ctx)
	for i := 0; i < count; i++ {
		g.executors = append(g.executors, executor)
	}
	return g
}

func (g *SimpleGroup[T]) AddExecutor(executor SimpleFunc[T]) {
	g.groupLock.Lock()
	defer g.groupLock.Unlock()

	if g.group != nil {
		panic(throw.IllegalState())
	}

	g.executors = append(g.executors, executor)
}

func (g *SimpleGroup[T]) Start(ctx context.Context) error {
	g.groupLock.Lock()
	defer g.groupLock.Unlock()

	if len(g.group) != 0 {
		panic(throw.IllegalState()) // double start is not permitted
	}

	ctx, g.cancelFn = context.WithCancel(ctx)
	for id, executor := range g.executors {
		g.group = append(g.group, NewSimpleStarted[T](WithContextID(ctx, id), executor))
	}
	return nil
}

func (g *SimpleGroup[T]) Any(ctx context.Context) error {
	var out []<-chan struct{}

	{
		g.groupLock.Lock()

		if len(g.group) == 0 {
			g.groupLock.Unlock()
			return ErrNotStarted
		}

		out = make([]<-chan struct{}, 0, len(g.group))
		for _, simple := range g.group {
			out = append(out, simple.C())
		}

		g.groupLock.Unlock()
	}

	if pos, _, _ := SelectSlice(ctx, out); pos == -1 {
		return ctx.Err()
	}
	return nil
}

func (g *SimpleGroup[T]) All(ctx context.Context) error {
	{
		g.groupLock.Lock()
		if g.group == nil {
			g.groupLock.Unlock()
			return ErrNotStarted
		}
		g.groupLock.Unlock()
	}

	for _, simple := range g.group {
		_, ok, err := simple.Wait(ctx)
		if !ok {
			return err
		}
	}

	return nil
}

func (g *SimpleGroup[T]) FirstError() error {
	g.groupLock.Lock()
	defer g.groupLock.Unlock()

	for _, simple := range g.group {
		_, ok, err := simple.Poll()

		if ok && err != nil {
			return err
		}
	}
	return nil
}

func (g *SimpleGroup[T]) Stop(ctx context.Context) error {
	g.cancelFn()
	return g.All(ctx)
}
