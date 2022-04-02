package elements

import (
	"errors"
	"music_a_gogo/utility"
)

// modes
const (
	Ionian     = 0
	Dorian     = 1
	Phrygian   = 2
	Lydian     = 3
	Mixolydian = 4
	Aeolian    = 5
	Locrian    = 6
	Major      = 7
	Minor      = 8
)

// TODO: Some of the Intervals should be provided in an isolated file, and provided publicly
var (
	// MajorScaleIntervals is Intervals for a Major Scale or an Ionian Scale
	MajorScaleIntervals = []int{2, 2, 1, 2, 2, 2, 1}
)

type Scale struct {
	Root      *Note      // like C or C#
	Intervals []Interval // for a major scale, it's [2,2,1,2,2,2,1]
}

func (s *Scale) intervalOfIntervals() int {
	return utility.IntervalSum(s.Intervals)
}

func NewScale(root *Note, nums []int) (*Scale, error) {
	// TODO: check the legitimacy of the root
	intervals := make([]Interval, len(nums))
	for i, v := range nums {
		// Check the legitimacy of the interval slice
		if v < 0 {
			return nil, errors.New("interval cannot be negative")
		}
		// TODO: the interval slice should not be something repeated, it's nonsense, like {1, 2, 3, 1, 2, 3} should be {1, 2, 3}
		intervals[i] = Interval(v)
	}

	return &Scale{
		Root:      root,
		Intervals: intervals,
	}, nil
}

func (s *Scale) NthNoteName(n int) NoteName {
	steps := s.intervalOfIntervals()*(n/len(s.Intervals)) + utility.IntervalSum(s.Intervals[:(n%len(s.Intervals))])
	note := s.Root.Add(Interval(steps))
	return note.NoteName()
}

func (s *Scale) UnrepeatedSequence() []NoteName {
	totalIntervals := 0
	for _, v := range s.Intervals {
		totalIntervals += int(v)
	}
	sequenceLen := utility.LCM(PerfectOctave, totalIntervals) / totalIntervals * len(s.Intervals)
	sequence := make([]NoteName, sequenceLen)
	for i := 0; i < sequenceLen; i++ {
		sequence[i] = s.NthNoteName(i)
	}
	return sequence
}

// ConstituentNoteNames return a set containing all the note names which can not be repeated
func (s *Scale) ConstituentNoteNames() []NoteName {
	set := make(map[NoteName]struct{})
	notes := make([]NoteName, 0)
	for _, v := range s.UnrepeatedSequence() {
		_, ok := set[v]
		if !ok {
			set[v] = struct{}{}
			notes = append(notes, v)
		}
	}
	return notes
}
