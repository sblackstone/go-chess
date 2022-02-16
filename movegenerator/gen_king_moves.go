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

			if (rank <= 6) {
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

  // TODO: NEEDS TO CHECK IF KING MOVES THROUGH CHECK

	if (turn == boardstate.WHITE) {
		if b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(5) && b.EmptySquare(6) {
			result = append(result, b.CopyPlayTurn(4, 6, boardstate.EMPTY))
		}
		if b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(1) && b.EmptySquare(2) && b.EmptySquare(3) {
			result = append(result, b.CopyPlayTurn(4, 2, boardstate.EMPTY))
		}
	}

	if (turn == boardstate.BLACK) {
		if b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(61) && b.EmptySquare(62) {
			result = append(result, b.CopyPlayTurn(60, 62, boardstate.EMPTY))
		}
		if b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(57) && b.EmptySquare(58) && b.EmptySquare(59) {
			result = append(result, b.CopyPlayTurn(60, 58, boardstate.EMPTY))
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
