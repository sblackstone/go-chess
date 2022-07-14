package bitops

import (
	"fmt"
	"math/bits"
)

var rankMasks [8]uint64
var fileMasks [8]uint64
var clearMasks [64]uint64
var setMasks [64]uint64
var ranks [64]int8
var files [64]int8
var perimeterMask uint64
var internalMask uint64

func init() {
	for pos := 0; pos < 64; pos++ {
		setMasks[pos] = (1 << pos)
		clearMasks[pos] = ^(1 << pos)
		ranks[pos] = int8(pos / 8)
		files[pos] = int8(pos % 8)
	}
	var i int8
	for i = 0; i < 8; i++ {
		rankMasks[i] = 255 << (i * 8)
		fileMasks[i] = Rotate90Clockwise(rankMasks[i])
	}
	perimeterMask = rankMasks[0] | rankMasks[7] | fileMasks[0] | fileMasks[7]
	internalMask = ^perimeterMask
}

func FileMask(file int8) uint64 {
	return fileMasks[file]
}

func RankMask(rank int8) uint64 {
	return rankMasks[rank]
}

func InternalMask() uint64 {
	return internalMask
}
func PerimeterMask() uint64 {
	return perimeterMask
}

func Mask(pos int8) uint64 {
	return setMasks[pos]
}

func SetBit(n uint64, pos int8) uint64 {
	n |= setMasks[pos]
	return n
}

func FindSetBits(n uint64) []int8 {
	var result []int8
	for n > 0 {
		leading := bits.TrailingZeros64(n)
		result = append(result, int8(leading))
		n &= clearMasks[leading]
	}

	return result
}

func ClearBit(n uint64, pos int8) uint64 {
	n &= clearMasks[pos]
	return n
}

func TestBit(n uint64, pos int8) bool {
	return n&(setMasks[pos]) > 0
}

func FlipBit(n uint64, pos int8) uint64 {
	n ^= (setMasks[pos])
	return n
}

func RankFileToSquare(rank int8, file int8) int8 {
	return rank*8 + file
}

func RankOfSquare(n int8) int8 {
	return ranks[n]
}

func FileOfSquare(n int8) int8 {
	return files[n]
}

func SquareToRankFile(n int8) (int8, int8) {
	return ranks[n], files[n]
}

func Print(n uint64, highlight int8) {
	var rank, file int8
	for rank = 7; rank >= 0; rank-- {
		for file = 0; file < 8; file++ {
			pos := RankFileToSquare(rank, file)
			if pos == highlight {
				fmt.Printf(" * ")
			} else if TestBit(n, pos) {
				fmt.Printf(" X ")
			} else {
				fmt.Printf(" - ")
			}
		}
		fmt.Println()
	}

}
