package signature

import (
	"bytes"
	"io"
	"testing"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/hashalgos"

	"github.com/stretchr/testify/assert"
)

func TestNewSignatureGenerator(t *testing.T) {
	data := []byte("hello world")
	chunkSize := 2

	reader := bytes.NewReader(data)

	sg := NewSignatureGenerator(reader, chunkSize, false)

	assert.Equal(t, chunkSize, sg.size)
	assert.Equal(t, reader, sg.reader)
}

func TestGenerateSignature(t *testing.T) {
	data := []byte("hello world")
	chunkSize := 2

	reader := bytes.NewReader(data)

	sg := NewSignatureGenerator(reader, chunkSize, false)

	signature := sg.GenerateSigature()

	assert.NotNil(t, signature.Checksums)

	h := hashalgos.MD5Checksum(data[0:chunkSize])

	assert.Equal(t, h, signature.Checksums[0].MD5Checksum)
}

func TestNextChunkHashes(t *testing.T) {
	data := []byte("hello world")
	chunkSize := 2

	reader := bytes.NewReader(data)

	expectedMD5 := hashalgos.MD5Checksum([]byte("he"))
	_, _, expectedAdler := hashalgos.Adler32Checksums([]byte("he"))

	sg := NewSignatureGenerator(reader, chunkSize, false)

	adlerHash, md5Hash, err := sg.NextChunkHashes()

	assert.Equal(t, expectedAdler, adlerHash)
	assert.Equal(t, expectedMD5, md5Hash)
	assert.NoError(t, err)
}

func TestNextChunkHashesWithZeroBytes(t *testing.T) {
	data := []byte{}
	chunkSize := 2

	reader := bytes.NewReader(data)

	sg := NewSignatureGenerator(reader, chunkSize, false)

	_, _, err := sg.NextChunkHashes()

	assert.IsType(t, io.EOF, err)
	assert.Error(t, err)
}

func TestZeroBytes(t *testing.T) {
	data := []byte{0, 0, 0, 0}

	assert.True(t, zeroBytes(data))
}

func TestNonZeroBytes(t *testing.T) {
	data := []byte{12, 2, 4, 0}

	assert.False(t, zeroBytes(data))
}
