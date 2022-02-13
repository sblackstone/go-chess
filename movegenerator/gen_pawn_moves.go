package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)



func genSinglePawnMovesWhite(b *boardstate.BoardState, pawnPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	var i uint8
	if b.ColorOfSquare(pawnPos+8) == boardstate.EMPTY {
		newRank := pawnPos / 8
		if (newRank < 7) {
			result = append(result, b.CopyPlayTurn(pawnPos, pawnPos+8))
		} else {
			for i = boardstate.ROOK; i <= boardstate.QUEEN; i++ {
				newBoard := b.CopyPlayTurn(pawnPos, pawnPos + 8)
				newBoard.SetSquare(pawnPos + 8, b.GetTurn(), i)
				result = append(result, newBoard)
			}
		}
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
