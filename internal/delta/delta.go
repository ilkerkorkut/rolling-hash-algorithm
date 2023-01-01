package delta

import (
	"encoding/json"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
)

type SingleDelta struct {
	AdlerChecksum uint32    `json:"adler,omitempty"`
	MD5Checksum   *[16]byte `json:"-"`
	Start         int       `json:"start,omitempty"`
	End           int       `json:"end,omitempty"`
	DiffBytes     []byte    `json:"diff,omitempty"`
}

type Delta struct {
	Inserted []*SingleDelta `json:"inserted,omitempty"`
	Deleted  []*SingleDelta `json:"deleted,omitempty"`
	Copied   []*SingleDelta `json:"copied,omitempty"`
}

func NewDelta() *Delta {
	return &Delta{
		Inserted: nil,
		Deleted:  nil,
		Copied:   nil,
	}
}

func (d *Delta) MarshalJSON() ([]byte, error) {
	y, err := json.Marshal(*d)
	if err != nil {
		logging.GetLogger().Errorf("err: %v\n", err)
		return nil, err
	}
	return y, nil
}

func UnmarshalJSON(bytes []byte) (*Delta, error) {
	s := &Delta{}
	err := json.Unmarshal(bytes, s)
	if err != nil {
		logging.GetLogger().Errorf("err: %v\n", err)
		return nil, err
	}
	return s, nil
}
