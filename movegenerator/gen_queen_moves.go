package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func genAllQueenMovesGeneric(b *boardstate.BoardState, color int8, updateFunc func(int8, int8, int8)) {
	b.PieceLocations.EachLocation(color, boardstate.QUEEN, func(pos int8) {
		genSingleRookMovesGeneric(b, pos, updateFunc)
		genSingleBishopMovesGeneric(b, pos, updateFunc)
	})
}

func genAllQueenAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	occupied := b.GetOccupiedBitboard()

	b.PieceLocations.EachLocation(color, boardstate.QUEEN, func(queenPos int8) {
		magic := rookMagicBitboards[queenPos]
		blockers := occupied & magic.preMask
		cacheKey := (blockers * magic.magicValue) >> magic.rotate
		result = result | magic.mapping[cacheKey]

		magic = bishopMagicBitboards[queenPos]
		blockers = occupied & magic.preMask
		cacheKey = (blockers * magic.magicValue) >> magic.rotate
		result = result | magic.mapping[cacheKey]
	})

	return result & (^b.GetColorBitboard(color))
}

func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState

	updateFunc := func(src, dst, promotionPiece int8) {
		result = append(result, b.CopyPlayTurn(src, dst, promotionPiece))
	}
	genAllQueenMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
