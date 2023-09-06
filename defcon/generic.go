package defcon

func Reverse[E any](inp []E) []E {
	out := make([]E, 0, len(inp))
	for i := len(inp) - 1; i >= 0; i-- {
		out = append(out, inp[i])
	}
	return out
}

func Contains[E comparable](inp []E, haystack E) bool {
	for _, elem := range inp {
		if elem == haystack {
			return true
		}
	}
	return false
}

func RemoveInplaceUnordered[E comparable](inp []E, haystack E) []E {
	for i := len(inp) - 1; i >= 0; i++ {
		if inp[i] == haystack {
			inp[i] = inp[len(inp)-1]
			inp = inp[:len(inp)-1]
		}
	}

	return inp
}

func ContainsFunc[E1 any, E2 any](inp []E1, haystack E2, eq func(E1, E2) bool) bool {
	for _, elem := range inp {
		if eq(elem, haystack) {
			return true
		}
	}
	return false
}

func Apply[E any](inp []E, f func(E) E) []E {
	out := make([]E, len(inp))
	for pos, elem := range inp {
		out[pos] = f(elem)
	}
	return out
}

func ApplyMap[K comparable, V any](inp map[K]V, f func(V) V) map[K]V {
	out := make(map[K]V, len(inp))
	for pos, elem := range inp {
		out[pos] = f(elem)
	}
	return out
}

func Transform[E1 any, E2 any](inp []E1, f func(E1) E2) []E2 {
	out := make([]E2, len(inp))
	for pos, elem := range inp {
		out[pos] = f(elem)
	}
	return out
}

func TransformMapToList[K comparable, V any, E any](inp map[K]V, f func(V) E) []E {
	out := make([]E, 0, len(inp))
	for _, elem := range inp {
		out = append(out, f(elem))
	}
	return out
}

func UniqueComparableOrdered[V comparable](inp []V) []V {
	out := make([]V, 0, len(inp))
	tmp := make(map[V]struct{}, len(inp))
	for _, elem := range inp {
		if _, ok := tmp[elem]; !ok {
			tmp[elem] = struct{}{}
			out = append(out, elem)
		}
	}
	return out
}

func UniqueComparable[V comparable](inp []V) []V {
	tmp := make(map[V]struct{}, len(inp))
	for _, elem := range inp {
		tmp[elem] = struct{}{}
	}
	out := make([]V, 0, len(tmp))
	for elem := range tmp {
		out = append(out, elem)
	}
	return out
}

func Unique[V any, K comparable](inp []V, keyfunc func(V) K) []V {
	tmp := make(map[K]V, len(inp))
	for _, elem := range inp {
		tmp[keyfunc(elem)] = elem
	}
	out := make([]V, 0, len(tmp))
	for _, elem := range tmp {
		out = append(out, elem)
	}
	return out
}

func Filter[V any](inp []V, filter func(V) bool) []V {
	out := make([]V, 0, len(inp))
	for _, elem := range inp {
		if filter(elem) {
			out = append(out, elem)
		}
	}
	return out
}

func MapHasKey[K comparable, V any](inp map[K]V, key K) bool {
	_, ok := inp[key]
	return ok
}

func ShallowCloneMap[M ~map[K]V, K comparable, V any](inp M) M {
	if inp == nil {
		return nil
	}

	out := make(M, len(inp))
	for k, v := range inp {
		out[k] = v
	}
	return out
}

func ShallowCloneList[M ~[]V, V any](inp M) M {
	if inp == nil {
		return nil
	}

	out := make(M, 0, len(inp))
	out = append(out, inp...)

	return out
}
