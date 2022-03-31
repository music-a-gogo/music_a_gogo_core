package elements

import "music_a_gogo/utility"

// modes
const (
	Ionian     = 0
	Dorian     = 1
	Phrygian   = 2
	Lydian     = 3
	Mixolydian = 4
	Aeolian    = 5
	Locrian    = 6
)

// tonality
const (
	Major = 0
	Minor = 1
)

var MajorScaleDiffs = []int{2, 2, 1, 2, 2, 2, 1}

type Scale struct {
	TonicNote int
	Tonality  int
	Mode      int
}

func NewTotalityScale(tonicNote int, tonality int) *Scale {
	return &Scale{
		TonicNote: tonicNote,
		Tonality:  tonality,
		Mode:      Ionian,
	}
}

func NewModeScale(tonicNote int, mode int) *Scale {
	return &Scale{
		TonicNote: tonicNote,
		Tonality:  Major,
		Mode:      mode,
	}
}

func (s *Scale) ConstituentNotes() []*Note {
	var scaleDiffs []int
	startNote := NewNote(s.TonicNote)

	if s.Tonality == Minor {
		scaleDiffs = utility.SliceLeftRolling(MajorScaleDiffs, 5)
	} else {
		scaleDiffs = MajorScaleDiffs
	}

	if s.Mode != Ionian {
		scaleDiffs = utility.SliceLeftRolling(scaleDiffs, s.Mode)
	}

	notes := make([]*Note, len(MajorScaleDiffs)+1)
	notes[0] = startNote
	for i := 1; i < len(notes); i++ {
		notes[i] = notes[i-1].Add(Interval(scaleDiffs[i-1]))
	}
	return notes
}

// TODO: if tonality is minor, show with flat representation
func (s *Scale) ConstituentNotesNames() []string {
	return utility.NoteNames(s.ConstituentNotes())
}
