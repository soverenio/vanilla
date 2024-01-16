package defcon

import (
	"reflect"

	"golang.org/x/exp/slices"

	"github.com/soverenio/vanilla/throw"
	"github.com/soverenio/vanilla/zero"
)

type MapIterator[K comparable, V any] interface {
	MapV(func(K, V) V) MapIterator[K, V]
	MapKV(func(K, V) (K, V)) MapIterator[K, V]
	Filter(func(K, V) bool) MapIterator[K, V]
	Unique(func(K, V) string) MapIterator[K, V]

	MapVIndex(func(int, K, V) V) MapIterator[K, V]
	MapKVIndex(func(int, K, V) (K, V)) MapIterator[K, V]
	FilterIndex(func(int, K, V) bool) MapIterator[K, V]

	Apply(func(K, V))
	Finalize() map[K]V
	Intermediate(less func(K, K) bool) MapIterator[K, V]

	Skip(count int) MapIterator[K, V]
	Limit(count int) MapIterator[K, V]

	KeyIterator() ListIterator[K]
	ValueIterator() ListIterator[V]
}

type MapGenerator[K comparable, V any] func() (K, V, bool)

var _ MapIterator[int, any] = &mapIterator[int, any]{}

type mapIterator[K comparable, V any] struct {
	generator   MapGenerator[K, V]
	maxElements int
}

func NewMapIterator[K comparable, V any](gen MapGenerator[K, V], maxElements int) MapIterator[K, V] {
	return &mapIterator[K, V]{
		generator:   gen,
		maxElements: maxElements,
	}
}

func (m *mapIterator[K, V]) MapV(f func(K, V) V) MapIterator[K, V] {
	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			key, val, ok := m.generator()
			if !ok {
				return zero.Zero[K](), zero.Zero[V](), false
			}
			return key, f(key, val), true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) MapKV(f func(K, V) (K, V)) MapIterator[K, V] {
	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			key, val, ok := m.generator()
			if !ok {
				return zero.Zero[K](), zero.Zero[V](), false
			}
			key, val = f(key, val)
			return key, val, true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) Filter(f func(K, V) bool) MapIterator[K, V] {
	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			for {
				key, val, ok := m.generator()
				if !ok {
					return zero.Zero[K](), zero.Zero[V](), false
				} else if f(key, val) {
					return key, val, true
				}
			}
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) MapVIndex(f func(int, K, V) V) MapIterator[K, V] {
	pos := 0

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			key, val, ok := m.generator()
			if !ok {
				return zero.Zero[K](), zero.Zero[V](), false
			}
			newVal := f(pos, key, val)
			pos++
			return key, newVal, true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) MapKVIndex(f func(int, K, V) (K, V)) MapIterator[K, V] {
	pos := 0

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			key, val, ok := m.generator()
			if !ok {
				return zero.Zero[K](), zero.Zero[V](), false
			}
			key, val = f(pos, key, val)
			pos++
			return key, val, true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) FilterIndex(f func(int, K, V) bool) MapIterator[K, V] {
	pos := 0

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			for {
				key, val, ok := m.generator()
				if !ok {
					return zero.Zero[K](), zero.Zero[V](), false
				}

				take := f(pos, key, val)
				pos++
				if take {
					return key, val, true
				}
			}
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) Unique(fn func(K, V) string) MapIterator[K, V] {
	uniqueChecker := make(map[string]struct{}, m.maxElements)

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			for {
				k, v, ok := m.generator()
				if !ok {
					return zero.Zero[K](), zero.Zero[V](), false
				}
				uniqueValue := fn(k, v)
				if _, ok := uniqueChecker[uniqueValue]; ok {
					continue
				}
				uniqueChecker[uniqueValue] = struct{}{}
				return k, v, true
			}
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) Apply(f func(K, V)) {
	for {
		k, v, ok := m.generator()
		if !ok {
			break
		}
		f(k, v)
	}
}

func (m *mapIterator[K, V]) Finalize() map[K]V {
	rv := make(map[K]V, m.maxElements)
	for {
		k, v, ok := m.generator()
		if !ok {
			break
		}
		rv[k] = v
	}
	return rv
}

func (m *mapIterator[K, V]) Intermediate(less func(K, K) bool) MapIterator[K, V] {
	return IteratorFromMap(m.Finalize(), less)
}

func (m *mapIterator[K, V]) KeyIterator() ListIterator[K] {
	return &listIterator[K]{
		generator: func() (K, bool) {
			k, _, ok := m.generator()
			if !ok {
				return zero.Zero[K](), false
			}
			return k, true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) ValueIterator() ListIterator[V] {
	return &listIterator[V]{
		generator: func() (V, bool) {
			_, v, ok := m.generator()
			if !ok {
				return zero.Zero[V](), false
			}
			return v, true
		},
		maxElements: m.maxElements,
	}
}

func (m *mapIterator[K, V]) Skip(count int) MapIterator[K, V] {
	if count < 0 {
		panic(throw.IllegalValue())
	}

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			for {
				k, v, ok := m.generator()
				if !ok {
					return zero.Zero[K](), zero.Zero[V](), false
				} else if count > 0 {
					count--
					continue
				}
				return k, v, true
			}
		},
		maxElements: m.maxElements - count,
	}
}

func (m *mapIterator[K, V]) Limit(count int) MapIterator[K, V] {
	if count < 0 {
		panic(throw.IllegalValue())
	}

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			for {
				k, v, ok := m.generator()
				if !ok || count <= 0 {
					return zero.Zero[K](), zero.Zero[V](), false
				}
				count--
				return k, v, true //nolint:staticcheck // SA4004: the surrounding loop is unconditionally terminated
			}
		},
		maxElements: m.maxElements - count,
	}
}

func IteratorFromMap[K comparable, V any](inp map[K]V, less func(K, K) bool) MapIterator[K, V] {
	if less != nil {
		keys := make([]K, 0, len(inp))
		for key := range inp {
			keys = append(keys, key)
		}
		
		slices.SortFunc(keys, func(a, b K) int {
			switch {
			case less(a, b):
				return -1
			case less(b, a):
				return 1
			default:
				return 0
			}
		})

		pos := 0

		return &mapIterator[K, V]{
			generator: func() (K, V, bool) {
				if pos >= len(keys) {
					return zero.Zero[K](), zero.Zero[V](), false
				}

				key, val := keys[pos], inp[keys[pos]]
				pos++
				return key, val, true
			},
			maxElements: len(keys),
		}
	}

	iter := reflect.ValueOf(inp).MapRange()

	return &mapIterator[K, V]{
		generator: func() (K, V, bool) {
			if !iter.Next() {
				return zero.Zero[K](), zero.Zero[V](), false
			}

			return iter.Key().Interface().(K), iter.Value().Interface().(V), true
		},
		maxElements: len(inp),
	}
}

func MapIteratorFromGenerator[K comparable, V any](generator func() (K, V, bool), count int) MapIterator[K, V] {
	return &mapIterator[K, V]{
		generator:   generator,
		maxElements: count,
	}
}
