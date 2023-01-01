package signature_test

import (
	"bytes"
	"testing"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/hashalgos"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/signature"

	"github.com/stretchr/testify/assert"
)

func TestNextChunkHashes(t *testing.T) {
	data := []byte("hello world")
	chunkSize := 2

	reader := bytes.NewReader(data)

	expectedMD5 := hashalgos.MD5Checksum([]byte("he"))
	_, _, expectedAdler := hashalgos.Adler32Checksums([]byte("he"))

	sg := signature.NewSignatureGenerator(reader, chunkSize, false)

	adlerHash, md5Hash, err := sg.NextChunkHashes()

	assert.Equal(t, expectedAdler, adlerHash)
	assert.Equal(t, expectedMD5, md5Hash)

	assert.NoError(t, err)
}
