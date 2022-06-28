package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func pregenerateAllKingMoves() [64][]int8 {
  var result [64][]int8;
  var rank,file int8
  for rank = 7; rank >= 0; rank-- {
    for file =0; file < 8; file++ {
      pos := bitopts.RankFileToSquare(rank,file)
			if (rank >= 1) {
				result[pos] = append(result[pos], pos - 8);
				if (file > 0) {
					result[pos] = append(result[pos], pos - 9);
				}
				if (file < 7) {
					result[pos] = append(result[pos], pos - 7);
				}
			}

			if (rank <= 6) {
				result[pos] = append(result[pos], pos + 8);
				if (file > 0) {
					result[pos] = append(result[pos], pos + 7);
				}
				if (file < 7) {
					result[pos] = append(result[pos], pos + 9);
				}
			}


			if (file > 0) {
				result[pos] = append(result[pos], pos - 1);
			}
			if (file < 7) {
				result[pos] = append(result[pos], pos + 1);
			}


		}
  }
  return result;
}


func contains(arr []int8, value int8) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func genSingleKingMoves(b *boardstate.BoardState, kingPos int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move;
	allKingMoves := pregenerateAllKingMoves(); // TODO: THIS MUST BE MEMOIZED SOMEHOW.
	for i := range(allKingMoves[kingPos]) {
		move := allKingMoves[kingPos][i];
		if b.ColorOfSquare(move) != b.ColorOfSquare(kingPos) {
			result = append(result, boardstate.CreateMove(kingPos, move, boardstate.EMPTY))
		}
	}

	kingColor := b.ColorOfSquare(kingPos)


	// If we aren't calculating attacks...
	if (!calculateChecks) {

		checkedSquares := GenAllCheckedSquares(b, b.EnemyColor())


		// And the king isn't in check....
		if !contains(checkedSquares, kingPos) {

			if (kingColor == boardstate.WHITE) {
				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) && b.EmptySquare(5) && b.EmptySquare(6) && !contains(checkedSquares, 5) && !contains(checkedSquares, 6) {
					result = append(result, boardstate.CreateMove(4, 6, boardstate.EMPTY))
				}
				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) && b.EmptySquare(1) && b.EmptySquare(2) && b.EmptySquare(3) && !contains(checkedSquares, 2) && !contains(checkedSquares, 3) {
					result = append(result, boardstate.CreateMove(4, 2, boardstate.EMPTY))
				}
			}

			if (kingColor == boardstate.BLACK) {
				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) && b.EmptySquare(61) && b.EmptySquare(62) && !contains(checkedSquares, 61) && !contains(checkedSquares, 62) {
					result = append(result, boardstate.CreateMove(60, 62, boardstate.EMPTY))
				}
				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) && b.EmptySquare(57) && b.EmptySquare(58) && b.EmptySquare(59) && !contains(checkedSquares, 58) && !contains(checkedSquares, 59) {
					result = append(result, boardstate.CreateMove(60, 58, boardstate.EMPTY))
				}
			}


		}

	}


	return result
}


/* TODO: CASTLING */
func genAllKingMoves(b *boardstate.BoardState, color int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move;
	kingPositions := b.FindPieces(color, boardstate.KING)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(kingPositions); i++ {
		result = append(result, genSingleKingMoves(b, kingPositions[i], calculateChecks)...)
	}
  return result;
}


func genKingSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllKingMoves(b, b.GetTurn(), false))
}
