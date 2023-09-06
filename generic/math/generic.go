package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Comparable interface {
	constraints.Integer | constraints.Float
}

func Max[T Comparable](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Comparable](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Abs[T Comparable](x T) T {
	return T(math.Abs(float64(x)))
}

func Floor[T constraints.Float](x T) T {
	return T(math.Floor(float64(x)))
}

func Ceil[T constraints.Float](x T) T {
	return T(math.Ceil(float64(x)))
}

func Trunc[T constraints.Float](x T) T {
	return T(math.Trunc(float64(x)))
}

func Round[T constraints.Float](x T) T {
	return T(math.Round(float64(x)))
}

func RoundToEven[T constraints.Float](x T) T {
	return T(math.RoundToEven(float64(x)))
}

func IsEven[T constraints.Integer](x T) bool {
	return x%2 == 0
}

func IsOdd[T constraints.Integer](x T) bool {
	return x%2 != 0
}
