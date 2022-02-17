package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"

//	"fmt"
)

func genSingleRookMoves(b *boardstate.BoardState, rookPos int8) []*boardstate.Move {
	var result []*boardstate.Move;
	for r := rookPos+8; r < 64; r += 8 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(rookPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}

	for r := rookPos-8; r >= 0; r -= 8 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(rookPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}

	for r := rookPos+1; bitopts.FileOfSquare(r) > 0; r += 1 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(rookPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}

	for r := rookPos-1; r >= 0 && bitopts.FileOfSquare(r) < 7; r -= 1 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(rookPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}
	return result
}

func genAllRookMoves(b *boardstate.BoardState) []*boardstate.Move {
	var result []*boardstate.Move;
	rookPositions := b.FindPieces(b.GetTurn(), boardstate.ROOK)
	for i := 0; i < len(rookPositions); i++ {
		result = append(result, genSingleRookMoves(b, rookPositions[i])...)
	}
	return result
}

func genRookSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllRookMoves(b))
}
