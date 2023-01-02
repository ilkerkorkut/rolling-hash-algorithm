package delta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDelta(t *testing.T) {
	delta := NewDelta()

	assert.NotNil(t, delta)
	assert.Nil(t, delta.Inserted)
	assert.Nil(t, delta.Copied)
	assert.Nil(t, delta.Deleted)
}

func TestDelta_MarshalJSON(t *testing.T) {
	delta := NewDelta()

	_, err := delta.MarshalJSON()
	assert.Nil(t, err)
}

func TestUnmarshalJSON(t *testing.T) {
	delta := NewDelta()

	deltaJson, _ := delta.MarshalJSON()

	_, err := UnmarshalJSON(deltaJson)

	assert.Nil(t, err)

	_, err = UnmarshalJSON(nil)

	assert.NotNil(t, err)
}
