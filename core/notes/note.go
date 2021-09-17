package notes

import (
	"fmt"
)

type Note struct {
	Name  Name
	Index int
	Frequency float32
	Octave    int
}

type Name string

const (
	aFlat  Name = "Ab"
	a      Name = "A"
	aSharp Name = "A#"
	bFlat  Name = "Bb"
	b      Name = "B"
	c      Name = "C"
	cSharp Name = "C#"
	dFlat  Name = "Db"
	d      Name = "D"
	dSharp Name = "D#"
	eFlat  Name = "Eb"
	e      Name = "E"
	f      Name = "F"
	fSharp Name = "F#"
	gFlat  Name = "Gb"
	g      Name = "G"
	gSharp Name = "G#"
)

type NoteMap map[Name]Note

func GetChromaticScale(sharped bool) []Note {
	scale := []Note{
		{Name: c, Index: 0},
		{Name: "", Index: 1},
		{Name: d, Index: 2},
		{Name: "", Index: 3},
		{Name: e, Index: 4},
		{Name: f, Index: 5},
		{Name: "", Index: 6},
		{Name: g, Index: 7},
		{Name: "", Index: 8},
		{Name: a, Index: 9},
		{Name: "", Index: 10},
		{Name: b, Index: 11},
	}
	if sharped {
		scale[1].Name = cSharp
		scale[3].Name = dSharp
		scale[6].Name = fSharp
		scale[8].Name = gSharp
		scale[10].Name = aSharp
	} else {
		scale[1].Name = dFlat
		scale[3].Name = eFlat
		scale[6].Name = gFlat
		scale[8].Name = aFlat
		scale[10].Name = bFlat
	}
	return scale
}

// TODO func GetNotes() returns all notes in slice
var AFlat = Note{
	Name:  aFlat,
	Index: 8,
}
var A = Note{
	Name:  a,
	Index: 9,
}
var ASharp = Note{
	Name:  aSharp,
	Index: 10,
}
var BFlat = Note{
	Name:  bFlat,
	Index: 10,
}
var B = Note{
	Name:  b,
	Index: 11,
}
var C = Note{
	Name:  c,
	Index: 0,
}
var CSharp = Note{
	Name:  cSharp,
	Index: 1,
}
var DFlat = Note{
	Name:  dFlat,
	Index: 1,
}
var D = Note{
	Name:  d,
	Index: 2,
}
var DSharp = Note{
	Name:  dSharp,
	Index: 3,
}
var EFlat = Note{
	Name:  eFlat,
	Index: 3,
}
var E = Note{
	Name:  e,
	Index: 4,
}
var F = Note{
	Name:  f,
	Index: 5,
}
var FSharp = Note{
	Name:  fSharp,
	Index: 6,
}
var GFlat = Note{
	Name:  gFlat,
	Index: 6,
}
var G = Note{
	Name:  g,
	Index: 7,
}
var GSharp = Note{
	Name:  gSharp,
	Index: 8,
}

// TODO this might not be needed
func GetAllNotes() []Note {
	noteList := []Note{
		A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp,
	}
	return noteList
}

func FindIndex(index int) (*Note, error) {
	for _, note := range GetAllNotes() {
		if index == note.Index {
			return &note, nil
		}
	}
	return nil, fmt.Errorf("index not found")
}

// Frequencies
/*
C1	32.70	1054.94
C#1/Db1 	34.65	995.73
D1	36.71	939.85
D#1/Eb1 	38.89	887.10
E1	41.20	837.31
F1	43.65	790.31
F#1/Gb1 	46.25	745.96
G1	49.00	704.09
G#1/Ab1 	51.91	664.57
A1	55.00	627.27
A#1/Bb1 	58.27	592.07
B1	61.74	558.84
C2	65.41	527.47
C#2/Db2 	69.30	497.87
D2	73.42	469.92
D#2/Eb2 	77.78	443.55
E2	82.41	418.65
F2	87.31	395.16
F#2/Gb2 	92.50	372.98
G2	98.00	352.04
G#2/Ab2 	103.83	332.29
A2	110.00	313.64
A#2/Bb2 	116.54	296.03
B2	123.47	279.42
C3	130.81	263.74
C#3/Db3 	138.59	248.93
D3	146.83	234.96
D#3/Eb3 	155.56	221.77
E3	164.81	209.33
F3	174.61	197.58
F#3/Gb3 	185.00	186.49
G3	196.00	176.02
G#3/Ab3 	207.65	166.14
A3	220.00	156.82
A#3/Bb3 	233.08	148.02
B3	246.94	139.71
C4	261.63	131.87
C#4/Db4 	277.18	124.47
D4	293.66	117.48
D#4/Eb4 	311.13	110.89
E4	329.63	104.66
F4	349.23	98.79
F#4/Gb4 	369.99	93.24
G4	392.00	88.01
G#4/Ab4 	415.30	83.07
A4	440.00	78.41
A#4/Bb4 	466.16	74.01
B4	493.88	69.85
C5	523.25	65.93
C#5/Db5 	554.37	62.23
D5	587.33	58.74
D#5/Eb5 	622.25	55.44
E5	659.25	52.33
F5	698.46	49.39
F#5/Gb5 	739.99	46.62
G5	783.99	44.01
G#5/Ab5 	830.61	41.54
A5	880.00	39.20
A#5/Bb5 	932.33	37.00
B5	987.77	34.93
C6	1046.50	32.97
C#6/Db6 	1108.73	31.12
D6	1174.66	29.37
D#6/Eb6 	1244.51	27.72
E6	1318.51	26.17
F6	1396.91	24.70
F#6/Gb6 	1479.98	23.31
G6	1567.98	22.00
G#6/Ab6 	1661.22	20.77
A6	1760.00	19.60
A#6/Bb6 	1864.66	18.50
B6	1975.53	17.46
C7	2093.00	16.48
C#7/Db7 	2217.46	15.56
D7	2349.32	14.69
D#7/Eb7 	2489.02	13.86
E7	2637.02	13.08
F7	2793.83	12.35
F#7/Gb7 	2959.96	11.66
G7	3135.96	11.00
G#7/Ab7 	3322.44	10.38
A7	3520.00	9.80
A#7/Bb7 	3729.31	9.25
B7	3951.07	8.73
C8	4186.01	8.24
C#8/Db8 	4434.92	7.78
D8	4698.63	7.34
D#8/Eb8 	4978.03	6.93
E8	5274.04	6.54
F8	5587.65	6.17
F#8/Gb8 	5919.91	5.83
G8	6271.93	5.50
G#8/Ab8 	6644.88	5.19
A8	7040.00	4.90
A#8/Bb8 	7458.62	4.63
B8	7902.13	4.37

*/
