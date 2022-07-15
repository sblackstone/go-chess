package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

func RookOccuppancyMasksForSquare(n int8) []uint64 {
	var result []uint64
	rank, file := bitops.SquareToRankFile(n)
	baseMask := (bitops.FileMask(file) | bitops.RankMask(rank))
	f := func(occupancyMask uint64) {
		if bitops.TestBit(occupancyMask, n) {
			result = append(result, occupancyMask)
		}
	}

	bitops.Subsets(baseMask, f)

	return result

}

func RookAttackSetForOccupancy(n int8, occupancy uint64) uint64 {
	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, boardstate.ROOK)
	squares := bitops.FindSetBits(occupancy)
	for _, pos := range squares {
		b.SetSquare(pos, boardstate.BLACK, boardstate.PAWN)
	}

	var result uint64

	f := func(src, dst int8) {
		result = bitops.SetBit(result, dst)
	}
	genSingleRookMovesGeneric(b, n, f)
	return result

}
