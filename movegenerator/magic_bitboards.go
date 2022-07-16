package movegenerator

import (
	"math/bits"
	"math/rand"

	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

func blockersForSquare(n int8, pieceType int8) []uint64 {
	var result []uint64

	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)

	baseMask := preMask(n, pieceType)

	f := func(blockerSet uint64) {
		if bitops.TestBit(blockerSet, n) {
			result = append(result, blockerSet)
		}
	}

	bitops.Subsets(baseMask, f)

	return result

}

func attackSetForBlockers(n int8, blockers uint64, pieceType int8) uint64 {
	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)
	squares := bitops.FindSetBits(blockers)
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

func genRookPreMask(n int8) uint64 {
	rank, file := bitops.SquareToRankFile(n)
	result := bitops.Mask(n)
	var i int8
	for i = 1; i < 7; i++ {
		result = bitops.SetBit(result, bitops.RankFileToSquare(rank, i))
		result = bitops.SetBit(result, bitops.RankFileToSquare(i, file))
	}
	//bitops.Print(result, 127)
	return result

}

func preMask(n int8, pieceType int8) uint64 {
	b := boardstate.Blank()
	b.SetSquare(n, boardstate.WHITE, pieceType)

	baseMask := bitops.Mask(n)

	if pieceType == boardstate.ROOK {
		return genRookPreMask(n)
	}

	if pieceType == boardstate.BISHOP {
		f1 := func(src, dst int8) {
			baseMask = bitops.SetBit(baseMask, dst)
		}
		// Take all the places a bishop can move
		// Remove the edges
		// Make sure bishop square is still set.
		genSingleBishopMovesGeneric(b, n, f1)
		baseMask = baseMask & bitops.InternalMask()
		baseMask = bitops.SetBit(baseMask, n)
	}

	return baseMask
}

type MagicDefinition struct {
	mapping    []uint64
	magicValue uint64
	rotate     int8
	preMask    uint64
}

func findMagic(n int8, preMask uint64, blockers []uint64, attackSets []uint64) *MagicDefinition {
	blockerMaskBits := bits.OnesCount64(preMask) - 1
	//fmt.Printf("blockerMaskBits = %v\n", blockerMaskBits)
	rotate := int8(64 - blockerMaskBits)
	totalCount := len(blockers)
	best := 0
	var magicValue uint64
	mapping := make([]uint64, 1<<blockerMaskBits)

	for {
		for i := range mapping {
			mapping[i] = 0
		}
		// We want few small bits in our test values.   Using more or fewer random numbers here hurts.
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
		//fmt.Printf("Generating magic %v square for %v\n", pieceType, n)
		blockers := blockersForSquare(n, pieceType)
		attackSets := make([]uint64, len(blockers))
		preMask := preMask(n, pieceType)
		for i, blocker := range blockers {
			attackSets[i] = attackSetForBlockers(n, blocker, pieceType)
		}
		result[n] = findMagic(n, preMask, blockers, attackSets)
	}
	return result
}
