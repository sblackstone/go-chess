package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

var bishopMagicBitboards [64]*MagicDefinition

func init() {
	bishopMagicBitboards = GenerateMagicBitboards(boardstate.BISHOP)
}

func genSingleBishopMovesGeneric(b *boardstate.BoardState, bishopPos int8, updateFunc func(int8, int8, int8)) {
	file := bitops.FileOfSquare(bishopPos)

	for r := bishopPos + 9; r < 64 && bitops.FileOfSquare(r) > file; r += 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos + 7; r < 64 && bitops.FileOfSquare(r) < file; r += 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 7; r >= 0 && bitops.FileOfSquare(r) > file; r -= 7 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}

	for r := bishopPos - 9; r >= 0 && bitops.FileOfSquare(r) < file; r -= 9 {
		if b.ColorOfSquare(r) != b.ColorOfSquare(bishopPos) {
			updateFunc(bishopPos, r, boardstate.EMPTY)
		}
		if !b.EmptySquare(r) {
			break
		}
	}
}

func genAllBishopMovesGeneric(b *boardstate.BoardState, color int8, updateFunc func(int8, int8, int8)) {
	b.PieceLocations.EachLocation(color, boardstate.BISHOP, func(pos int8) {
		genSingleBishopMovesGeneric(b, pos, updateFunc)
	})
}

func genAllBishopAttacks(b *boardstate.BoardState, color int8) uint64 {
	// var result uint64
	// updateFunc := func(src, dst int8) {
	// 	result = bitops.SetBit(result, dst)
	// }
	// genAllRookMovesGeneric(b, color, updateFunc)
	var result uint64

	occupied := b.GetOccupiedBitboard()

	b.PieceLocations.EachLocation(color, boardstate.BISHOP, func(bishopPos int8) {
		// fmt.Printf("BOARD\n")
		// b.Print(127)
		magic := bishopMagicBitboards[bishopPos]
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

func genBishopSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	updateFunc := func(src int8, dst int8, promotionPiece int8) {
		result = append(result, b.CopyPlayTurn(src, dst, promotionPiece))
	}
	genAllBishopMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
