package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
	//"fmt"
)

var pregeneratedKingMoves [64][]int8
var pregeneratedKingMovesBitboard [64]uint64

func getKingMoves() [64][]int8 {
	return pregeneratedKingMoves
}

func init() {
	var rank, file int8
	for rank = 7; rank >= 0; rank-- {
		for file = 0; file < 8; file++ {
			pos := bitopts.RankFileToSquare(rank, file)
			pregeneratedKingMovesBitboard[pos] = 0

			appendPos := func(dst int8) {
				pregeneratedKingMovesBitboard[pos] = bitopts.SetBit(pregeneratedKingMovesBitboard[pos], dst)
				pregeneratedKingMoves[pos] = append(pregeneratedKingMoves[pos], dst)
			}

			if rank >= 1 {
				appendPos(pos - 8)
				if file > 0 {
					appendPos(pos - 9)
				}
				if file < 7 {
					appendPos(pos - 7)
				}
			}

			if rank <= 6 {
				appendPos(pos + 8)
				if file > 0 {
					appendPos(pos + 7)
				}
				if file < 7 {
					appendPos(pos + 9)
				}
			}

			if file > 0 {
				appendPos(pos - 1)
			}
			if file < 7 {
				appendPos(pos + 1)
			}

		}
	}
}

func contains(arr uint64, value int8) bool {
	return bitopts.TestBit(arr, value)
}

func genSingleKingMovesGeneric(b *boardstate.BoardState, kingPos int8, calculateChecks bool, updateFunc func(int8)) []*boardstate.Move {
	var result []*boardstate.Move

	kingColor := b.ColorOfSquare(kingPos)

	for _, move := range pregeneratedKingMoves[kingPos] {
		if b.ColorOfSquare(move) != kingColor {
			updateFunc(move)
		}
	}

	// If we aren't calculating attacks...
	if !calculateChecks {

		checkedSquares := GenAllCheckedSquares(b, b.EnemyColor())

		// And the king isn't in check....
		if !contains(checkedSquares, kingPos) {

			if kingColor == boardstate.WHITE {
				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) && b.EmptySquare(5) && b.EmptySquare(6) && !contains(checkedSquares, 5) && !contains(checkedSquares, 6) {
					updateFunc(6)

				}
				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) && b.EmptySquare(1) && b.EmptySquare(2) && b.EmptySquare(3) && !contains(checkedSquares, 2) && !contains(checkedSquares, 3) {
					updateFunc(2)
				}
			}

			if kingColor == boardstate.BLACK {
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
	kingPos := b.GetKingPos(color)
	if kingPos != boardstate.NO_KING {
		return (pregeneratedKingMovesBitboard[kingPos] ^ b.GetColorBitboard(b.ColorOfSquare(kingPos))) & pregeneratedKingMovesBitboard[kingPos]
	} else {
		return 0
	}
}

// This will be almost identical everywhere.
func genAllKingMoves(b *boardstate.BoardState, color int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move

	kingPositions := b.FindPieces(color, boardstate.KING)

	if len(kingPositions) == 0 {
		return result
	}
	updateFunc := func(dst int8) {
		result = append(result, &boardstate.Move{Src: kingPositions[0], Dst: dst, PromotePiece: boardstate.EMPTY})
	}

	genSingleKingMovesGeneric(b, kingPositions[0], calculateChecks, updateFunc)

	return result
}

func genKingSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	kingPos := b.GetKingPos(b.GetTurn())
	var result []*boardstate.BoardState

	if kingPos == boardstate.NO_KING {
		return result
	}

	updateFunc := func(dst int8) {
		result = append(result, b.CopyPlayTurn(kingPos, dst, boardstate.EMPTY))
	}

	genSingleKingMovesGeneric(b, kingPos, false, updateFunc)

	return result
}
