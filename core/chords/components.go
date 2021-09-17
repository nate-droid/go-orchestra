package chords

const (
	P1          Interval = 0
	minor2      Interval = 1
	major2      Interval = 2
	minor3      Interval = 3
	major3      Interval = 4
	p4          Interval = 5
	diminished5 Interval = 6
	P5          Interval = 7
	A5          Interval = 8
	major6dim7  Interval = 9
	minor7      Interval = 10
	major7      Interval = 11
)

type Interval int
type Component []Interval
type ChordMap map[ChordType][]Interval

// TODO find a way to manage chords better

var ChordList = ChordMap{
	MajorChord:            []Interval{P1, major3, P5},
	MinorChord:            []Interval{P1, minor3, P5},
	DiminishedChord:       []Interval{P1, minor3, diminished5},
	AugmentedTriadChord:   []Interval{P1, major3, A5},
	DiminishedSeventh:     []Interval{P1, minor3, diminished5, major6dim7},
	HalfDiminishedSeventh: []Interval{P1, minor3, diminished5, minor7},
	MajorSixth:            []Interval{P1, major3, P5, major6dim7},
	MajorSeventh:          []Interval{P1, major3, P5, major7},
	MinorSeventh:          []Interval{P1, minor3, P5, minor7},
}

// Triads is a list of triad chords
var Triads = ChordMap{
	MajorChord:          ChordList[MajorChord],
	MinorChord:          ChordList[MinorChord],
	DiminishedChord:     ChordList[DiminishedChord],
	AugmentedTriadChord: ChordList[AugmentedTriadChord],
}
