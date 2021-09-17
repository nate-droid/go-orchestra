package musician

import (
	"github.com/nate-droid/go-orchestra/chords"
	"github.com/nate-droid/go-orchestra/conductor"
)

type Musician struct {

}

func (m *Musician) playSong(song conductor.Song) ([]chords.Chord, error) {
	return chords.Progression(song.Mode, song.ChordProgression, song.Key)
}