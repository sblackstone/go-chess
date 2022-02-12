package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func genAllKingMoves() [64][]uint8 {
  var result [64][]uint8;
  var rank,file uint8
  for rank = 7; rank < 8; rank-- {
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


func genSingleKingMoves(b *boardstate.BoardState, kingPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	allKingMoves := genAllKingMoves(); // TODO: THIS MUST BE MEMOIZED SOMEHOW.
	for i := range(allKingMoves[kingPos]) {
		move := allKingMoves[kingPos][i];
		if b.ColorOfSquare(move) != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(kingPos, move))
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
