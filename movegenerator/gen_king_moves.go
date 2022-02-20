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


func genSingleKingMoves(b *boardstate.BoardState, kingPos int8) []*boardstate.Move {
	var result []*boardstate.Move;
	allKingMoves := pregenerateAllKingMoves(); // TODO: THIS MUST BE MEMOIZED SOMEHOW.
	for i := range(allKingMoves[kingPos]) {
		move := allKingMoves[kingPos][i];
		if b.ColorOfSquare(move) != b.ColorOfSquare(kingPos) {
			result = append(result, boardstate.CreateMove(kingPos, move, boardstate.EMPTY))
		}
	}

	turn := b.ColorOfSquare(kingPos)

  // TODO: NEEDS TO CHECK IF KING MOVES THROUGH CHECK

	if (turn == boardstate.WHITE) {
		if b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(5) && b.EmptySquare(6) {
			result = append(result, boardstate.CreateMove(4, 6, boardstate.EMPTY))
		}
		if b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(1) && b.EmptySquare(2) && b.EmptySquare(3) {
			result = append(result, boardstate.CreateMove(4, 2, boardstate.EMPTY))
		}
	}

	if (turn == boardstate.BLACK) {
		if b.HasCastleRights(turn, boardstate.CASTLE_SHORT) && b.EmptySquare(61) && b.EmptySquare(62) {
			result = append(result, boardstate.CreateMove(60, 62, boardstate.EMPTY))
		}
		if b.HasCastleRights(turn, boardstate.CASTLE_LONG) && b.EmptySquare(57) && b.EmptySquare(58) && b.EmptySquare(59) {
			result = append(result, boardstate.CreateMove(60, 58, boardstate.EMPTY))
		}
	}

	return result
}


/* TODO: CASTLING */
func genAllKingMoves(b *boardstate.BoardState, color int8) []*boardstate.Move {
	var result []*boardstate.Move;
	kingPositions := b.FindPieces(color, boardstate.KING)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(kingPositions); i++ {
		result = append(result, genSingleKingMoves(b, kingPositions[i])...)
	}
  return result;
}


func genKingSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllKingMoves(b, b.GetTurn()))
}
