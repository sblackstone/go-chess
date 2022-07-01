package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func genSingleQueenMoves(b *boardstate.BoardState, queenPos int8) []*boardstate.Move {
	var result []*boardstate.Move;
	result = append(result, genSingleBishopMoves(b, queenPos)...)
	result = append(result, genSingleRookMoves(b, queenPos)...)
	return result
}

func genAllQueenMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move;
	queenPositions := b.FindPieces(color, boardstate.QUEEN)
	for i := 0; i < len(queenPositions); i++ {
		result = append(result, genSingleQueenMoves(b, queenPositions[i])...)
	}
	return result
}

func genAllQueenAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	queenPositions := b.FindPieces(color, boardstate.QUEEN)
	for i := 0; i < len(queenPositions); i++ {
		result = result | genSingleRookMovesBitboard(b, queenPositions[i])
		result = result | genSingleBishopMovesBitboard(b, queenPositions[i])
	}
	return result
}


func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllQueenMoves(b, b.GetTurn()))
}
