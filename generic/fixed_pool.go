package generic

import (
	"sync"

	"github.com/soverenio/vanilla/throw"
)

type FixedSizePoolElementWrapped[T any] struct {
	elem     T
	poolLock sync.Mutex
	pool     *FixedSizePool[T]
}

func (e FixedSizePoolElementWrapped[T]) Unwrap() T { // nolint
	e.poolLock.Lock()
	defer e.poolLock.Unlock()

	if e.pool == nil {
		panic(throw.W(throw.IllegalState(), "trying to use released element"))
	}

	return e.elem
}

func (e FixedSizePoolElementWrapped[T]) Release() { // nolint
	e.poolLock.Lock()
	defer e.poolLock.Unlock()

	if e.pool == nil {
		panic(throw.W(throw.IllegalState(), "trying to release already released element"))
	}

	e.pool.Release(e.elem)
}

// FixedSizePool is offensive concurrent-safe pool with preallocated count of elements
// it throws exceptions in case of:
//   - element leakage (if buffer is drained)
//   - trying to use more elements that were allocated (if buffer is drained)
//   - double frees (if buffer is overflowed)
type FixedSizePool[T any] struct {
	pool       chan T
	maxSize    int
	newElement func() T
}

func createFixedSizePool[T any](new func() T, maxSize int) chan T {
	buffer := make(chan T, maxSize)
	for i := 0; i < maxSize; i++ {
		select {
		case buffer <- new():
		default:
			panic(throw.Impossible())
		}
	}
	return buffer
}

func NewFixedSizePool[T any](new func() T, maxSize int) FixedSizePool[T] {
	return FixedSizePool[T]{
		pool:       createFixedSizePool(new, maxSize),
		maxSize:    maxSize,
		newElement: new,
	}
}

func (f *FixedSizePool[T]) Get() T {
	select {
	case elem := <-f.pool:
		return elem
	default:
		panic(throw.IllegalState())
	}
}

func (f *FixedSizePool[T]) Release(elem T) {
	select {
	case f.pool <- elem:
	default:
		panic(throw.IllegalState())
	}
}

func (f *FixedSizePool[T]) GetWrapped() FixedSizePoolElementWrapped[T] {
	return FixedSizePoolElementWrapped[T]{
		elem: f.Get(),
		pool: f,
	}
}
