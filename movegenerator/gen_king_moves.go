package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func genAllKingMoves() [64][]int8 {
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

			if (rank <= 7) {
				result[pos] = append(result[pos], pos + 8);
				if (file > 0) {
					result[pos] = append(result[pos], pos + 7);
				}
				if (file < 7) {
					result[pos] = append(result[pos], pos + 9);
				}
			}


			if (file >= 1) {
				result[pos] = append(result[pos], pos - 1);
			}
			if (file <= 6) {
				result[pos] = append(result[pos], pos + 1);
			}


		}
  }
  return result;
}


func genSingleKingMoves(b *boardstate.BoardState, kingPos int8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	allKingMoves := genAllKingMoves(); // TODO: THIS MUST BE MEMOIZED SOMEHOW.
	for i := range(allKingMoves[kingPos]) {
		move := allKingMoves[kingPos][i];
		if b.ColorOfSquare(move) != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(kingPos, move, boardstate.EMPTY))
		}
	}

	turn := b.GetTurn()

  // TODO: NEEDS TO CHECK IF KING MOVES THROUGH CHECK!

	if (turn == boardstate.WHITE) {
		if kingPos == 3 && b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(1) && b.EmptySquare(2) {
			result = append(result, b.CopyPlayTurn(3, 1, boardstate.EMPTY))
		}
		if kingPos == 3 && b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(4) && b.EmptySquare(5) && b.EmptySquare(6) {
			result = append(result, b.CopyPlayTurn(3, 5, boardstate.EMPTY))
		}
	}

	if (turn == boardstate.BLACK) {
		if kingPos == 59 && b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(57) && b.EmptySquare(58) {
			result = append(result, b.CopyPlayTurn(59, 57, boardstate.EMPTY))
		}
		if kingPos == 59 && b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(60) && b.EmptySquare(61) && b.EmptySquare(62) {
			result = append(result, b.CopyPlayTurn(59, 61, boardstate.EMPTY))
		}
	}

	return result
}


/* TODO: CASTLING */
func genKingMoves(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	kingPositions := b.FindPieces(b.GetTurn(), boardstate.KING)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(kingPositions); i++ {
		result = append(result, genSingleKingMoves(b, kingPositions[i])...)
	}

  return result;
}
