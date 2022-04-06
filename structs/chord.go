package elements

import (
	"errors"
	"fmt"
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

var FilteredIntervals Intervals
var ErrNotAValidChord = errors.New("not a valid chord")

func NewChordWithNoteName(n NoteName, intervals []Interval) (*Chord, error) {
	note := NewNote(int(n))
	return NewChord(note, intervals)
}

func (c *Chord) GetFilteredIntervals() Intervals {
	if len(FilteredIntervals) == 0 {
		FilteredIntervals = Intervals(c.IntervalsToRoot).RemoveDuplicates().ToRelativeIntervals()
	}
	return FilteredIntervals
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
func (c *Chord) hasRoot() bool {
	return Intervals(c.IntervalsToRoot).Contains(Unison)
}

func (c *Chord) has7th() bool {
	return c.GetFilteredIntervals().Contains(Major7th, Minor7th, Diminished7th)
}

func (c *Chord) has6th() bool {
	return c.GetFilteredIntervals().Contains(Major6th, Minor6th)
}

func (c *Chord) has5th() bool {
	return c.GetFilteredIntervals().Contains(Perfect5th, Diminished5th, Augmented5th)
}

func (c *Chord) has4th() bool {
	return c.GetFilteredIntervals().Contains(Perfect4th)
}

func (c *Chord) has3rd() bool {
	return c.GetFilteredIntervals().Contains(Major3rd, Minor3rd)
}

func (c *Chord) has2nd() bool {
	return c.GetFilteredIntervals().Contains(Major2nd, Minor2nd)
}

// TODO: more test cases
// TODO: default to be standard representation, but leave an option
func (c *Chord) ChordName() (string, error) {
	withInversion := func(name string, err error) (string, error) {
		lowest := Intervals(c.IntervalsToRoot).Min()
		if lowest.ToRelative() == Unison {
			return name, err
		}
		// has inversion
		noteName := c.Root.Add(lowest.ToRelative()).NoteName()
		return fmt.Sprintf("%s/%s", noteName, name), err
	}

	if c.has7th() {
		// 13, 11, 9, 7
		if c.has6th() {
			// TODO: 13th chord
			return withInversion(thirteenthChordName(c))
		} else if c.has4th() {
			// TODO: 11th chord
			return withInversion(eleventhChordName(c))
		} else if c.has2nd() {
			// TODO: 9th chord
			return withInversion(ninthChordName(c))
		} else {
			// TODO: 7th chord
			return withInversion(seventhChordName(c))
		}
	} else {
		if c.has6th() {
			// TODO: 6 chord
			return withInversion(sixthChordName(c))
		} else if c.has3rd() {
			// TODO: can be complex
			return withInversion(triadName(c))
		} else if c.has4th() {
			// TODO: sus4
			return withInversion(sus4ChordName(c))
		} else if c.has2nd() {
			// TODO: sus2
			return withInversion(sus2ChordName(c))
		} else if c.has5th() && c.hasRoot() {
			return withInversion(powerChordName(c))
		} else {
			return "", ErrNotAValidChord
		}
	}
}

/*
	TODO: need to accomplish
*/

func thirteenthChordName(c *Chord) (string, error) {
	if !c.has7th() || !c.has6th() {
		return "", errors.New("not a 13th chord")
	}
	return "", nil
}

func eleventhChordName(c *Chord) (string, error) {
	if !c.has7th() || !c.has4th() {
		return "", errors.New("not a 11th chord")
	}
	return "", nil
}

func ninthChordName(c *Chord) (string, error) {
	if !c.has7th() || !c.has2nd() {
		return "", errors.New("not a 9th chord")
	}
	return "", nil
}

func seventhChordName(c *Chord) (string, error) {
	if !c.has7th() {
		return "", errors.New("not a 7th chord")
	}
	return "", nil
}

func sixthChordName(c *Chord) (string, error) {
	if c.has7th() || !c.has6th() {
		return "", errors.New("not a 6th chord")
	}
	return "", nil
}

func sus4ChordName(c *Chord) (string, error) {
	if c.has7th() || c.has6th() || c.has3rd() || !c.has4th() {
		return "", errors.New("not a sus4")
	}
	return "", nil
}

func sus2ChordName(c *Chord) (string, error) {
	if c.has7th() || c.has6th() || c.has4th() || !c.has2nd() {
		return "", errors.New("not a sus2")
	}
	return "", nil
}

func triadName(c *Chord) (string, error) {
	if c.has7th() || c.has6th() {
		return "", errors.New("not a triad")
	}
	return "", nil
}

func powerChordName(c *Chord) (string, error) {
	if len(c.GetFilteredIntervals()) != 2 {
		return "", errors.New("not a power chord")
	}
	if !c.GetFilteredIntervals().Contains(Unison, Perfect5th) {
		return "", errors.New("not a power chord")
	}
	return fmt.Sprintf("%s5", c.Root.NoteName()), nil
}
