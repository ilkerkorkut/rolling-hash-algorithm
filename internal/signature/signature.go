package signature

import (
	"encoding/json"
	"io"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/hashalgos"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
)

type Signature struct {
	Checksums []*Checksum `json:"checksums"`
	ChunkSize int         `json:"chunk_size"`
	Hash      string      `json:"hash"`
}

func (s *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}

func UnmarshalJSON(bytes []byte) (*Signature, error) {
	var s Signature
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

type SignatureGenerator struct {
	reader  io.ReadSeeker
	buf     []byte
	curPos  int
	size    int
	content bool
}

func NewSignatureGenerator(reader io.ReadSeeker, chunkSize int, content bool) *SignatureGenerator {
	return &SignatureGenerator{
		reader:  reader,
		buf:     make([]byte, chunkSize),
		curPos:  0,
		size:    chunkSize,
		content: content,
	}
}

func zeroBytes(s []byte) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func (sg *SignatureGenerator) NextChunkHashes() (uint32, [16]byte, error) {
	n, err := sg.reader.Read(sg.buf)
	sg.curPos += n
	if err != nil {
		return 0, [16]byte{}, err
	}

	_, _, adlerHash := hashalgos.Adler32Checksums(sg.buf)
	if adlerHash == 1 || adlerHash == 0 {
		logging.GetLogger().Fatal("Adler32 hash generated invalid hashsum")
	}

	md5Hash := hashalgos.MD5Checksum(sg.buf)

	return adlerHash, md5Hash, nil
}

func (sg *SignatureGenerator) GenerateSigature() *Signature {
	s := &Signature{
		Checksums: nil,
		ChunkSize: sg.size,
		Hash:      "adler32",
	}

	var ch *Checksum
	for {
		adlerHash, md5Hash, err := sg.NextChunkHashes()
		if err != nil {
			if err == io.EOF {
				if zeroBytes(sg.buf) {
					logging.GetLogger().Error("the file is empty")
					return nil
				}
				return s
			}
			logging.GetLogger().Fatalf("Error reading file: %v", err)
			panic(err)
		}

		ch = &Checksum{
			AdlerChecksum: adlerHash,
			MD5Checksum:   md5Hash,
			Start:         sg.curPos - sg.size,
			End:           sg.curPos,
		}
		if sg.content {
			ch.Content = append(ch.Content, sg.buf...)
		}
		s.Checksums = append(s.Checksums, ch)
	}
}
