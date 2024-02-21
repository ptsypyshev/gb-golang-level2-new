package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "Simple unsorted slice",
			a:    []int{3, 2, 5, 4, 6, 1, 5, 9, 0},
			want: []int{0, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name: "Reversed sorted slice",
			a:    []int{9, 6, 5, 5, 4, 3, 2, 1, 0},
			want: []int{0, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name: "Simple sorted slice",
			a:    []int{0, 1, 2, 3, 4, 5, 5, 6, 9},
			want: []int{0, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name: "Same elements",
			a:    []int{1, 1, 1, 1, 1, 1},
			want: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "Empty slice",
			a:    []int{},
			want: []int{},
		},
		{
			name: "Nil slice",
			a:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		BubbleSort(tt.a)
		assert.Equal(t, tt.want, tt.a)
	}
}
