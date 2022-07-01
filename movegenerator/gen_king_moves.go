package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)

var pregeneratedKingMoves [64][]int8
var pregeneratedKingMovesBitboard [64]uint64

func getKingMoves() [64][]int8 {
	return pregeneratedKingMoves;
}

func init() {
  var rank,file int8
  for rank = 7; rank >= 0; rank-- {
    for file =0; file < 8; file++ {
      pos := bitopts.RankFileToSquare(rank,file)
			pregeneratedKingMovesBitboard[pos] = 0

			appendPos := func (dst int8) {
				pregeneratedKingMovesBitboard[pos] = bitopts.SetBit(pregeneratedKingMovesBitboard[pos], dst)
				pregeneratedKingMoves[pos]         = append(pregeneratedKingMoves[pos], dst)
			}

			if (rank >= 1) {
				appendPos(pos - 8)
				if (file > 0) {
					appendPos(pos -9)
				}
				if (file < 7) {
					appendPos(pos -7)
				}
			}

			if (rank <= 6) {
				appendPos(pos +8)
				if (file > 0) {
					appendPos(pos +7)
				}
				if (file < 7) {
					appendPos(pos +9)
				}
			}


			if (file > 0) {
				appendPos(pos - 1)
			}
			if (file < 7) {
				appendPos(pos + 1)
			}


		}
  }
}


func contains(arr uint64, value int8) bool {
	return bitopts.TestBit(arr, value);
}

func genSingleKingMovesGeneric(b *boardstate.BoardState, kingPos int8, calculateChecks bool, updateFunc func(int8)) []*boardstate.Move {
	var result []*boardstate.Move;
	for _, move := range(pregeneratedKingMoves[kingPos]) {
		if b.ColorOfSquare(move) != b.ColorOfSquare(kingPos) {
			updateFunc(move)
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
					updateFunc(6)

				}
				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) && b.EmptySquare(1) && b.EmptySquare(2) && b.EmptySquare(3) && !contains(checkedSquares, 2) && !contains(checkedSquares, 3) {
					updateFunc(2)
				}
			}

			if (kingColor == boardstate.BLACK) {
				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) && b.EmptySquare(61) && b.EmptySquare(62) && !contains(checkedSquares, 61) && !contains(checkedSquares, 62) {
					updateFunc(62)
				}
				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) && b.EmptySquare(57) && b.EmptySquare(58) && b.EmptySquare(59) && !contains(checkedSquares, 58) && !contains(checkedSquares, 59) {
					updateFunc(58)
				}
			}
		}
	}

	return result
}



// This will be almost identical everywhere.
func genAllKingAttacks(b *boardstate.BoardState, color int8) uint64 {

	kingPositions := b.FindPieces(color, boardstate.KING)
	if len(kingPositions) == 0 {
			return 0
	} else {
		return (pregeneratedKingMovesBitboard[kingPositions[0]] ^ b.GetColorBitboard(b.ColorOfSquare(kingPositions[0]))) & pregeneratedKingMovesBitboard[kingPositions[0]];
	}

}

// This will be almost identical everywhere.
func genAllKingMoves(b *boardstate.BoardState, color int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move;

	kingPositions := b.FindPieces(color, boardstate.KING)

  if (len(kingPositions) == 0) {
		return result
	}
	updateFunc := func(dst int8) {
		result = append(result, boardstate.CreateMove(kingPositions[0], dst, boardstate.EMPTY))
	}

	genSingleKingMovesGeneric(b, kingPositions[0], calculateChecks, updateFunc)

	return result;
}


func genKingSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllKingMoves(b, b.GetTurn(), false))
}
