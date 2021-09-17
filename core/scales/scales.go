package scales

import (
	"github.com/nate-droid/core/notes"
)

type Scale struct {
	Name    string
	Pattern []string
	Modes   []Mode
}

var MajorScale = Scale{
	Name: "Major",
	Modes: []Mode{
		{
			Name:     Ionian,
			Position: 1,
			Pattern:  []int{Whole, Whole, Half, Whole, Whole, Whole, Half},
		},
		{
			Name:     Dorian,
			Position: 2,
			Pattern:  []int{Whole, Half, Whole, Whole, Whole, Half, Whole},
		},
		{
			Name:     Phrygian,
			Position: 3,
			Pattern:  []int{Half, Whole, Whole, Whole, Half, Whole, Whole},
		},
		{
			Name:     Lydian,
			Position: 4,
			Pattern:  []int{Whole, Whole, Whole, Half, Whole, Whole, Half},
		},
		{
			Name:     Mixolydian,
			Position: 5,
			Pattern:  []int{Whole, Whole, Half, Whole, Whole, Half, Whole},
		},
		{
			Name:     Aeolian,
			Position: 6,
			Pattern:  []int{Whole, Half, Whole, Whole, Half, Whole, Whole},
		},
		{
			Name:     Locrian,
			Position: 7,
			Pattern:  []int{Half, Whole, Whole, Half, Whole, Whole, Whole},
		},
	},
}

type ModeList []notes.Note
type ModeName string

type Mode struct {
	Name    ModeName
	Pattern []int
	Position int
}

const (
	Whole = 2
	Half  = 1

	Major      ModeName = "Ionian"
	Ionian     ModeName = "Ionian"
	Dorian     ModeName = "Dorian"
	Phrygian   ModeName = "Phrygian"
	Lydian     ModeName = "Lydian"
	Mixolydian ModeName = "Mixolydian"
	Aeolian    ModeName = "Aeolian"
	Locrian    ModeName = "Locrian"
)

// Major Scale Modes
var ModePatterns = map[ModeName][]int{
	Ionian:     {Whole, Whole, Half, Whole, Whole, Whole, Half},
	Dorian:     {Whole, Half, Whole, Whole, Whole, Half, Whole},
	Phrygian:   {Half, Whole, Whole, Whole, Half, Whole, Whole},
	Lydian:     {Whole, Whole, Whole, Half, Whole, Whole, Half},
	Mixolydian: {Whole, Whole, Half, Whole, Whole, Half, Whole},
	Aeolian:    {Whole, Half, Whole, Whole, Half, Whole, Whole},
	Locrian:    {Half, Whole, Whole, Half, Whole, Whole, Whole},
}

func GetMajorScale(scale notes.Note) (ModeList, error) {
	var majorScale ModeList
	pattern := []int{Whole, Whole, Half, Whole, Whole, Whole, Half}
	// note Index that will traverse the scale
	currentIndex := scale.Index
	for _, interval := range pattern {
		// x, _ := notes.FindIndex(currentIndex)
		nextNoteIndex := currentIndex + interval
		if nextNoteIndex > len(notes.GetAllNotes()) {
			// Do stuff
			nextNoteIndex = nextNoteIndex - len(notes.GetAllNotes())
			note, err := notes.FindIndex(currentIndex)
			if err != nil || note == nil {
				return nil, err
			}
			majorScale = append(majorScale, *note)
			currentIndex = nextNoteIndex
		} else {
			note, err := notes.FindIndex(currentIndex)
			if err != nil || note == nil {
				return nil, err
			}
			majorScale = append(majorScale, *note)
			currentIndex = nextNoteIndex
		}

	}
	return majorScale, nil
}

func GetMode(mode ModeName, startNote notes.Note) (ModeList, error) {
	var modeList ModeList
	pattern := ModePatterns[mode]
	currentIndex := startNote.Index
	for _, interval := range pattern {
		// x, _ := notes.FindIndex(currentIndex)
		nextNoteIndex := currentIndex + interval
		if nextNoteIndex >= len(notes.GetAllNotes()) {
			// Do stuff
			nextNoteIndex = nextNoteIndex - len(notes.GetAllNotes())
			note, err := notes.FindIndex(currentIndex)
			if err != nil || note == nil {
				return nil, err
			}
			modeList = append(modeList, *note)
			currentIndex = nextNoteIndex
		} else {
			note, err := notes.FindIndex(currentIndex)
			if err != nil || note == nil {
				return nil, err
			}
			modeList = append(modeList, *note)
			currentIndex = nextNoteIndex
		}
	}

	return modeList, nil
}
