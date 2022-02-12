package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func genSingleQueenMoves(b *boardstate.BoardState, queenPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	result = append(result, genSingleRookMoves(b, queenPos)...)
	result = append(result, genSingleBishopMoves(b, queenPos)...)
	return result
}

func genQueenMoves(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
	queenPositions := b.FindPieces(b.GetTurn(), boardstate.QUEEN)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(queenPositions); i++ {
		result = append(result, genSingleQueenMoves(b, queenPositions[i])...)
	}

  return result;
}
