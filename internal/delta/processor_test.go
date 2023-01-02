package delta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeltaProcessor(t *testing.T) {
	assert.Panics(t, func() {
		NewDeltaProcessor(nil, nil, false)
	})
}
