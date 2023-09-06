// Package defcon
// means DEFault CONtainers
package defcon

import (
	"reflect"

	"github.com/soverenio/vanilla/throw"
	"github.com/soverenio/vanilla/zero"
	"golang.org/x/exp/slices"
)

type ListIterator[E any] interface {
	Map(func(E) E) ListIterator[E]
	MapIndex(func(int, E) E) ListIterator[E]
	Filter(func(E) bool) ListIterator[E]
	FilterIndex(func(int, E) bool) ListIterator[E]
	Sort(less func(E, E) bool) ListIterator[E]
	Unique(func(E) string) ListIterator[E]

	Skip(count int) ListIterator[E]
	Limit(count int) ListIterator[E]

	Count() int

	Apply(func(E)) int
	ApplyIndex(func(int, E))
	Finalize() []E
	Intermediate() ListIterator[E]
}

type ListGenerator[K any] func() (K, bool)

var _ ListIterator[any] = &listIterator[any]{}

type listIterator[E any] struct {
	generator   ListGenerator[E]
	maxElements int
}

func NewListIterator[E any](gen ListGenerator[E], maxElements int) ListIterator[E] {
	return &listIterator[E]{
		generator:   gen,
		maxElements: maxElements,
	}
}

func (l *listIterator[E]) Map(f func(E) E) ListIterator[E] {
	return &listIterator[E]{
		generator: func() (E, bool) {
			val, ok := l.generator()
			if !ok {
				return zero.Zero[E](), false
			}
			return f(val), true
		},
		maxElements: l.maxElements,
	}
}

func (l *listIterator[E]) MapIndex(f func(int, E) E) ListIterator[E] {
	pos := 0

	return &listIterator[E]{
		generator: func() (E, bool) {
			val, ok := l.generator()
			if !ok {
				return zero.Zero[E](), false
			}
			rv := f(pos, val)
			pos++
			return rv, true
		},
		maxElements: l.maxElements,
	}
}

// Filter returns a new iterator that will only return elements for which the given function returns true.
func (l *listIterator[E]) Filter(f func(E) bool) ListIterator[E] {
	return &listIterator[E]{
		generator: func() (E, bool) {
			for {
				val, ok := l.generator()
				if !ok {
					return zero.Zero[E](), false
				}
				if f(val) {
					return val, true
				}
			}
		},
		maxElements: l.maxElements,
	}
}

func (l *listIterator[E]) FilterIndex(f func(int, E) bool) ListIterator[E] {
	pos := 0

	return &listIterator[E]{
		generator: func() (E, bool) {
			for {
				val, ok := l.generator()
				if !ok {
					return zero.Zero[E](), false
				}
				use := f(pos, val)
				pos++
				if use {
					return val, true
				}
			}
		},
		maxElements: l.maxElements,
	}
}

func (l *listIterator[E]) Sort(less func(E, E) bool) ListIterator[E] {
	var sorted []E
	if l.maxElements > 0 {
		sorted = make([]E, 0, l.maxElements)
	}

	for {
		e, ok := l.generator()
		if !ok {
			break
		}
		sorted = append(sorted, e)
	}

	slices.SortFunc(sorted, less)
	return IteratorFromList[E](sorted)
}

func (l *listIterator[E]) Unique(fn func(E) string) ListIterator[E] {
	uniqueChecker := make(map[string]struct{}, l.maxElements)

	return &listIterator[E]{
		generator: func() (E, bool) {
			for {
				e, ok := l.generator()
				if !ok {
					return zero.Zero[E](), false
				}
				uniqueValue := fn(e)
				if _, ok := uniqueChecker[uniqueValue]; ok {
					continue
				}
				uniqueChecker[uniqueValue] = struct{}{}
				return e, true
			}
		},
		maxElements: l.maxElements,
	}
}

func (l *listIterator[E]) Apply(f func(E)) int {
	count := 0

	for {
		e, ok := l.generator()
		if !ok {
			break
		}
		f(e)
		count++
	}

	return count
}

func (l *listIterator[E]) ApplyIndex(f func(int, E)) {
	for pos := 0; ; pos++ {
		e, ok := l.generator()
		if !ok {
			break
		}
		f(pos, e)
	}
}

func (l *listIterator[E]) Finalize() []E {
	var rv []E
	if l.maxElements > 0 {
		rv = make([]E, 0, l.maxElements)
	}

	for {
		e, ok := l.generator()
		if !ok {
			break
		}
		rv = append(rv, e)
	}

	return rv[:len(rv):len(rv)]
}

func (l *listIterator[E]) Intermediate() ListIterator[E] {
	return IteratorFromList(l.Finalize())
}

func (l *listIterator[E]) Skip(count int) ListIterator[E] {
	if count < 0 {
		panic(throw.IllegalValue())
	}

	return &listIterator[E]{
		generator: func() (E, bool) {
			for {
				rv, ok := l.generator()
				if !ok {
					return zero.Zero[E](), false
				} else if count > 0 {
					count--
					continue
				}
				return rv, true
			}
		},
		maxElements: l.maxElements - count,
	}
}

func (l *listIterator[E]) Count() int {
	count := 0

	for {
		_, ok := l.generator()
		if !ok {
			break
		}
		count++
	}

	return count
}

func (l *listIterator[E]) Limit(count int) ListIterator[E] {
	if count < 0 {
		panic(throw.IllegalValue())
	}

	maxElements := count
	if l.maxElements >= 0 && maxElements > l.maxElements {
		maxElements = l.maxElements
	}

	return &listIterator[E]{
		generator: func() (E, bool) {
			rv, ok := l.generator()
			if !ok || count <= 0 {
				return zero.Zero[E](), false
			}
			count--
			return rv, true
		},
		maxElements: maxElements,
	}
}

func IteratorFromList[E any](inp []E) ListIterator[E] {
	pos := 0

	return &listIterator[E]{
		generator: func() (E, bool) {
			if pos >= len(inp) {
				return zero.Zero[E](), false
			}
			elem := inp[pos]
			pos++
			return elem, true
		},
		maxElements: len(inp),
	}
}

func IteratorFromListWithCallback[E any, I any](inp []E, cb func(E) I) ListIterator[I] {
	pos := 0

	return &listIterator[I]{
		generator: func() (I, bool) {
			if pos >= len(inp) {
				return zero.Zero[I](), false
			}
			elem := inp[pos]
			pos++
			return cb(elem), true
		},
		maxElements: len(inp),
	}
}

func IteratorFromMapKeys[K comparable, V any](inp map[K]V) ListIterator[K] {
	iter := reflect.ValueOf(inp).MapRange()

	return &listIterator[K]{
		generator: func() (K, bool) {
			if !iter.Next() {
				return zero.Zero[K](), false
			}
			return iter.Key().Interface().(K), true
		},
		maxElements: len(inp),
	}
}

func IteratorFromMapValues[K comparable, V any](inp map[K]V) ListIterator[V] {
	iter := reflect.ValueOf(inp).MapRange()

	return &listIterator[V]{
		generator: func() (V, bool) {
			if !iter.Next() {
				return zero.Zero[V](), false
			}
			return iter.Value().Interface().(V), true
		},
		maxElements: len(inp),
	}
}

func ListIteratorFromGenerator[E any](generator func() (E, bool), count int) ListIterator[E] {
	return &listIterator[E]{
		generator:   generator,
		maxElements: count,
	}
}
