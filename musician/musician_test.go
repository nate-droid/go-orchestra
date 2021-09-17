package musician

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/conductor"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/notes"
	"github.com/nate-droid/go-orchestra/core/scales"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSong(t *testing.T) {
	m := &Musician{}
	songStructure := conductor.SongStructure{
		Key:              notes.C,
		Mode:             scales.Major,
		ChordProgression: []chords.ChordInterval{chords.I, chords.IV, chords.V},
	}
	song, err := m.playSong(songStructure)
	assert.NoError(t, err)
	fmt.Println(song)
}