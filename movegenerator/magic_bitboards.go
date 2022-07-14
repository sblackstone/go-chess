package movegenerator

import "github.com/sblackstone/go-chess/bitops"

func RookOccuppancyMasksForSquare(n int8) []uint64 {
	var result []uint64
	rank, file := bitops.SquareToRankFile(n)
	baseMask := (bitops.FileMask(file) | bitops.RankMask(rank)) & bitops.InternalMask()
	baseMask = bitops.SetBit(baseMask, n)

	f := func(occupancyMask uint64) {
		result = append(result, occupancyMask)
	}

	bitops.Subsets(baseMask, f)

	return result

}
