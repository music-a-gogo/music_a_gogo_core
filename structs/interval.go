package elements

// number of semitones
const (
	Unison                               = 0
	Minor2nd                             = 1
	Major2nd                             = 2
	Minor3rd                             = 3
	Major4th                             = 4
	Perfect4th                           = 5
	Augmented4th, Diminished5th, Tritone = 6, 6, 6
	Perfect5th                           = 7
	Minor6th                             = 8
	Major6th                             = 9
	Minor7th                             = 10
	Major7th                             = 11
	PerfectOctave                        = 12
)

/*
	Interval represents distance between two notes in semitones
*/
type Interval int
