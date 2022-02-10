package server

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"github.com/nate-droid/go-orchestra/messages/fake"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	m, err := newMem()
	assert.NoError(t, err)
	id, _ := uuid.NewV4()
	song := &symphony.Song{
		ChordProgression: []chords.Chord{{Name: chords.MajorChord}},
		SymphonyID:       id.String(),
	}
	err = m.StoreSong(song)
	assert.NoError(t, err)
}

func TestFetchSymphonies(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	m, err := newMem()
	assert.NoError(t, err)

	results, err := m.FetchAllSongs()
	assert.NoError(t, err)
	fmt.Println("results: ", results)
}
