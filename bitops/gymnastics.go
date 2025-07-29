package bitops

// Standard bitboard operations...
// Some of these algorithms are originally credited to Knuth.
//
// https://www.chessprogramming.org/Flipping_Mirroring_and_Rotating#Rotationby90degreesClockwise

func FlipDiagA1H8(x uint64) uint64 {
	var t uint64
	var k1 uint64 = 0x5500550055005500
	var k2 uint64 = 0x3333000033330000
	var k4 uint64 = 0x0f0f0f0f00000000
	t = k4 & (x ^ (x << 28))
	x ^= t ^ (t >> 28)
	t = k2 & (x ^ (x << 14))
	x ^= t ^ (t >> 14)
	t = k1 & (x ^ (x << 7))
	x ^= t ^ (t >> 7)
	return x
}

func FlipVertical(x uint64) uint64 {
	return (x << 56) |
		((x << 40) & (0x00ff000000000000)) |
		((x << 24) & (0x0000ff0000000000)) |
		((x << 8) & (0x000000ff00000000)) |
		((x >> 8) & (0x00000000ff000000)) |
		((x >> 24) & (0x0000000000ff0000)) |
		((x >> 40) & (0x000000000000ff00)) |
		(x >> 56)
}

func Rotate90Clockwise(x uint64) uint64 {
	return FlipVertical(FlipDiagA1H8(x))
}

func Rotate90AntiClockwise(x uint64) uint64 {
	return FlipDiagA1H8(FlipVertical(x))
}
