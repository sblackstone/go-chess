package movegenerator

import (
	"fmt"
	"math/bits"
	"math/rand"

	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

func RookBlockerMasksForSquare(n int8) []uint64 {
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
		if pos != n {
			b.SetSquare(pos, boardstate.BLACK, boardstate.PAWN)
		}
	}

	var result uint64

	f := func(src, dst int8) {
		result = bitops.SetBit(result, dst)
	}
	genSingleRookMovesGeneric(b, n, f)
	return result

}

func RookPreMask(n int8) uint64 {
	rank, file := bitops.SquareToRankFile(n)
	return bitops.FileMask(file) | bitops.RankMask(rank)
}

type MagicDefinition struct {
	square     int8
	mapping    [16384]uint64
	magicValue uint64
	rotate     int8
	preMask    uint64
}

func RookFindMagic(n int8) *MagicDefinition {
	blockers := RookBlockerMasksForSquare(n)
	preMask := RookPreMask(n)
	blockerMaskBits := bits.OnesCount64(preMask) - 1
	rotate := int8(64 - blockerMaskBits)
	attackSets := make([]uint64, len(blockers))
	totalCount := len(blockers)
	best := 0
	var magicValue uint64
	for i, blocker := range blockers {
		attackSets[i] = RookAttackSetForOccupancy(n, blocker)
	}

	for {
		var mapping [16384]uint64
		magicValue = rand.Uint64() & rand.Uint64() & rand.Uint64()
		for i, blocker := range blockers {
			cacheKey := (blocker * magicValue) >> rotate
			attackSet := attackSets[i]
			if mapping[cacheKey] == 0 {
				mapping[cacheKey] = attackSet
			}
			if mapping[cacheKey] != attackSet {
				if i > best {
					best = i
					fmt.Printf("Collision detected at %v of %v with magic %v at cache key %v\n", i, totalCount, magicValue, cacheKey)
				}
				break
			}
			if i == totalCount-1 {
				return &MagicDefinition{
					square:     n,
					mapping:    mapping,
					magicValue: magicValue,
					rotate:     rotate,
					preMask:    preMask,
				}
			}
		}
	}
}

func GenerateRookMagicBitboards() [64]*MagicDefinition {
	var result [64]*MagicDefinition
	var i int8
	for i = 0; i < 64; i++ {
		fmt.Printf("Generating magic rook square for %v\n", i)
		result[i] = RookFindMagic(i)
	}
	return result
}
