package movegenerator

import (
	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

var pregeneratedPawnAttacks [2][64]uint64

func init() {
	var color, pos int8

	updateFunc := func(src, dst int8) {
		pregeneratedPawnAttacks[color][src] = bitopts.SetBit(pregeneratedPawnAttacks[color][src], dst)
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

func genSinglePawnMovesGeneric(b *boardstate.BoardState, pawnPos int8, calculateChecks bool, updateFunc func(int8, int8)) {
	pawnPosRank, pawnPosFile := bitopts.SquareToRankFile(pawnPos)
	var pushFoardTwoRank, pushForwardOne, pushForwardTwo, captureToLowerFilePos, captureToHigherFilePos, fromEnpassantRank int8

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
		if captureToHigherFilePos >= 0 && captureToHigherFilePos <= 63 && bitopts.FileOfSquare(captureToHigherFilePos) > pawnPosFile {
			updateFunc(pawnPos, captureToHigherFilePos)
		}

		// Cpature to Lower file
		if captureToLowerFilePos >= 0 && captureToLowerFilePos <= 63 && bitopts.FileOfSquare(captureToLowerFilePos) < pawnPosFile {
			updateFunc(pawnPos, captureToLowerFilePos)
		}

		return

	}

	// Capture to Higher file
	if captureToHigherFilePos <= 63 && b.EnemyOccupiedSquare(captureToHigherFilePos) && bitopts.FileOfSquare(captureToHigherFilePos) > pawnPosFile {
		updateFunc(pawnPos, captureToHigherFilePos)
	}

	// Cpature to Lower file
	if captureToLowerFilePos >= 0 && b.EnemyOccupiedSquare(captureToLowerFilePos) && bitopts.FileOfSquare(captureToLowerFilePos) < pawnPosFile {
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
	if bitopts.RankOfSquare(pawnPos) == pushFoardTwoRank && b.EmptySquare(pushForwardOne) && b.EmptySquare(pushForwardTwo) {
		updateFunc(pawnPos, pushForwardTwo)
	}

	// Push 1
	if b.EmptySquare(pushForwardOne) {
		updateFunc(pawnPos, pushForwardOne)
	}
}

// This will be almost identical everywhere.
func genSinglePawnMoves(b *boardstate.BoardState, piecePos int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move

	updateFunc := func(src, dst int8) {
		rank := bitopts.RankOfSquare(dst)
		if rank == 0 || rank == 7 {
			// With Promotion
			var i int8
			for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
				result = append(result, &boardstate.Move{Src: src, Dst: dst, PromotePiece: i})
			}
		} else {
			result = append(result, &boardstate.Move{Src: src, Dst: dst, PromotePiece: boardstate.EMPTY})
		}
	}

	genSinglePawnMovesGeneric(b, piecePos, calculateChecks, updateFunc)

	return result
}

func genAllPawnMoves(b *boardstate.BoardState, color int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move
	pawnPositions := b.FindPieces(color, boardstate.PAWN)
	for _, pos := range pawnPositions {
		result = append(result, genSinglePawnMoves(b, pos, calculateChecks)...)
	}
	return result

}

func genAllPawnAttacks(b *boardstate.BoardState, color int8) uint64 {
	var result uint64
	pawnPositions := b.FindPieces(color, boardstate.PAWN)
	for _, pos := range pawnPositions {
		result = result | (pregeneratedPawnAttacks[color][pos]^b.GetColorBitboard(color))&pregeneratedPawnAttacks[color][pos]
	}
	return result
}

func genPawnSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	color := b.GetTurn()
	var result []*boardstate.BoardState
	pawnPositions := b.FindPieces(color, boardstate.PAWN)

	updateFunc := func(src, dst int8) {
		rank := bitopts.RankOfSquare(dst)
		if rank == 0 || rank == 7 {
			var i int8
			for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
				result = append(result, b.CopyPlayTurn(src, dst, i))
			}
		} else {
			result = append(result, b.CopyPlayTurn(src, dst, boardstate.EMPTY))
		}
	}
	for _, pos := range pawnPositions {

		genSinglePawnMovesGeneric(b, pos, false, updateFunc)
	}
	return result
}
