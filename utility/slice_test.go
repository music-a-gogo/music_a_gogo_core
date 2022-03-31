package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceLeftRolling(t *testing.T) {
	testcases := []struct {
		original []int
		step     int
		want     []int
	}{
		{
			original: []int{1, 2, 3},
			step:     2,
			want:     []int{3, 1, 2},
		},
	}

	for _, tt := range testcases {
		assert.EqualValues(t, tt.want, SliceLeftRolling(tt.original, tt.step))
	}
}
