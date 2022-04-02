package elements

import "fmt"

const (
	C      NoteName = 0
	CSharp          = 1
	D               = 2
	DSharp          = 3
	E               = 4
	F               = 5
	FSharp          = 6
	G               = 7
	GSharp          = 8
	A               = 9
	ASharp          = 10
	B               = 11
)

var NoteNameToStrMap = map[NoteName]string{
	C:      "C",
	CSharp: "C#",
	D:      "D",
	DSharp: "D#",
	E:      "E",
	F:      "F",
	FSharp: "F#",
	G:      "G",
	GSharp: "G#",
	A:      "A",
	ASharp: "A#",
	B:      "B",
}

type Note struct {
	IntervalToMidC Interval
}

type NoteName int

func NewNote(n int) *Note {
	return &Note{IntervalToMidC: Interval(n)}
}

func (n *Note) NoteName() NoteName {
	remainder := n.IntervalToMidC % PerfectOctave
	if remainder < 0 {
		remainder += PerfectOctave
	}
	return NoteName(remainder)
}

/*
	SPN is short for Scientific Pitch Notation, also ASPN(American Standard Pitch Notation) or IPN(International Pitch Notation).
	Used to represent the absolute pitch of a note, range from C0.
	In this system, the middle C is represented as 'C4'.
*/
func (n *Note) SPN() string {
	name := n.NoteName()
	group := n.SPNGroup()
	return fmt.Sprintf("%s%d", NoteNameToStrMap[name], group)
}

func (n *Note) SPNGroup() int8 {
	return 4 + int8(n.IntervalToMidC/12)
}

func (n *Note) Add(interval Interval) *Note {
	return &Note{IntervalToMidC: n.IntervalToMidC + interval}
}
