package manager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testManager(t *testing.T) {
	man, err := NewManager()
	if err != nil {
		assert.NoError(t, err)
	}
	d := <-man.WaitForSymphony
	fmt.Println(d)
	man.hireOrchestra(d)
}
