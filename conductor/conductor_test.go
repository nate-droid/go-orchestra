package conductor

import (
	"context"
	"fmt"
	"github.com/nate-droid/go-orchestra/core/chords"
	"github.com/nate-droid/go-orchestra/core/symphony"
	"github.com/nate-droid/go-orchestra/messages/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConductor(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	s := symphony.NewSymphony()

	cond, err := NewConductor()
	if err != nil {
		t.Fatal(err)
	}
	err = cond.sendSymphony(s)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSendSong(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	// the band is created, now send out instructions
	s := symphony.NewSymphony()
	cond, err := NewConductor()
	if err != nil {
		t.Fatal(err)
	}

	for _, section := range s.Sections {
		for i := 0; i < section.GroupSize; i++ {
			err := cond.sendSongStructure(s.SongStructure)
			assert.NoError(t, err)
		}
	}

}

func testRun(t *testing.T) {
	fakeNats := fake.RunDefaultServer()
	defer fakeNats.Shutdown()

	cond, err := NewConductor()
	assert.NoError(t, err)
	err = cond.Run(context.Background())
	assert.NoError(t, err)
}

func TestImport(t *testing.T) {
	fmt.Println(chords.CommonProgressions)
}
