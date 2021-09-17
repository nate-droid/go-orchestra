package chords

import (
	"github.com/nate-droid/go-orchestra/notes"
	"github.com/nate-droid/go-orchestra/scales"
)

// TODO not in love with the name
type ChordInterval int

// TODO add names of positions, ie Tonic, subdominant, etc...
const (
	I   ChordInterval = 1
	II  ChordInterval = 2
	III ChordInterval = 3
	IV  ChordInterval = 4
	V   ChordInterval = 5
	VI  ChordInterval = 6
	VII ChordInterval = 7
)

func Progression(mode scales.ModeName, intervals []ChordInterval, key notes.Note) ([]Chord, error) {
	// get the scale
	var progression []Chord
	scale, err := scales.GetMode(mode, key)
	if err != nil {
		return nil, err
	}

	for _, interval := range intervals {
		base := scale[interval-1]

		// TODO start with only maj/min
		// TODO need to get current
		// TODO need to get scale from Mode
		// chordTypes, err := GetChordQualitiesForScalePosition(int(ChordInterval), mode)

		chord := Chord{
			Root: base.Name, // TODO add the next notes

		}
		// TODO calculate maj/min etc from scale degree
		progression = append(progression, chord)
	}

	// todo add chord flavors
	return progression, nil
}
