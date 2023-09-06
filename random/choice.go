package random

import (
	"github.com/soverenio/vanilla/defcon"
	"github.com/soverenio/vanilla/defcon/set"
	"github.com/soverenio/vanilla/throw"
	"github.com/soverenio/vanilla/zero"
)

func Choice[T any](inp []T) T {
	if len(inp) == 0 {
		panic(throw.IllegalValue())
	}
	return inp[SignedMathRandBetween[int](0, len(inp)-1)]
}

func MapChoice[K comparable, V any](inp map[K]V) (K, V, bool) {
	if len(inp) == 0 {
		return zero.Zero[K](), zero.Zero[V](), false
	}

	pos := SignedMathRandBetween[int](0, len(inp)-1)

	for k, v := range inp {
		if pos == 0 {
			return k, v, true
		}
		pos--
	}

	panic(throw.Impossible())
}

func ListChoicesIter[T any](inp []T, count int, unsorted bool) defcon.ListIterator[T] {
	switch {
	case count < 0:
		panic(throw.IllegalValue())
	case count > len(inp):
		count = len(inp)
	}

	set := set.NewSet[int](count)
	for set.Len() < count {
		set.Add(SignedMathRandBetween[int](0, len(inp)-1))
	}

	if unsorted {
		return defcon.ConvertIterators(set.Iterator(), func(i int) T { return inp[i] })
	}

	return defcon.IteratorFromList(inp).FilterIndex(func(pos int, _ T) bool { return set.Has(pos) })
}

func ListChoices[T any](inp []T, count int, unsorted bool) []T {
	return ListChoicesIter(inp, count, unsorted).Finalize()
}

func MapChoicesIter[K comparable, V any](inp map[K]V, count int) defcon.MapIterator[K, V] {
	switch {
	case count < 0:
		panic(throw.IllegalValue())
	case count > len(inp):
		count = len(inp)
	}

	set := set.NewSet[int](count)
	for set.Len() < count {
		set.Add(SignedMathRandBetween[int](0, len(inp)-1))
	}

	return defcon.IteratorFromMap(inp, nil).FilterIndex(func(pos int, k K, v V) bool { return set.Has(pos) })
}
