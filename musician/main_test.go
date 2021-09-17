package main

import (
	"fmt"
	"github.com/nate-droid/core/chords"
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
	"github.com/nate-droid/core/symphony"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSong(t *testing.T) {
	m := &Musician{}
	songStructure := &symphony.SongStructure{
		Key:              notes.C,
		Mode:             scales.Major,
		ChordProgression: []chords.ChordInterval{chords.I, chords.IV, chords.V},
	}
	song, err := m.playSong(songStructure)
	assert.NoError(t, err)
	fmt.Println(song)
}

func TestCreateSongAfterListen(t *testing.T) {
	m, err := newMusician()
	assert.NoError(t, err)
	songStructure := <-m.ReceiveSongStructure
	song, err := m.playSong(songStructure)
	assert.NoError(t, err)
	fmt.Println(song)
}

