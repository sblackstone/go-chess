package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
)

func genSinglePawnMoves(b *boardstate.BoardState, pawnPos int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move;
	pawnPosRank, pawnPosFile := bitopts.SquareToRankFile(pawnPos)
	var promotionRank,pushFoardTwoRank,pushForwardOne,pushForwardTwo,captureToLowerFilePos,captureToHigherFilePos,fromEnpassantRank int8

	if (b.ColorOfSquare(pawnPos) == boardstate.WHITE) {
		promotionRank          = int8(7)
		pushFoardTwoRank       = int8(1)
		pushForwardOne         = pawnPos + 8
		pushForwardTwo         = pawnPos + 16
		captureToLowerFilePos  = pawnPos + 7
		captureToHigherFilePos = pawnPos + 9
		fromEnpassantRank      = int8(4)
	} else {
		promotionRank          = int8(0)
		pushFoardTwoRank       = int8(6)
		pushForwardOne         = pawnPos - 8
		pushForwardTwo         = pawnPos - 16
		captureToLowerFilePos  = pawnPos - 9
		captureToHigherFilePos = pawnPos - 7
		fromEnpassantRank      = int8(3)
	}

	appendPawnMovesFn := func(newPos int8) {
		// Non-Promotion
		if (bitopts.RankOfSquare(newPos) != promotionRank) {
			result = append(result, boardstate.CreateMove(pawnPos, newPos, boardstate.EMPTY))
		} else {
			// With Promotion
			var i int8
			for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
				newBoard := boardstate.CreateMove(pawnPos, newPos, i)
				result = append(result, newBoard)
			}
		}
	}

	if (calculateChecks) {
		// Capture to Higher file
		if (captureToHigherFilePos >= 0 && captureToHigherFilePos <= 63 && bitopts.FileOfSquare(captureToHigherFilePos) > pawnPosFile) {
			appendPawnMovesFn(captureToHigherFilePos)
		}

		// Cpature to Lower file
		if (captureToLowerFilePos >= 0 && captureToLowerFilePos <= 64 && bitopts.FileOfSquare(captureToLowerFilePos) < pawnPosFile) {
			appendPawnMovesFn(captureToLowerFilePos)
		}

		return result


	}


	// Capture to Higher file
	if (captureToHigherFilePos <= 63 && b.EnemyOccupiedSquare(captureToHigherFilePos) && bitopts.FileOfSquare(captureToHigherFilePos) > pawnPosFile) {
		appendPawnMovesFn(captureToHigherFilePos)
	}

  // Cpature to Lower file
	if (captureToLowerFilePos >= 0 && b.EnemyOccupiedSquare(captureToLowerFilePos) && bitopts.FileOfSquare(captureToLowerFilePos) < pawnPosFile) {
		appendPawnMovesFn(captureToLowerFilePos)
	}

	// The pawn is on the rank where taking enpassant is possible
	if pawnPosRank == fromEnpassantRank {
		enpassantSquare := b.GetEnpassant()
		// The enpassant flag is set.
		if enpassantSquare != boardstate.NO_ENPASSANT {
			// Capture enpassant lower file
			if captureToLowerFilePos == enpassantSquare {
				appendPawnMovesFn(captureToLowerFilePos)
			}
			if captureToHigherFilePos == enpassantSquare {
				appendPawnMovesFn(captureToHigherFilePos)
			}
		}
	}

	// Push 2, never has to promote.
	if bitopts.RankOfSquare(pawnPos) == pushFoardTwoRank && b.EmptySquare(pushForwardOne) && b.EmptySquare(pushForwardTwo) {
		result = append(result, boardstate.CreateMove(pawnPos, pushForwardTwo, boardstate.EMPTY))
	}

	// Push 1
	if b.EmptySquare(pushForwardOne) {
		appendPawnMovesFn(pushForwardOne)
	}


	return result
}

func genAllPawnMoves(b *boardstate.BoardState, color int8, calculateChecks bool) []*boardstate.Move {
	var result []*boardstate.Move;
	pawnPositions := b.FindPieces(color, boardstate.PAWN)
	for i := 0; i < len(pawnPositions); i++ {
		result = append(result, genSinglePawnMoves(b, pawnPositions[i], calculateChecks)...)
	}
  return result;

}


func genPawnSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	return b.GenerateSuccessors(genAllPawnMoves(b, b.GetTurn(), false))
}
