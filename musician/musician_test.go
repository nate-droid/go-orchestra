package musician

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/nate-droid/go-orchestra/chords"
	"github.com/nate-droid/go-orchestra/conductor"
	"github.com/nate-droid/go-orchestra/notes"
	"github.com/nate-droid/go-orchestra/scales"
	"testing"
)

func TestCreateSong(t *testing.T) {
	m := &Musician{}
	song := conductor.Song{
		Key:              notes.C,
		Mode:             scales.Major,
		ChordProgression: []chords.ChordInterval{chords.I, chords.IV, chords.V},
	}
	chords, err := m.playSong(song)
	assert.NoError(t, err)
	fmt.Println(chords)
}