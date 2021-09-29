package main

import (
	"context"
	"fmt"
	"github.com/nate-droid/core/chords"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConductor(t *testing.T) {
	symphony := newSymphony()

	cond, err := newConductor()
	err = cond.sendSymphony(symphony)
	assert.NoError(t, err) // TODO make MUST
	<-cond.SymphonyReady
	fmt.Println("weeeeee")
}

func TestSendSong(t *testing.T) {
	// the band is created, now send out instructions
	symphony := newSymphony()
	cond, err := newConductor()
	assert.NoError(t, err) // TODO make MUST

	for _, section := range symphony.Sections {
		for i := 0; i < section.GroupSize; i++ {
			err := cond.sendSongStructure(&symphony.SongStructure)
			assert.NoError(t, err)
		}
	}
	fmt.Println()
}

func TestRun(t *testing.T) {
	cond, err := newConductor()
	assert.NoError(t, err)
	err = cond.Run(context.Background())
	assert.NoError(t, err)
}



func TestImport(t *testing.T) {
	fmt.Println(chords.CommonProgressions)
}
