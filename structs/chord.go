package elements

import (
	"errors"
	"sort"
)

/*
	A chord is defined as root note and interval slice to the root.
	For example, a `C Maj` Triad can be represented as:
	Chord {
		Root: C4
		IntervalsToRoot []Interval{Unison, Major3rd, Perfect5th}
	}
*/
type Chord struct {
	Root            *Note
	IntervalsToRoot []Interval
}

const (
	// Chord represent method, for example, a Bm7(b5) can be also represented as Bø
	Standard      = 0
	MinorAsHyphen = 1
	// TODO: major as triangle something can be added
)

func NewChordWithNoteName(n NoteName, intervals []Interval) (*Chord, error) {
	note := NewNote(int(n))
	return NewChord(note, intervals)
}

func NewChord(r *Note, intervals []Interval) (*Chord, error) {
	if len(intervals) == 1 {
		return nil, errors.New("must have more than 2 notes")
	}
	if Intervals(intervals).HasDuplicates() {
		return nil, errors.New("can not have duplicates in intervals")
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i] < intervals[j]
	})
	return &Chord{
		Root:            r,
		IntervalsToRoot: intervals,
	}, nil
}

// containsRoot indicates whether the chord struct concludes its power, for example, {B D F} without G can also be a G7 chord
func (c *Chord) containsRoot() bool {
	return Intervals(c.IntervalsToRoot).Contains(Unison)
}

// Contains only root note and Perfect 5th, but it can be C5/G which the root note is C, and has an interval of -5
func (c *Chord) isPowerChord() bool {
	// transfer to relative intervals, remove the duplicates and sort
	is := Intervals(c.IntervalsToRoot).RemoveDuplicates().ToRelativeIntervals().Sort()
	return len(is) == 2 && is[0] == Unison && is[1] == Perfect5th
}

// fixme: Using relative intervals is inaccurate!!!!
func (c *Chord) isMajorTriad() bool {
	is := Intervals(c.IntervalsToRoot).RemoveDuplicates().ToRelativeIntervals().Sort()
	return (c.containsRoot() && len(is) == 3 && is[1] == Major3rd && is[2] == Perfect5th) ||
		(!c.containsRoot() && len(is) == 2 && is[0] == Major3rd && is[1] == Perfect5th)
}

// TODO: very important
//func (c *Chord) ChordName() string {
//
//}
