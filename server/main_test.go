package main

import (
	"fmt"
	"github.com/nate-droid/core/chords"
	"github.com/nate-droid/core/symphony"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	fmt.Println("hey")
	s, err := newStore()
	assert.NoError(t, err)
	id, _ := uuid.NewV4()
	song := &symphony.Song{
		ChordProgression: []chords.Chord{{Name: chords.MajorChord}},
		SymphonyID: id.String(),
	}
	err = s.SaveSong(song)
	assert.NoError(t, err)
}

func TestFetch(t *testing.T) {
	s, err := newStore()
	assert.NoError(t, err)
	err = s.FetchSong("")
	assert.NoError(t, err)
}

func TestFetchSymphonies(t *testing.T) {
	s, err := newStore()
	assert.NoError(t, err)
	_, err = s.FetchSymphoniesIDs()
	assert.NoError(t, err)
	results, err := s.FetchLatestSymphonies()
	assert.NoError(t, err)
	fmt.Println("results: ", results)
}

