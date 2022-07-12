package bitopts

import (
	"errors"
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

var clearMasks [64]uint64
var setMasks [64]uint64
var ranks [64]int8
var files [64]int8

func init() {
	for pos := 0; pos < 64; pos++ {
		setMasks[pos] = (1 << pos)
		clearMasks[pos] = ^(1 << pos)
		ranks[pos] = int8(pos / 8)
		files[pos] = int8(pos % 8)
	}
}

func Mask(pos int8) uint64 {
	return setMasks[pos]
}

func SquareToAlgebraic(pos int8) string {
	rank, file := SquareToRankFile(pos)
	ranks := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return files[int(file)] + ranks[int(rank)]
}

func AlgebraicToSquare(algrbraicSquare string) (int8, error) {
	var result int8

	if len(algrbraicSquare) != 2 {
		return -1, errors.New("algrbraicSquare must be a string of len 2 " + fmt.Sprint(result))
	}

	algrbraicSquare = strings.ToLower(algrbraicSquare)
	var parts = strings.Split(algrbraicSquare, "")
	result = int8(algrbraicSquare[0]) - int8('a')

	if result < 0 || result > 7 {
		return -1, errors.New("Could not convert square value " + algrbraicSquare + ": Invalid file: " + fmt.Sprint(result))
	}

	rank, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, err
	}
	rank -= 1

	if rank > 7 {
		return -1, errors.New("Could not convert square value " + algrbraicSquare + ": Invalid Rank: " + fmt.Sprint(rank))
	}

	result += int8(rank) * 8
	return result, nil
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
