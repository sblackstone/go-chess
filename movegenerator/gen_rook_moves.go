package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

var rookMagicBitboards [64]*MagicDefinition

func init() {
	rookMagicBitboards = GenerateMagicBitboards(boardstate.ROOK)
}

func genSingleRookMovesGeneric(b *boardstate.BoardState, rookPos int8, updateFunc func(int8, int8, int8)) {
	for r := rookPos + 8; r < 64; r += 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 8; r >= 0; r -= 8 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos + 1; r <= 63 && bitops.FileOfSquare(r) > 0; r += 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := rookPos - 1; r >= 0 && bitops.FileOfSquare(r) < 7; r -= 1 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(rookPos) {
			updateFunc(rookPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

func genAllRookMovesGeneric(b *boardstate.BoardState, color int8, updatefunc func(int8, int8, int8)) {
	b.PieceLocations.EachLocation(color, boardstate.ROOK, func(rookPos int8) {
		genSingleRookMovesGeneric(b, rookPos, updatefunc)
	})
}

func genAllRookAttacks(b *boardstate.BoardState, color int8) uint64 {
	// var result uint64
	// updateFunc := func(src, dst int8) {
	// 	result = bitops.SetBit(result, dst)
	// }
	// genAllRookMovesGeneric(b, color, updateFunc)
	var result uint64
	occupied := b.GetOccupiedBitboard()

	b.PieceLocations.EachLocation(color, boardstate.ROOK, func(rookPos int8) {
		magic := rookMagicBitboards[rookPos]
		// fmt.Printf("PREMASK\n")
		// bitops.Print(magic.preMask, 127)
		blockers := occupied & magic.preMask
		// fmt.Printf("BLOCKERS\n")
		// bitops.Print(blockers, rookPos)
		cacheKey := (blockers * magic.magicValue) >> magic.rotate
		// fmt.Printf("ATTACK MASK\n")
		// bitops.Print(magic.mapping[cacheKey], 127)
		result = result | magic.mapping[cacheKey]
	})
	return result & (^b.GetColorBitboard(color))
}

func genRookSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	updateFunc := func(src, dst, promotionPiece int8) {
		result = append(result, b.CopyPlayTurn(src, dst, promotionPiece))
	}
	genAllRookMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
