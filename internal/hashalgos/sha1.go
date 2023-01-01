package hashalgos

import (
	"crypto/sha1"
)

func SHA1Checksum(data []byte) [20]byte {
	return sha1.Sum(data)
}
