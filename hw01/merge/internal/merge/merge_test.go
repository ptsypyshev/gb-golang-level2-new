package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedSlices(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple slices",
			args: args{
				a: []int{1, 3, 5, 9},
				b: []int{1, 2, 4, 5, 8},
			},
			want: []int{1, 1, 2, 3, 4, 5, 5, 8, 9},
		},
		{
			name: "All 'a'-elements are greater than 'b'-elements",
			args: args{
				a: []int{6, 7, 8, 9},
				b: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "All 'b'-elements are greater than 'a'-elements",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				b: []int{6, 7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Slices looks same",
			args: args{
				a: []int{1, 3, 5, 7},
				b: []int{1, 3, 5, 7, 9},
			},
			want: []int{1, 1, 3, 3, 5, 5, 7, 7, 9},
		},
		{
			name: "All elements are equal",
			args: args{
				a: []int{1, 1, 1, 1},
				b: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Nil slice 'a'",
			args: args{
				a: nil,
				b: []int{1, 3, 5, 7, 9},
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "Nil slice 'b'",
			args: args{
				a: []int{1, 3, 5, 7},
				b: nil,
			},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "Both slices are nil",
			args: args{
				a: nil,
				b: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MergeSortedSlices(tt.args.a, tt.args.b))
		})
	}
}
