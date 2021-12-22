package notes

import (
	"fmt"
)

type Note struct {
	Name      Name
	Index     int
	Frequency float32
	Octave    int
	IsFlat    bool
}

type Name string

var ChromaticScaleLength = 12
var ChromaticScale = map[Name]Note{
	"C":  {Name: "C", Index: 0, Frequency: 261.63, Octave: 4},
	"C#": {Name: "C#", Index: 1, Frequency: 277.18, Octave: 4},
	"Db": {Name: "Db", Index: 1, Frequency: 277.18, Octave: 4, IsFlat: true},
	"D":  {Name: "D", Index: 2, Frequency: 293.66, Octave: 4},
	"D#": {Name: "D#", Index: 3, Frequency: 311.13, Octave: 4},
	"Eb": {Name: "Eb", Index: 3, Frequency: 311.13, Octave: 4, IsFlat: true},
	"E":  {Name: "E", Index: 4, Frequency: 329.63, Octave: 4},
	"F":  {Name: "F", Index: 5, Frequency: 349.23, Octave: 4},
	"F#": {Name: "F#", Index: 6, Frequency: 369.99, Octave: 4},
	"Gb": {Name: "Gb", Index: 6, Frequency: 369.99, Octave: 4, IsFlat: true},
	"G":  {Name: "G", Index: 7, Frequency: 392.00, Octave: 4},
	"G#": {Name: "G#", Index: 8, Frequency: 415.30, Octave: 4},
	"Ab": {Name: "Ab", Index: 8, Frequency: 415.30, Octave: 4, IsFlat: true},
	"A":  {Name: "A", Index: 9, Frequency: 440.00, Octave: 4},
	"A#": {Name: "A#", Index: 10, Frequency: 466.16, Octave: 4},
	"Bb": {Name: "Bb", Index: 10, Frequency: 466.16, Octave: 4, IsFlat: true},
	"B":  {Name: "B", Index: 11, Frequency: 493.88, Octave: 4},
}

// calculate intervals and "names" based on circle of fifths / modulo 12
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

var AFlat = ChromaticScale[aFlat]
var A = ChromaticScale[a]
var ASharp = ChromaticScale[aSharp]
var BFlat = ChromaticScale[bFlat]
var B = ChromaticScale[b]
var C = ChromaticScale[c]
var CSharp = ChromaticScale[cSharp]
var DFlat = ChromaticScale[dFlat]
var D = ChromaticScale[d]
var DSharp = ChromaticScale[dSharp]
var EFlat = ChromaticScale[eFlat]
var E = ChromaticScale[e]
var F = ChromaticScale[f]
var FSharp = ChromaticScale[fSharp]
var GFlat = ChromaticScale[gFlat]
var G = ChromaticScale[g]
var GSharp = ChromaticScale[gSharp]

// TODO might not be useful, but just a placeholder for me to remember a few different approaches
var testChromaticScale = struct {
	A      Note
	ASharp Note
}{
	A: ChromaticScale[a],
}

func GetAllNotes() []Note {
	noteList := []Note{
		A, ASharp, B, C, CSharp, D, DSharp, E, F, FSharp, G, GSharp,
	}
	return noteList
}

func FindIndex(index int, isFlat bool) (*Note, error) {
	for _, note := range ChromaticScale {
		if !isFlat && note.IsFlat {
			// this is an enharmonic equivalent. "right note" but wrong context
			continue
		}
		if index == note.Index {
			return &note, nil
		}
	}
	return nil, fmt.Errorf("index not found")
}
