package movegenerator

import (
	"fmt"
	"math/bits"
	"math/rand"

	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

func BlockerMasksForSquare(n int8, pieceType int8) []uint64 {
	var result []uint64

	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)

	baseMask := bitops.Mask(n)

	f1 := func(src, dst int8) {
		baseMask = bitops.SetBit(baseMask, dst)
	}
	if pieceType == boardstate.ROOK {
		genSingleRookMovesGeneric(b, n, f1)
	}
	if pieceType == boardstate.BISHOP {
		genSingleBishopMovesGeneric(b, n, f1)
	}
	f := func(occupancyMask uint64) {
		if bitops.TestBit(occupancyMask, n) {
			result = append(result, occupancyMask)
		}
	}

	bitops.Subsets(baseMask, f)

	return result

}

func AttackSetForOccupancy(n int8, occupancy uint64, pieceType int8) uint64 {
	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)
	squares := bitops.FindSetBits(occupancy)
	for _, pos := range squares {
		if pos != n {
			b.SetSquare(pos, boardstate.BLACK, pieceType)
		}
	}

	var result uint64

	f := func(src, dst int8) {
		result = bitops.SetBit(result, dst)
	}
	if pieceType == boardstate.ROOK {
		genSingleRookMovesGeneric(b, n, f)
	}
	if pieceType == boardstate.BISHOP {
		genSingleBishopMovesGeneric(b, n, f)
	}
	return result

}

func PreMask(n int8, pieceType int8) uint64 {
	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)

	baseMask := bitops.Mask(n)

	f1 := func(src, dst int8) {
		baseMask = bitops.SetBit(baseMask, dst)
	}
	if pieceType == boardstate.ROOK {
		genSingleRookMovesGeneric(b, n, f1)
	}
	if pieceType == boardstate.BISHOP {
		genSingleBishopMovesGeneric(b, n, f1)
	}

	return baseMask
}

type MagicDefinition struct {
	square     int8
	mapping    []uint64
	magicValue uint64
	rotate     int8
	preMask    uint64
}

func FindMagic(n int8, preMask uint64, blockers []uint64, attackSets []uint64) *MagicDefinition {
	blockerMaskBits := bits.OnesCount64(preMask) - 1
	rotate := int8(64 - blockerMaskBits)
	totalCount := len(blockers)
	best := 0
	var magicValue uint64

	for {
		mapping := make([]uint64, 1<<blockerMaskBits)
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
					//fmt.Printf("Collision detected at %v of %v with magic %v at cache key %v\n", i, totalCount, magicValue, cacheKey)
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

func GenerateMagicBitboards(pieceType int8) [64]*MagicDefinition {
	var result [64]*MagicDefinition
	var n int8
	for n = 0; n < 64; n++ {
		fmt.Printf("Generating magic %v square for %v\n", pieceType, n)
		blockers := BlockerMasksForSquare(n, pieceType)
		attackSets := make([]uint64, len(blockers))
		preMask := PreMask(n, pieceType)
		for i, blocker := range blockers {
			attackSets[i] = AttackSetForOccupancy(n, blocker, pieceType)
		}
		result[n] = FindMagic(n, preMask, blockers, attackSets)
	}
	return result
}
