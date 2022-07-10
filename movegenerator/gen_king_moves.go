package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

var pregeneratedKingMoves [64][]int8
var pregeneratedKingMovesBitboard [64]uint64

func getKingMoves() [64][]int8 {
	return pregeneratedKingMoves
}

var castlingTestEmptyMasks [2][2]uint64
var castlingTestAttackMasks [2][2]uint64

func initCastlingMasks() {

	castlingTestEmptyMasks[boardstate.WHITE][boardstate.CASTLE_SHORT] = bitopts.Mask(5) | bitopts.Mask(6)
	castlingTestAttackMasks[boardstate.WHITE][boardstate.CASTLE_SHORT] = bitopts.Mask(5) | bitopts.Mask(6)

	castlingTestEmptyMasks[boardstate.WHITE][boardstate.CASTLE_LONG] = bitopts.Mask(1) | bitopts.Mask(2) | bitopts.Mask(3)
	castlingTestAttackMasks[boardstate.WHITE][boardstate.CASTLE_LONG] = bitopts.Mask(2) | bitopts.Mask(3)

	castlingTestEmptyMasks[boardstate.BLACK][boardstate.CASTLE_SHORT] = bitopts.Mask(61) | bitopts.Mask(62)
	castlingTestAttackMasks[boardstate.BLACK][boardstate.CASTLE_SHORT] = bitopts.Mask(61) | bitopts.Mask(62)

	castlingTestEmptyMasks[boardstate.BLACK][boardstate.CASTLE_LONG] = bitopts.Mask(57) | bitopts.Mask(58) | bitopts.Mask(59)
	castlingTestAttackMasks[boardstate.BLACK][boardstate.CASTLE_LONG] = bitopts.Mask(58) | bitopts.Mask(59)

}

func init() {
	initPregeneratedKingMoves()
	initCastlingMasks()
}

func initPregeneratedKingMoves() {
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
		occupied := b.GetOccupiedBitboard()
		// And the king isn't in check....
		if !contains(checkedSquares, kingPos) {

			if kingColor == boardstate.WHITE {
				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) &&
					(occupied&castlingTestEmptyMasks[boardstate.WHITE][boardstate.CASTLE_SHORT] == 0) &&
					(checkedSquares&castlingTestAttackMasks[boardstate.WHITE][boardstate.CASTLE_SHORT] == 0) {
					updateFunc(6)
				}

				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) &&
					(occupied&castlingTestEmptyMasks[boardstate.WHITE][boardstate.CASTLE_LONG] == 0) &&
					(checkedSquares&castlingTestAttackMasks[boardstate.WHITE][boardstate.CASTLE_LONG] == 0) {
					updateFunc(2)
				}
			}

			if kingColor == boardstate.BLACK {

				if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) &&
					(occupied&castlingTestEmptyMasks[boardstate.BLACK][boardstate.CASTLE_SHORT] == 0) &&
					(checkedSquares&castlingTestAttackMasks[boardstate.BLACK][boardstate.CASTLE_SHORT] == 0) {
					updateFunc(62)
				}

				if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) &&
					(occupied&castlingTestEmptyMasks[boardstate.BLACK][boardstate.CASTLE_LONG] == 0) &&
					(checkedSquares&castlingTestAttackMasks[boardstate.BLACK][boardstate.CASTLE_LONG] == 0) {
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
