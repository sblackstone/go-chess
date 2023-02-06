package movegenerator

import (
	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"
)

var pregeneratedPawnAttacks [2][64]uint64

func init() {
	var color, pos int8

	updateFunc := func(src, dst, promotePiece int8) {
		pregeneratedPawnAttacks[color][src] = bitops.SetBit(pregeneratedPawnAttacks[color][src], dst)
	}

	for color = 0; color < 2; color++ {
		for pos = 0; pos < 64; pos++ {
			pregeneratedPawnAttacks[color][pos] = 0
			b := boardstate.Blank()
			b.SetSquare(pos, color, boardstate.PAWN)

			genSinglePawnMovesGeneric(b, pos, true, updateFunc)

		}
	}
}

func genSinglePawnMovesGeneric(b *boardstate.BoardState, pawnPos int8, calculateChecks bool, origUpdateFunc func(int8, int8, int8)) {
	pawnPosRank, pawnPosFile := bitops.SquareToRankFile(pawnPos)
	var pushFoardTwoRank, pushForwardOne, pushForwardTwo, captureToLowerFilePos, captureToHigherFilePos, fromEnpassantRank int8

	updateFunc := func(src, dst int8) {
		rank := bitops.RankOfSquare(dst)
		if rank == 0 || rank == 7 {
			origUpdateFunc(src, dst, boardstate.ROOK)
			origUpdateFunc(src, dst, boardstate.KNIGHT)
			origUpdateFunc(src, dst, boardstate.BISHOP)
			origUpdateFunc(src, dst, boardstate.QUEEN)
		} else {
			origUpdateFunc(src, dst, boardstate.EMPTY)
		}
	}

	if b.ColorOfSquare(pawnPos) == boardstate.WHITE {
		pushFoardTwoRank = int8(1)
		pushForwardOne = pawnPos + 8
		pushForwardTwo = pawnPos + 16
		captureToLowerFilePos = pawnPos + 7
		captureToHigherFilePos = pawnPos + 9
		fromEnpassantRank = int8(4)
	} else {
		pushFoardTwoRank = int8(6)
		pushForwardOne = pawnPos - 8
		pushForwardTwo = pawnPos - 16
		captureToLowerFilePos = pawnPos - 9
		captureToHigherFilePos = pawnPos - 7
		fromEnpassantRank = int8(3)
	}

	if calculateChecks {
		// Capture to Higher file
		if captureToHigherFilePos >= 0 && captureToHigherFilePos <= 63 && bitops.FileOfSquare(captureToHigherFilePos) > pawnPosFile {
			updateFunc(pawnPos, captureToHigherFilePos)
		}

		// Cpature to Lower file
		if captureToLowerFilePos >= 0 && captureToLowerFilePos <= 63 && bitops.FileOfSquare(captureToLowerFilePos) < pawnPosFile {
			updateFunc(pawnPos, captureToLowerFilePos)
		}

		return

	}

	// Capture to Higher file
	if captureToHigherFilePos <= 63 && b.EnemyOccupiedSquare(captureToHigherFilePos) && bitops.FileOfSquare(captureToHigherFilePos) > pawnPosFile {
		updateFunc(pawnPos, captureToHigherFilePos)
	}

	// Cpature to Lower file
	if captureToLowerFilePos >= 0 && b.EnemyOccupiedSquare(captureToLowerFilePos) && bitops.FileOfSquare(captureToLowerFilePos) < pawnPosFile {
		updateFunc(pawnPos, captureToLowerFilePos)
	}

	// The pawn is on the rank where taking enpassant is possible
	if pawnPosRank == fromEnpassantRank {
		enpassantSquare := b.GetEnpassant()
		// The enpassant flag is set.
		if enpassantSquare != boardstate.NO_ENPASSANT {
			// Capture enpassant lower file
			if captureToLowerFilePos == enpassantSquare {
				updateFunc(pawnPos, captureToLowerFilePos)
			}
			if captureToHigherFilePos == enpassantSquare {
				updateFunc(pawnPos, captureToHigherFilePos)
			}
		}
	}

	// Push 2, never has to promote.
	if bitops.RankOfSquare(pawnPos) == pushFoardTwoRank && b.EmptySquare(pushForwardOne) && b.EmptySquare(pushForwardTwo) {
		updateFunc(pawnPos, pushForwardTwo)
	}

	// Push 1
	if b.EmptySquare(pushForwardOne) {
		updateFunc(pawnPos, pushForwardOne)
	}
}

func genAllPawnMovesGeneric(b *boardstate.BoardState, color int8, calculateChecks bool, updateFunc func(int8, int8, int8)) {

	b.PieceLocations.EachLocation(color, boardstate.PAWN, func(pos int8) {
		genSinglePawnMovesGeneric(b, pos, calculateChecks, updateFunc)
	})
}

func genAllPawnAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	b.PieceLocations.EachLocation(color, boardstate.PAWN, func(pos int8) {
		result = result | (pregeneratedPawnAttacks[color][pos]^b.GetColorBitboard(color))&pregeneratedPawnAttacks[color][pos]
	})
	return result
}

func genPawnSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState
	updateFunc := func(src, dst, promotePiece int8) {
		result = append(result, b.CopyPlayTurn(src, dst, promotePiece))
	}
	genAllPawnMovesGeneric(b, b.GetTurn(), false, updateFunc)
	return result
}
