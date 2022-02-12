package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"

//	"fmt"
)

func genSingleBishopMoves(b *boardstate.BoardState, bishopPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;

	_, file := bitopts.SquareToRankFile(bishopPos)

	for r := bishopPos+9; r < 64 && r % 8 > file; r += 9 {
			color := b.ColorOfSquare(r)
			if color == boardstate.EMPTY || color != b.GetTurn() {
				result = append(result, b.CopyPlayTurn(bishopPos, r))
			}
			if color != boardstate.EMPTY {
				break;
			}
	}

	for r := bishopPos+7; r < 64 && r % 8 < file; r += 7 {
			color := b.ColorOfSquare(r)
			if color == boardstate.EMPTY || color != b.GetTurn() {
				result = append(result, b.CopyPlayTurn(bishopPos, r))
			}
			if color != boardstate.EMPTY {
				break;
			}
	}


	for r := bishopPos-7; r < 64 && r % 8 > file; r -= 7 {
			color := b.ColorOfSquare(r)
			if color == boardstate.EMPTY || color != b.GetTurn() {
				result = append(result, b.CopyPlayTurn(bishopPos, r))
			}
			if color != boardstate.EMPTY {
				break;
			}
	}

	for r := bishopPos-9; r < 64 && r % 8 < file; r -= 9 {
			color := b.ColorOfSquare(r)
			if color == boardstate.EMPTY || color != b.GetTurn() {
				result = append(result, b.CopyPlayTurn(bishopPos, r))
			}
			if color != boardstate.EMPTY {
				break;
			}
	}




/*
	for r := rookPos-8; r < 64; r -= 8 {
		color := b.ColorOfSquare(r)
		if color == boardstate.EMPTY || color != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(rookPos, r))
		}
		if color != boardstate.EMPTY {
			break;
		}
	}

	for r := rookPos+1; r % 8 > 0; r += 1 {
		color := b.ColorOfSquare(r)
		if color == boardstate.EMPTY || color != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(rookPos, r))
		}
		if color != boardstate.EMPTY {
			break;
		}
	}

	for r := rookPos-1; r % 8 < 7; r -= 1 {
		color := b.ColorOfSquare(r)
		if color == boardstate.EMPTY || color != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(rookPos, r))
		}
		if color != boardstate.EMPTY {
			break;
		}
	}
	*/

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
