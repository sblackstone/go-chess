package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
)

func genSinglePawnMoves(b *boardstate.BoardState, pawnPos int8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	pawnPosRank, pawnPosFile := bitopts.SquareToRankFile(pawnPos)
	var promotionRank,pushFoardTwoRank,pushForwardOne,pushForwardTwo,captureToLowerFilePos,captureToHigherFilePos,fromEnpassantRank int8

	if (b.GetTurn() == boardstate.WHITE) {
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
			result = append(result, b.CopyPlayTurn(pawnPos, newPos, boardstate.EMPTY))
		} else {
			// With Promotion
			var i int8
			for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
				newBoard := b.CopyPlayTurn(pawnPos, newPos, i)
				result = append(result, newBoard)
			}
		}
	}

	// Push 1
	if b.EmptySquare(pushForwardOne) {
		appendPawnMovesFn(pushForwardOne)
	}

	// Push 2, never has to promote.
	if bitopts.RankOfSquare(pawnPos) == pushFoardTwoRank && b.EmptySquare(pushForwardOne) && b.EmptySquare(pushForwardTwo) {
		result = append(result, b.CopyPlayTurn(pawnPos, pushForwardTwo, boardstate.EMPTY))
	}

  // Cpature to Lower file
	if (b.EnemyOccupiedSquare(captureToLowerFilePos) && bitopts.FileOfSquare(captureToLowerFilePos) < pawnPosFile) {
		appendPawnMovesFn(captureToLowerFilePos)
	}

	// Capture to Higher file
	if (b.EnemyOccupiedSquare(captureToHigherFilePos) && bitopts.FileOfSquare(captureToHigherFilePos) > pawnPosFile) {
		appendPawnMovesFn(captureToHigherFilePos)
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
	return result
}


func genPawnMoves(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	pawnPositions := b.FindPieces(b.GetTurn(), boardstate.PAWN)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(pawnPositions); i++ {
		result = append(result, genSinglePawnMoves(b, pawnPositions[i])...)
	}

  return result;
}
