package elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	type want struct {
		wantName     []string
		wantSpnGroup int8
		wantSpn      []string
	}
	cases := []struct {
		desc string
		note *Note
		want want
	}{
		{
			desc: "01. mid C",
			note: &Note{IntervalToMidC: 0},
			want: want{wantName: []string{"C"}, wantSpnGroup: 4, wantSpn: []string{"C4"}},
		},
	}
	for _, tt := range cases {
		assert.Equal(t, tt.want.wantName, tt.note.NoteName())
		assert.Equal(t, tt.want.wantSpnGroup, tt.note.SPNGroup())
		assert.Equal(t, tt.want.wantSpn, tt.note.SPN())
	}
}
