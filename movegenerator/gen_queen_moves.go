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

func genAllQueenMoves(b *boardstate.BoardState) []*boardstate.Move {
	var result []*boardstate.Move;
	queenPositions := b.FindPieces(b.GetTurn(), boardstate.QUEEN)
	for i := 0; i < len(queenPositions); i++ {
		result = append(result, genSingleQueenMoves(b, queenPositions[i])...)
	}
	return result
}

func genQueenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllQueenMoves(b))
}
