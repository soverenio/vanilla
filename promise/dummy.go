package promise

import (
	"context"

	"github.com/soverenio/vanilla/zero"
)

func NewDummy[T any](ctx context.Context) *Simple[T] {
	return NewSimpleStarted[T](ctx, func(ctx context.Context) (T, error) { return zero.Zero[T](), nil })
}
