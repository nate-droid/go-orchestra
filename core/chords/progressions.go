package chords

import (
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
)

type ChordInterval int

// TODO add names of positions, ie Tonic, subdominant, etc...
const (
	I   ChordInterval = 1
	i	ChordInterval = -1
	II  ChordInterval = 2
	ii	ChordInterval = -2
	III ChordInterval = 3
	iii ChordInterval = -3
	IV  ChordInterval = 4
	iv 	ChordInterval = -4
	V   ChordInterval = 5
	v	ChordInterval = -5
	VI  ChordInterval = 6
	vi 	ChordInterval = -6
	VII ChordInterval = 7
	vii ChordInterval = -7
)

var CommonProgressions = [][]ChordInterval{
	{I, IV, V},
	{I, V, vi, IV},
	{ii, V, I},
	{I, vi, IV, V},
}

func Progression(mode scales.ModeName, intervals []ChordInterval, key notes.Note) ([]Chord, error) {
	// get the scale
	// intervals = I1, IV4, V5
	var progression []Chord
	scale, err := scales.GetMode(mode, key)
	if err != nil {
		return nil, err
	}

	for _, interval := range intervals {
		var chordType ChordType
		if int(interval) > 0 {
			chordType = MajorChord
		} else {
			chordType = MinorChord
			interval = interval * -1
		}
		base := scale[interval-1]

		// TODO need to get current
		// TODO need to get scale from Mode
		// chordTypes, err := GetChordQualitiesForScalePosition(int(ChordInterval), mode)

		chordNotes, err := GetChordTones(base, chordType)
		if err != nil {
			return nil, err
		}
		chord := Chord{
			Root: base.Name,
			Name: chordType,
			Notes: chordNotes,
		}

		progression = append(progression, chord)
	}

	return progression, nil
}
