package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func genAllQueenMovesGeneric(b *boardstate.BoardState, color int8, updateFunc func(int8, int8)) {
	queenPositions := b.FindPieces(color, boardstate.QUEEN)
	for _, qp := range queenPositions {
		genSingleRookMovesGeneric(b, qp, updateFunc)
		genSingleBishopMovesGeneric(b, qp, updateFunc)
	}
}

func genAllQueenAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	updateFunc := func(src int8, dst int8) {
		result = bitopts.SetBit(result, dst)
	}
	genAllQueenMovesGeneric(b, color, updateFunc)
	return result
}

func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState

	updateFunc := func(src int8, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}
	genAllQueenMovesGeneric(b, b.GetTurn(), updateFunc)
	return result
}
