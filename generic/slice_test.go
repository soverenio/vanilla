package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverseInPlace(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "empty",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "one element",
			input:  []int{1},
			output: []int{1},
		},
		{
			name:   "two elements",
			input:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			name:   "three elements",
			input:  []int{1, 2, 3},
			output: []int{3, 2, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, InverseInPlace(test.input))
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		inp []int
		cut int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{inp: []int{}, cut: 0},
			want: []int(nil),
		},
		{
			name: "will be empty",
			args: args{inp: []int{0, 0, 0, 0}, cut: 0},
			want: []int(nil),
		},
		{
			name: "trim from start",
			args: args{inp: []int{0, 0, 0, 0, 1}, cut: 0},
			want: []int{1},
		},
		{
			name: "trim from end",
			args: args{inp: []int{1, 0, 0, 0, 0}, cut: 0},
			want: []int{1},
		},
		{
			name: "trim from both sides",
			args: args{inp: []int{0, 0, 0, 0, 1, 0, 0, 0, 0}, cut: 0},
			want: []int{1},
		},
		{
			name: "trim from both sides 2",
			args: args{inp: []int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}, cut: 0},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "trim from both sides, center ignored",
			args: args{inp: []int{0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0}, cut: 0},
			want: []int{1, 1, 0, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValuesf(t, tt.want, Trim(tt.args.inp, tt.args.cut), "Trim(%v, %v)", tt.args.inp, tt.args.cut)
		})
	}
}

func TestRemoveFirstFromSliceInPlace(t *testing.T) {
	type args struct {
		inp  []int
		elem int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nothing to remove",
			args: args{
				inp:  []int{1, 2, 3},
				elem: 0,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "empty input",
			args: args{
				inp:  []int{},
				elem: 0,
			},
			want: []int{},
		},
		{
			name: "first to remove input",
			args: args{
				inp:  []int{0, 1, 2},
				elem: 0,
			},
			want: []int{1, 2},
		},
		{
			name: "remove from middle",
			args: args{
				inp:  []int{-1, 0, 1},
				elem: 0,
			},
			want: []int{-1, 1},
		},
		{
			name: "remove from end",
			args: args{
				inp:  []int{-2, -1, 0},
				elem: 0,
			},
			want: []int{-2, -1},
		},
		{
			name: "only element",
			args: args{
				inp:  []int{0},
				elem: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveFirstFromSliceInPlace(tt.args.inp, tt.args.elem), "RemoveFirstFromSliceInPlace(%v, %v)", tt.args.inp, tt.args.elem)
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		inp []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil",
			args: args{inp: nil},
			want: nil,
		},
		{
			name: "empty",
			args: args{inp: []int{}},
			want: []int{},
		},
		{
			name: "one element",
			args: args{inp: []int{1}},
			want: []int{},
		},
		{
			name: "ten elements",
			args: args{inp: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			want: []int{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Truncate(tt.args.inp), "Truncate(%v)", tt.args.inp)
			assert.Equalf(t, cap(tt.args.inp), cap(Truncate(tt.args.inp)), "truncate_cap(%v)", tt.args.inp)
		})
	}
}

func TestNewPermutationIterator3(t *testing.T) {
	var (
		cnt  = 0
		iter = NewPermutationIterator([]int{1, 2, 3})
	)

	for {
		next := iter.Next()
		if next == nil {
			break
		}
		cnt += 1
	}

	assert.Equal(t, 6, cnt)
}

func TestNewPermutationIterator4(t *testing.T) {
	var (
		cnt  = 0
		iter = NewPermutationIterator([]int{1, 2, 3, 4})
	)

	for {
		next := iter.Next()
		if next == nil {
			break
		}
		cnt += 1
	}

	assert.Equal(t, 24, cnt)
}

func TestPermutations(t *testing.T) {
	assert.Len(t, Permutations([]int{1}), 1)
	assert.Len(t, Permutations([]int{1, 2}), 2)
	assert.Len(t, Permutations([]int{1, 2, 3}), 6)
	assert.Len(t, Permutations([]int{1, 2, 3, 4}), 24)
}

func TestNewPermutationIterator(t *testing.T) {
	type args struct {
		firstPermutation []int
	}
	tests := []struct {
		name string
		args args
		len  int
	}{
		{
			name: "1",
			args: args{firstPermutation: []int{1}},
			len:  1,
		},
		{
			name: "2",
			args: args{firstPermutation: []int{1, 2}},
			len:  2,
		},
		{
			name: "3",
			args: args{firstPermutation: []int{1, 2, 3}},
			len:  6,
		},
		{
			name: "4",
			args: args{firstPermutation: []int{1, 2, 3, 4}},
			len:  24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				cnt  = 0
				iter = NewPermutationIterator(tt.args.firstPermutation)
			)

			for {
				next := iter.Next()
				if next == nil {
					break
				}
				cnt += 1
			}

			assert.Equalf(t, tt.len, cnt, "NewPermutationIterator(%v)", tt.args.firstPermutation)
		})
	}
}
