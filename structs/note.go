package elements

import "fmt"

const (
	C      = 0
	CSharp = 1
	D      = 2
	DSharp = 3
	E      = 4
	F      = 5
	FSharp = 6
	G      = 7
	GSharp = 8
	A      = 9
	ASharp = 10
	B      = 11
)

type Note struct {
	IntervalToMidC Interval
}

func NewNote(n int) *Note {
	return &Note{IntervalToMidC: Interval(n)}
}

// TODO: this can help with some scientific calculation
// func (n *Note) GetFrequency() int64 {}

func (n *Note) NoteName() []string {
	remainder := n.IntervalToMidC % PerfectOctave
	if remainder < 0 {
		remainder += PerfectOctave
	}
	switch remainder {
	case 0:
		return []string{"C"}
	case 1:
		return []string{"C sharp", "D flat"}
	case 2:
		return []string{"D"}
	case 3:
		return []string{"D sharp", "E flat"}
	case 4:
		return []string{"E"}
	case 5:
		return []string{"F"}
	case 6:
		return []string{"F sharp", "G flat"}
	case 7:
		return []string{"G"}
	case 8:
		return []string{"G sharp", "A flat"}
	case 9:
		return []string{"A"}
	case 10:
		return []string{"A sharp", "B flat"}
	case 11:
		return []string{"B"}
	default:
		return []string{}
	}
}

/*
	SPN is short for Scientific Pitch Notation, also ASPN(American Standard Pitch Notation) or IPN(International Pitch Notation).
	Used to represent the absolute pitch of a note, range from C0.
	In this system, the middle C is represented as 'C4'.
*/
func (n *Note) SPN() []string {
	names := n.NoteName()
	group := n.SPNGroup()
	spns := make([]string, len(names))
	for i, name := range names {
		spns[i] = fmt.Sprintf("%s%d", name, group)
	}
	return spns
}

func (n *Note) SPNGroup() int8 {
	return 4 + int8(n.IntervalToMidC/12)
}

func (n *Note) Add(interval Interval) *Note {
	return &Note{IntervalToMidC: n.IntervalToMidC + interval}
}
