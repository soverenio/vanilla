package set

import (
	"github.com/soverenio/vanilla/defcon"
	"github.com/soverenio/vanilla/zero"
)

type Set[T comparable] struct {
	internal map[T]struct{}
}

func NewSet[T comparable](expectedSize int) Set[T] {
	return Set[T]{internal: make(map[T]struct{}, expectedSize)}
}

func NewSetFromList[T comparable](inp []T) Set[T] {
	return NewSetFromListIterator(defcon.IteratorFromList(inp))
}

func NewSetFromListFunc[I any, T comparable](inp []I, f func(I) T) Set[T] {
	internal := defcon.ConvertListToMapIterator(defcon.IteratorFromList(inp),
		func(iv I) (T, struct{}) { return f(iv), struct{}{} },
	).Finalize()

	return Set[T]{internal: internal}
}

func NewSetFromListIterator[T comparable](inp defcon.ListIterator[T]) Set[T] {
	internal := defcon.ConvertListToMapIterator(inp, func(iv T) (T, struct{}) { return iv, struct{}{} }).Finalize()
	return Set[T]{internal: internal}
}

func NewSetFromMapKeys[K comparable, V any](inp map[K]V) Set[K] {
	keyIter := defcon.IteratorFromMapKeys(inp)
	internal := defcon.ConvertListToMapIterator(keyIter, func(iv K) (K, struct{}) { return iv, struct{}{} })
	return Set[K]{internal: internal.Finalize()}
}

func (s *Set[T]) Add(key T) bool {
	if s.Has(key) {
		return false
	}
	s.internal[key] = struct{}{}
	return true
}

func (s Set[T]) Len() int {
	return len(s.internal)
}

func (s Set[T]) Has(key T) bool {
	return defcon.MapHasKey(s.internal, key)
}

func (s Set[T]) Iterator() defcon.ListIterator[T] {
	return defcon.IteratorFromMapKeys(s.internal)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T](s.Len() + other.Len())
	s.Iterator().Apply(func(key T) { result.Add(key) })
	other.Iterator().Apply(func(key T) { result.Add(key) })
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T](s.Len())
	s.Iterator().Apply(func(key T) {
		if !other.Has(key) {
			result.Add(key)
		}
	})
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T](s.Len())
	s.Iterator().Apply(func(key T) {
		if other.Has(key) {
			result.Add(key)
		}
	})
	return result
}

func (s Set[T]) Take() T {
	for k := range s.internal {
		return k
	}
	return zero.Zero[T]()
}

func (s Set[T]) TakeN(count int) []T {
	return s.Iterator().Limit(count).Finalize()
}

func (s Set[T]) Delete(key T) bool {
	if !s.Has(key) {
		return false
	}
	delete(s.internal, key)
	return true
}

func FullIntersect[T comparable](left, right Set[T]) (outerLeft Set[T], inner Set[T], outerRight Set[T]) {
	outerLeft = NewSet[T](0)
	inner = NewSet[T](0)
	outerRight = NewSet[T](0)

	left.Iterator().Apply(func(key T) {
		if right.Has(key) {
			inner.Add(key)
		} else {
			outerLeft.Add(key)
		}
	})

	right.Iterator().Apply(func(key T) {
		if !left.Has(key) {
			outerRight.Add(key)
		}
	})

	return outerLeft, inner, outerRight
}
