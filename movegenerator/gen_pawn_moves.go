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
	pushForwardOne := pawnPos+8
	pushForwardTwo := pawnPos+16
	if b.EmptySquare(pushForwardOne) {
		if (bitopts.RankOfSquare(pushForwardOne) < 7) {
			result = append(result, b.CopyPlayTurn(pawnPos, pushForwardOne, boardstate.EMPTY))
		} else {
			result = append(result, genPromotionBoards(b, pawnPos, pushForwardOne)...)
		}
	}

	if bitopts.RankOfSquare(pawnPos) == 1 && b.EmptySquare(pushForwardOne) && b.EmptySquare(pushForwardTwo) {
		result = append(result, b.CopyPlayTurn(pawnPos, pushForwardTwo, boardstate.EMPTY))

	}

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
