package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"

//	"fmt"
)

func genSingleBishopMoves(b *boardstate.BoardState, bishopPos int8) []*boardstate.Move {
	var result []*boardstate.Move;

	file := bitopts.FileOfSquare(bishopPos)

	for r := bishopPos+9; r < 64 && bitopts.FileOfSquare(r) > file; r += 9 {
			if b.EmptyOrEnemyOccupiedSquare(r) {
				result = append(result, boardstate.CreateMove(bishopPos, r, boardstate.EMPTY))
			}
			if !b.EmptySquare(r) {
				break;
			}
	}

	for r := bishopPos+7; r < 64 && bitopts.FileOfSquare(r) < file; r += 7 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}


	for r := bishopPos-7; r >= 0 && bitopts.FileOfSquare(r) > file; r -= 7 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}

	for r := bishopPos-9; r >= 0 && bitopts.FileOfSquare(r) < file; r -= 9 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, boardstate.CreateMove(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}
	return result
}


func genAllBishopMoves(b *boardstate.BoardState) []*boardstate.Move {
	var result []*boardstate.Move;
	rookPositions := b.FindPieces(b.GetTurn(), boardstate.BISHOP)
	for i := 0; i < len(rookPositions); i++ {
		result = append(result, genSingleBishopMoves(b, rookPositions[i])...)
	}
	return result
}

func genBishopSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllBishopMoves(b))
}
