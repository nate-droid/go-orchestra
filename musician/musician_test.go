package musician

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/notes"
	"github.com/nate-droid/go-orchestra/core/scales"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"github.com/nate-droid/go-orchestra/messages/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCreateSong(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()
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
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	// m, err := NewMusician()
	// assert.NoError(t, err)
	// songStructure := <-m.ReceiveSongStructure
	//song, err := m.playSong(songStructure)
	//assert.NoError(t, err)
	//fmt.Println(song)
}

