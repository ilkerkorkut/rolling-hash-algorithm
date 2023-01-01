package signature

type Checksum struct {
	AdlerChecksum uint32   `json:"adler_checksum"`
	MD5Checksum   [16]byte `json:"md5_checksum"`
	Start         int      `json:"start"`
	End           int      `json:"end"`
	Content       []byte   `json:"-"`
}
