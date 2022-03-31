package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale_ConstituentNotes(t *testing.T) {
	testcases := []struct {
		scale *Scale
		want  []*Note
	}{
		{
			scale: NewModeScale(C, Ionian),
			want:  []*Note{NewNote(C), NewNote(D), NewNote(E), NewNote(F), NewNote(G), NewNote(A), NewNote(B), NewNote(C + 12)},
		},
	}

	for _, tt := range testcases {
		assert.EqualValues(t, tt.scale.ConstituentNotes(), tt.want)
	}
}

// TODO: write test for constituent notes names
