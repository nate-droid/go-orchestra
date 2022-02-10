package symphony

import (
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/notes"
	"github.com/nate-droid/go-orchestra/core/scales"
	uuid "github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

type Symphony struct {
	*SongStructure
	Sections []Section
	ID       string
}

type Section struct {
	Type      string
	GroupSize int
}

type SongStructure struct {
	Key              notes.Note
	ChordProgression []chords.ChordInterval
	Mode             scales.ModeName
	SymphonyID       string
	Section
}

type Song struct {
	Progression []chords.ChordInterval
	ChordProgression []chords.Chord
	Section
	SymphonyID string
}

// NewSymphony will randomly generate a symphony
func NewSymphony() *Symphony {
	rand.Seed(time.Now().Unix())

	prog := chords.CommonProgressions[rand.Intn(len(chords.CommonProgressions))]
	key := notes.GetAllNotes()[rand.Intn(notes.ChromaticScaleLength)]

	u, _ := uuid.NewV4()

	return &Symphony{
		SongStructure: &SongStructure{
			Key:              key,
			ChordProgression: prog,
			Mode:             scales.Major,
			SymphonyID:       u.String(),
			Section:          Section{
				Type: "strings",
				GroupSize: 1,
			},
		},
		Sections: []Section{*NewSection()},
		ID: u.String(),
	}
}

func NewSection() *Section {
	return &Section{
		GroupSize: 1,
		Type:      "strings",
	}
}
