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

type CastlingConfig struct {
	emptyMask  uint64
	attackMask uint64
	kingDst    int8
}

var castlingConfigs [2][2]*CastlingConfig

func initCastlingMasks() {
	castlingConfigs[boardstate.WHITE][boardstate.CASTLE_SHORT] = &CastlingConfig{
		emptyMask:  bitopts.Mask(5) | bitopts.Mask(6),
		attackMask: bitopts.Mask(5) | bitopts.Mask(6),
		kingDst:    6,
	}

	castlingConfigs[boardstate.WHITE][boardstate.CASTLE_LONG] = &CastlingConfig{
		emptyMask:  bitopts.Mask(1) | bitopts.Mask(2) | bitopts.Mask(3),
		attackMask: bitopts.Mask(2) | bitopts.Mask(3),
		kingDst:    2,
	}

	castlingConfigs[boardstate.BLACK][boardstate.CASTLE_SHORT] = &CastlingConfig{
		emptyMask:  bitopts.Mask(61) | bitopts.Mask(62),
		attackMask: bitopts.Mask(61) | bitopts.Mask(62),
		kingDst:    62,
	}

	castlingConfigs[boardstate.BLACK][boardstate.CASTLE_LONG] = &CastlingConfig{
		emptyMask:  bitopts.Mask(57) | bitopts.Mask(58) | bitopts.Mask(59),
		attackMask: bitopts.Mask(58) | bitopts.Mask(59),
		kingDst:    58,
	}
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

func genSingleKingMovesGeneric(b *boardstate.BoardState, kingPos int8, calculateChecks bool, updateFunc func(int8, int8)) {

	kingColor := b.ColorOfSquare(kingPos)

	for _, move := range pregeneratedKingMoves[kingPos] {
		if b.ColorOfSquare(move) != kingColor {
			updateFunc(kingPos, move)
		}
	}

	// If we aren't calculating attacks...
	if !calculateChecks {

		checkedSquares := GenAllCheckedSquares(b, b.EnemyColor())
		occupied := b.GetOccupiedBitboard()

		// And the king isn't in check....
		if !bitopts.TestBit(checkedSquares, kingPos) {

			if b.HasCastleRights(kingColor, boardstate.CASTLE_SHORT) &&
				(occupied&castlingConfigs[kingColor][boardstate.CASTLE_SHORT].emptyMask == 0) &&
				(checkedSquares&castlingConfigs[kingColor][boardstate.CASTLE_SHORT].attackMask == 0) {
				updateFunc(kingPos, castlingConfigs[kingColor][boardstate.CASTLE_SHORT].kingDst)
			}

			if b.HasCastleRights(kingColor, boardstate.CASTLE_LONG) &&
				(occupied&castlingConfigs[kingColor][boardstate.CASTLE_LONG].emptyMask == 0) &&
				(checkedSquares&castlingConfigs[kingColor][boardstate.CASTLE_LONG].attackMask == 0) {
				updateFunc(kingPos, castlingConfigs[kingColor][boardstate.CASTLE_LONG].kingDst)
			}
		}
	}
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

func genAllKingMovesGeneric(b *boardstate.BoardState, color int8, calculateChecks bool, updateFunc func(int8, int8)) {
	kingPositions := b.FindPieces(color, boardstate.KING)

	if len(kingPositions) > 0 {
		for _, kingPos := range kingPositions {
			genSingleKingMovesGeneric(b, kingPos, calculateChecks, updateFunc)
		}
	}

}

func genKingSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	kingPos := b.GetKingPos(b.GetTurn())
	var result []*boardstate.BoardState

	if kingPos == boardstate.NO_KING {
		return result
	}

	updateFunc := func(src, dst int8) {
		result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
	}

	genSingleKingMovesGeneric(b, kingPos, false, updateFunc)

	return result
}
