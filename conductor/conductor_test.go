package conductor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConductor(t *testing.T) {
	symphony := newSymphony()

	cond, err := newConductor()
	err = cond.sendSymphony(*symphony)
	if err != nil {
		assert.NoError(t, err)
	}
	<-cond.SymphonyReady
	fmt.Println("weeeeee")
}
