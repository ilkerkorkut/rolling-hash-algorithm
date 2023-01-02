package hashalgos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdler32(t *testing.T) {
	chunk := []byte("hello world")
	x, y, c := Adler32Checksums(chunk)

	assert.Equal(t, uint32(0x45d), x)
	assert.Equal(t, uint32(0x1a0b), y)
	assert.Equal(t, uint32(0x1a097db8), c)
}
