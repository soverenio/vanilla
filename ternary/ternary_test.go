package ternary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name string
		cond bool
		a    interface{}
		b    interface{}
		want interface{}
	}{
		{
			name: "true",
			cond: true,
			a:    1,
			b:    2,
			want: 1,
		},
		{
			name: "false",
			cond: false,
			a:    1,
			b:    2,
			want: 2,
		},
		{
			name: "true string",
			cond: true,
			a:    "a",
			b:    "b",
			want: "a",
		},
		{
			name: "false string",
			cond: false,
			a:    "a",
			b:    "b",
			want: "b",
		},
		{
			name: "true struct",
			cond: true,
			a:    struct{ foo string }{foo: "foo"},
			b:    struct{ foo string }{foo: "bar"},
			want: struct{ foo string }{foo: "foo"},
		},
		{
			name: "false struct",
			cond: false,
			a:    struct{ foo string }{foo: "foo"},
			b:    struct{ foo string }{foo: "bar"},
			want: struct{ foo string }{foo: "bar"},
		},
		{
			name: "true slice",
			cond: true,
			a:    []string{"foo", "bar", "baz"},
			b:    []string{"baz"},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "false slice",
			cond: false,
			a:    []string{"foo", "bar", "baz"},
			b:    []string{"baz"},
			want: []string{"baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, If(tt.cond, tt.a, tt.b))
		})
	}

	add := func(a, b int) int { return a + b }
	sub := func(a, b int) int { return a - b }

	t.Run("true func", func(t *testing.T) {
		res := If(true, add, sub)
		assert.Equal(t, 17, res(10, 7))
	})

	t.Run("false func", func(t *testing.T) {
		res := If(false, add, sub)
		assert.Equal(t, 3, res(10, 7))
	})
}
