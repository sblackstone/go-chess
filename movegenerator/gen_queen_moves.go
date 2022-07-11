package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func genSingleQueenMoves(b *boardstate.BoardState, queenPos int8) []*boardstate.Move {
	var result []*boardstate.Move
	result = append(result, genSingleBishopMoves(b, queenPos)...)
	result = append(result, genSingleRookMoves(b, queenPos)...)
	return result
}

func genAllQueenMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move
	queenPositions := b.FindPieces(color, boardstate.QUEEN)
	for _, qp := range queenPositions {
		result = append(result, genSingleQueenMoves(b, qp)...)
	}
	return result
}

func genAllQueenAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	queenPositions := b.FindPieces(color, boardstate.QUEEN)
	for _, qp := range queenPositions {
		result = result | genSingleRookMovesBitboard(b, qp)
		result = result | genSingleBishopMovesBitboard(b, qp)
	}
	return result
}

// func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
// 	return b.GenerateSuccessors(genAllQueenMoves(b, b.GetTurn()))
// }
func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	color := b.GetTurn()
	var result []*boardstate.BoardState
	rookPositions := b.FindPieces(color, boardstate.QUEEN)

	updateFunc := func(src int8, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}

	for _, pos := range rookPositions {
		genSingleBishopMovesGeneric(b, pos, updateFunc)
		genSingleRookMovesGeneric(b, pos, updateFunc)
	}
	return result
}
