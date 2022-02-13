package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"

//	"fmt"
)

func genSingleBishopMoves(b *boardstate.BoardState, bishopPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;

	file := bitopts.FileOfSquare(bishopPos)

	for r := bishopPos+9; r < 64 && bitopts.FileOfSquare(r) > file; r += 9 {
			if b.EmptyOrEnemyOccupiedSquare(r) {
				result = append(result, b.CopyPlayTurn(bishopPos, r, boardstate.EMPTY))
			}
			if !b.EmptySquare(r) {
				break;
			}
	}

	for r := bishopPos+7; r < 64 && bitopts.FileOfSquare(r) < file; r += 7 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, b.CopyPlayTurn(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}


	for r := bishopPos-7; r < 64 && bitopts.FileOfSquare(r) > file; r -= 7 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, b.CopyPlayTurn(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}

	for r := bishopPos-9; r < 64 && bitopts.FileOfSquare(r) < file; r -= 9 {
		if b.EmptyOrEnemyOccupiedSquare(r) {
			result = append(result, b.CopyPlayTurn(bishopPos, r, boardstate.EMPTY))
		}
		if !b.EmptySquare(r) {
			break;
		}
	}
	return result
}

func genBishopMoves(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
	bishopPositions := b.FindPieces(b.GetTurn(), boardstate.BISHOP)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(bishopPositions); i++ {
		result = append(result, genSingleBishopMoves(b, bishopPositions[i])...)
	}

  return result;
}
