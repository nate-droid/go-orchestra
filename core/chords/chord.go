package chords

import (
	"fmt"
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
)

// Chord is a representation of a musical chord, including a name (maj) a root (C)
//and it's intervals (C, F, G)
type Chord struct {
	Name      ChordType
	Root      notes.Name
	Intervals []Interval
	Notes     []notes.Note
}

type ChordComponents struct {
	Name       ChordType
	Components []Interval
}

// ChordType is a string representation of a chord ie "major" or "diminished"
type ChordType string

const (
	// Chord Names
	MajorChord            ChordType = "maj"
	MinorChord            ChordType = "min"
	DiminishedChord       ChordType = "dim"
	DiminishedSeventh     ChordType = "dim7"
	HalfDiminishedSeventh ChordType = "halfdim7"
	AugmentedTriadChord   ChordType = "aug"
	MajorSixth            ChordType = "M6"
	DominantSeventh       ChordType = "7"
	MajorSeventh          ChordType = "M7"
	AugmentedSeventh      ChordType = "aug7"
	MinorSixth            ChordType = "min6"
	MinorSeventh          ChordType = "min7"
	MinorMajorSeventh     ChordType = "min/maj7"
)

func NewChord(root notes.Name, name ChordType) Chord {
	// TODO this just returns intervals, maybe clean this up?
	ch := Chord{
		Name: name,
		Root: root,
		Intervals: nil,
	}
	components := ChordList[name]

	ch.Intervals = components

	return ch
}

// GetChordTones returns a slice containing the tones that make up a specified chord type
func GetChordTones(note notes.Note, chordType ChordType) ([]notes.Note, error) {
	var chordTones []notes.Note
	components := ChordList[chordType]
	for _, interval := range components {
		nextToneIndex := note.Index + int(interval)
		if nextToneIndex >= len(notes.GetAllNotes()) {
			nextToneIndex = nextToneIndex - len(notes.GetAllNotes())
			tone, err := notes.FindIndex(nextToneIndex)
			if err != nil {
				return nil, err
			}
			chordTones = append(chordTones, *tone)
		} else {
			tone, err := notes.FindIndex(note.Index + int(interval))
			if err != nil {
				return nil, err
			}
			chordTones = append(chordTones, *tone)
		}
	}

	return chordTones, nil
}

// GetChordQualitiesForScalePosition will calculate what chords can be played at a particular mode
func GetChordQualitiesForScalePosition(scalePosition int, scale scales.Scale) ([]ChordType, error) {
	if scalePosition > 6 || scalePosition < 0 {
		return nil, fmt.Errorf("scalePosition must be between 0..6")
	}

	var chordTypes []ChordType
	mode := scale.Modes[scalePosition]
	chordElements := []int{0}
	index := 0

	for _, position := range mode.Pattern {
		chordElements = append(chordElements, index+position)
		index = index + position
	}

	for cType, chordIntervals := range ChordList {
		ok := true
	intervalLoop:
		for _, interval := range chordIntervals {
			_, ok = Find(chordElements, int(interval))
			if !ok {
				break intervalLoop
			}
		}
		if !ok {
			continue
		} else {
			chordTypes = append(chordTypes, cType)
		}

	}

	return chordTypes, nil
}

// Find will check to see if a specified int is present in a given int slice
// and return the index of the int. -1 indicates no value
func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// SameChordTypeSlice will compare to slices of ChordType to determine if they
// have the same ChordType present despite their order
func SameChordTypeSlice(x, y []ChordType) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[ChordType]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
