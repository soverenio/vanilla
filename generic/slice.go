package generic

func InverseInPlace[T any](inp []T) []T {
	for i := 0; i < len(inp)/2; i++ {
		inp[i], inp[len(inp)-i-1] = inp[len(inp)-i-1], inp[i]
	}
	return inp
}

func Trim[T comparable](inp []T, cut T) []T {
	var (
		begin, end = 0, len(inp)
	)

	for i := 0; i < len(inp); i++ {
		if inp[i] != cut {
			break
		}
		begin = i + 1
	}

	if begin == len(inp) {
		return nil
	}

	for i := len(inp) - 1; i >= 0; i-- {
		if inp[i] != cut {
			break
		}
		end = i
	}

	return inp[begin:end]
}

func RemoveFirstFromSliceInPlace[T comparable](inp []T, elem T) []T {
	for i := 0; i < len(inp); i++ {
		if inp[i] == elem {
			if i != len(inp)-1 {
				copy(inp[i:], inp[i+1:])
			}
			return inp[:len(inp)-1]
		}
	}
	return inp
}

func MergeSlices[T comparable](s1 []T, s2 []T) []T {
	res := append([]T{}, s1...)
	for _, newEl := range s2 {
		found := false
		for _, el := range s1 {
			if newEl == el {
				found = true
				break
			}
		}
		if !found {
			res = append(res, newEl)
		}
	}
	return res
}

func Truncate[T any](inp []T) []T {
	return inp[:0]
}

func Permutations[T any](arr []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}

	helper(arr, len(arr))

	return res
}

type PermutationIterator[T any] struct {
	first       bool
	permutation []int
	elements    []T
}

func NewPermutationIterator[T any](firstPermutation []T) PermutationIterator[T] {
	return PermutationIterator[T]{
		first:       true,
		permutation: make([]int, len(firstPermutation)),
		elements:    firstPermutation,
	}
}

func (iter *PermutationIterator[T]) generatePattern() {
	for i := len(iter.permutation) - 1; i >= 0; i-- {
		if i == 0 || iter.permutation[i] < len(iter.permutation)-i-1 {
			iter.permutation[i]++
			return
		}
		iter.permutation[i] = 0
	}
}

func (iter *PermutationIterator[T]) Next() []T {
	defer iter.generatePattern()

	if iter.first {
		iter.first = false
		return iter.elements
	}

	if iter.permutation[0] >= len(iter.permutation) {
		return nil
	}

	result := append([]T{}, iter.elements...)
	for i, v := range iter.permutation {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
