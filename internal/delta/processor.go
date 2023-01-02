package delta

import (
	"io"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/hashalgos"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/signature"
)

type DeltaProcessor struct {
	reader      io.ReadSeeker
	signature   *signature.Signature
	delta       *Delta
	chunkSize   int
	window      []byte
	pos         int
	x           uint32
	y           uint32
	sum         uint32
	visited     []byte
	foundHashes map[uint32]bool
	content     bool
}

func NewDeltaProcessor(signature *signature.Signature, reader io.ReadSeeker, content bool) *DeltaProcessor {
	if signature.ChunkSize == 0 {
		logging.GetLogger().Fatal("The size of the chunk is invalid: 0")
	}
	return &DeltaProcessor{
		signature:   signature,
		reader:      reader,
		pos:         0,
		chunkSize:   signature.ChunkSize,
		window:      make([]byte, signature.ChunkSize),
		visited:     make([]byte, 0),
		foundHashes: map[uint32]bool{},
		content:     content,
		delta:       NewDelta(),
	}
}

func (dp *DeltaProcessor) BuildDelta() *Delta {
	adlerChecksumsMap := map[uint32][16]byte{}

	for _, chs := range dp.signature.Checksums {
		adlerChecksumsMap[chs.AdlerChecksum] = chs.MD5Checksum
	}

	for {
		err := dp.Roll(adlerChecksumsMap)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		dp.foundHashes[dp.sum] = true
	}

	for _, item := range dp.signature.Checksums {
		var d *SingleDelta
		if _, ok := dp.foundHashes[item.AdlerChecksum]; !ok {
			d = &SingleDelta{
				AdlerChecksum: item.AdlerChecksum,
				MD5Checksum:   &item.MD5Checksum,
				Start:         item.Start,
				End:           item.End,
			}

			if dp.content && len(item.Content) != 0 {
				d.DiffBytes = append(d.DiffBytes, item.Content...)
			}

			dp.delta.Deleted = append(dp.delta.Deleted, d)
		}
	}

	return dp.delta
}

func (dp *DeltaProcessor) AppendDiff(md5Hash [16]byte) bool {
	newMD5Hash := hashalgos.MD5Checksum(dp.window)

	if md5Hash == newMD5Hash {
		d := &SingleDelta{
			AdlerChecksum: dp.sum,
			MD5Checksum:   &newMD5Hash,
			Start:         dp.pos - dp.chunkSize,
			End:           dp.pos,
		}

		if dp.content && len(dp.window) != 0 {
			d.DiffBytes = append(d.DiffBytes, dp.window...)
		}

		dp.delta.Copied = append(dp.delta.Copied, d)
		return true
	}

	return false
}

func (dp *DeltaProcessor) SetRemainingBytes() {
	if len(dp.visited) != 0 {
		dp.delta.Inserted = append(dp.delta.Inserted, &SingleDelta{
			Start:     dp.pos - len(dp.visited),
			End:       dp.pos,
			DiffBytes: dp.visited,
		})
	}
}

func (dp *DeltaProcessor) Roll(checksums map[uint32][16]byte) error {
	n, err := dp.reader.Read(dp.window)
	dp.pos += n
	if err != nil {
		if err == io.EOF {
			dp.SetRemainingBytes()
		}
		return err
	}

	dp.x, dp.y, dp.sum = hashalgos.Adler32Checksums(dp.window)
	if md5Hash, ok := checksums[dp.sum]; ok {
		if ok := dp.AppendDiff(md5Hash); ok {
			return nil
		}
	}

	dp.visited = append(dp.visited, dp.window...)

	for {
		err = dp.Next()
		dp.foundHashes[dp.sum] = true

		if err != nil {
			if err == io.EOF {
				dp.SetRemainingBytes()
				return err
			}
			logging.GetLogger().Fatalf("Error while rolling the window: %s", err)
		}

		if md5Hash, ok := checksums[dp.sum]; ok {
			if dp.visited != nil {
				dp.delta.Inserted = append(dp.delta.Inserted, &SingleDelta{
					Start:     dp.pos - len(dp.visited),
					End:       dp.pos,
					DiffBytes: dp.visited,
				})
				dp.visited = nil
			}

			dp.AppendDiff(md5Hash)
			return nil
		}
		dp.visited = append(dp.visited, dp.window[len(dp.window)-1])
	}
}

func (dp *DeltaProcessor) Next() error {
	prev := dp.window[0]
	next := make([]byte, 1)

	n, err := dp.reader.Read(next)
	dp.pos += n
	if err != nil {
		return err
	}

	dp.x, dp.y, dp.sum = hashalgos.Adler32Slide(dp.x, dp.y, prev, next[0], dp.chunkSize)
	dp.window = append(dp.window[1:], next[0])

	return nil
}
