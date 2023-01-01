package hashalgos

//revive:disable
const MOD_ADLER uint32 = 65521

func Adler32Checksums(chunk []byte) (uint32, uint32, uint32) {
	var x, y uint32 = 1, 0
	for i := 0; i < len(chunk); i++ {
		x = (x + uint32(chunk[i])) % MOD_ADLER
		y = (y + x) % MOD_ADLER
	}
	return x, y, x + y*MOD_ADLER
}

func Adler32Slide(x, y uint32, left, right byte, size int) (uint32, uint32, uint32) {
	l, r := uint32(left), uint32(right)
	x = (x - l + r) % MOD_ADLER
	y = ((y - uint32(size)*l) + x - 1) % MOD_ADLER
	return x, y, x + y*MOD_ADLER
}

//revive:enable
