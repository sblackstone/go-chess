package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
)

func genPromotionBoards(b *boardstate.BoardState, src uint8, dst uint8) []*boardstate.BoardState {
	var i  uint8
	var result []*boardstate.BoardState;
	for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
		newBoard := b.CopyPlayTurn(src, dst, i)
		result = append(result, newBoard)
	}
	return result
}


func genSinglePawnMovesWhite(b *boardstate.BoardState, pawnPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	promotionRank := uint8(7)
	pushFoardTwoRank := uint8(1)
	pushForwardOne := pawnPos+8
	pushForwardTwo := pawnPos+16
	captureToLowerFilePos := pawnPos + 7
	captureToHigherFilePos := pawnPos + 9

	appendPawnMovesFn := func(newPos uint8) {
		if (bitopts.RankOfSquare(newPos) != promotionRank) {
			result = append(result, b.CopyPlayTurn(pawnPos, newPos, boardstate.EMPTY))
		} else {
			result = append(result, genPromotionBoards(b, pawnPos, newPos)...)
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
	if (b.EnemyOccupiedSquare(captureToLowerFilePos) && bitopts.FileOfSquare(captureToLowerFilePos) < bitopts.FileOfSquare(pawnPos)) {
		appendPawnMovesFn(captureToLowerFilePos)
	}

	// Capture to Higher file
	if (b.EnemyOccupiedSquare(captureToHigherFilePos) && bitopts.FileOfSquare(captureToHigherFilePos) > bitopts.FileOfSquare(pawnPos)) {
		appendPawnMovesFn(captureToHigherFilePos)
	}



	// TODO: ENPASSANT CAPTURE

	return result
}

func genSinglePawnMovesBlack(b *boardstate.BoardState, pawnPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	return result
}


func genPawnMoves(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	var genFn func(*boardstate.BoardState, uint8) []*boardstate.BoardState;

	if (b.GetTurn() == boardstate.WHITE) {
		genFn = genSinglePawnMovesWhite
	} else {
		genFn = genSinglePawnMovesBlack
	}

	pawnPositions := b.FindPieces(b.GetTurn(), boardstate.PAWN)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(pawnPositions); i++ {
		result = append(result, genFn(b, pawnPositions[i])...)
	}

  return result;
}
