package scales

import (
	"fmt"
	"github.com/nate-droid/go-orchestra/core/notes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStuff(t *testing.T) {
	scale, _ := GetMajorScale(notes.A)
	fmt.Println(scale)
	mode, err := GetMode(Locrian, notes.A)
	assert.NoError(t, err)
	fmt.Println(mode)
}
