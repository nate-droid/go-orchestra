package manager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager(t *testing.T) {
	man, err := newManager()
	if err != nil {
		assert.NoError(t, err)
	}
	d := <-man.WaitForSymphony
	fmt.Println(d)
	man.hireOrchestra(d)
}

func TestDocker(t *testing.T) {
	err := dockStuff()
	assert.NoError(t, err)
}


