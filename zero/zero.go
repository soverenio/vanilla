package zero

import (
	"reflect"
)

func Zero[T any]() T {
	var rv T
	return rv
}

func IsZero[T comparable](inp T) bool {
	return inp == Zero[T]()
}

func FirstOf2[V1 any, V2 any](a V1, _ V2) V1 {
	return a
}

func SecondOf2[V1 any, V2 any](_ V1, a V2) V2 {
	return a
}

func FirstOf3[V1 any, V2 any, V3 any](a V1, _ V2, _ V3) V1 {
	return a
}

func SecondOf3[V1 any, V2 any, V3 any](_ V1, a V2, _ V3) V2 {
	return a
}

func ThirdOf3[V1 any, V2 any, V3 any](_ V1, _ V2, a V3) V3 {
	return a
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[V1 any](v1 V1, err error) V1 {
	if err != nil {
		panic(err)
	}
	return v1
}

func Must2[V1 any, V2 any](v1 V1, v2 V2, err error) (V1, V2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

func Must3[V1 any, V2 any, V3 any](v1 V1, v2 V2, v3 V3, err error) (V1, V2, V3) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}

func GetType[T any]() reflect.Type {
	return reflect.TypeOf(Zero[T]())
}

func Curry1[V1 any](f func(V1), v1 V1) func() {
	return func() { f(v1) }
}
