package defcon

import (
	"github.com/soverenio/vanilla/zero"
)

func ConvertIterators[I any, O any](inp ListIterator[I], cb func(I) O) ListIterator[O] {
	inpImplementation := inp.(*listIterator[I])

	return &listIterator[O]{
		generator: func() (O, bool) {
			i, ok := inpImplementation.generator()
			if !ok {
				return zero.Zero[O](), false
			}
			return cb(i), true
		},
		maxElements: inpImplementation.maxElements,
	}
}

func ConvertListToMapIterator[OK comparable, IV, OV any](
	inp ListIterator[IV], cb func(IV) (OK, OV),
) MapIterator[OK, OV] {
	inpImplementation := inp.(*listIterator[IV])

	return &mapIterator[OK, OV]{
		generator: func() (OK, OV, bool) {
			k, continued := inpImplementation.generator()
			if !continued {
				return zero.Zero[OK](), zero.Zero[OV](), false
			}
			ok, ov := cb(k)
			return ok, ov, true
		},
		maxElements: inpImplementation.maxElements,
	}
}

func ConvertMapToListIterator[IK comparable, IV, OV any](inp MapIterator[IK, IV], cb func(IK, IV) OV) ListIterator[OV] {
	inpImplementation := inp.(*mapIterator[IK, IV])

	return &listIterator[OV]{
		generator: func() (OV, bool) {
			k, v, ok := inpImplementation.generator()
			if !ok {
				return zero.Zero[OV](), false
			}
			return cb(k, v), true
		},
		maxElements: inpImplementation.maxElements,
	}
}

func ConvertMapIterators[IK, OK comparable, IV, OV any](
	inp MapIterator[IK, IV], cb func(IK, IV) (OK, OV),
) MapIterator[OK, OV] {
	inpImplementation := inp.(*mapIterator[IK, IV])

	return &mapIterator[OK, OV]{
		generator: func() (OK, OV, bool) {
			k, v, continued := inpImplementation.generator()
			if !continued {
				return zero.Zero[OK](), zero.Zero[OV](), false
			}
			ok, ov := cb(k, v)
			return ok, ov, true
		},
		maxElements: inpImplementation.maxElements,
	}
}

func MergeIterators[V any](iters ...ListIterator[V]) ListIterator[V] {
	if len(iters) == 0 {
		return &listIterator[V]{
			generator:   func() (V, bool) { return zero.Zero[V](), false },
			maxElements: 0,
		}
	}

	maxElements := 0
	for _, iter := range iters {
		inMaxElements := iter.(*listIterator[V]).maxElements
		if inMaxElements < 0 {
			maxElements = -1
			break
		}
		maxElements += inMaxElements
	}

	var pos = 0
	return &listIterator[V]{
		generator: func() (V, bool) {
			for {
				v, continued := iters[pos].(*listIterator[V]).generator()
				switch {
				case continued:
					return v, true
				case pos == len(iters)-1:
					return zero.Zero[V](), false
				default:
					pos++
					continue
				}
			}
		},
		maxElements: maxElements,
	}
}
