package random

import (
	"github.com/soverenio/vanilla/defcon"
)

func ListShuffleIter[T any](inp []T) defcon.ListIterator[T] {
	permutation := mRandInstance.Perm(len(inp))

	return defcon.ConvertIterators(defcon.IteratorFromList(permutation), func(i int) T { return inp[i] })
}
