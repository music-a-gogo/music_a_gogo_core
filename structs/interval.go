package elements

// number of semitones
const (
	Unison                               = 0
	Minor2nd                             = 1
	Major2nd                             = 2
	Minor3rd                             = 3
	Major3rd                             = 4
	Perfect4th                           = 5
	Augmented4th, Diminished5th, Tritone = 6, 6, 6
	Perfect5th                           = 7
	Minor6th, Augmented5th               = 8, 8
	Major6th, Diminished7th              = 9, 9
	Minor7th                             = 10
	Major7th                             = 11
	PerfectOctave                        = 12
	Minor9th                             = 13
	Major9th                             = 14
	Perfect11th                          = 17
	Augmented11th                        = 18
	Minor13rd                            = 20
	Major13rd                            = 21
)

/*
	Interval represents distance between two notes in semitones
*/
type Interval int

/*
	For example, a minus 4th should be a perfect 5th
*/
func (i Interval) ToRelative() Interval {
	v := i % PerfectOctave
	if v < 0 {
		v += PerfectOctave
	}
	return v
}
