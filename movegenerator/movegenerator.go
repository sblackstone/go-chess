package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func GenSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
  result = append(result, genPawnSuccessors(b)...);
  result = append(result, genKingSuccessors(b)...);
  result = append(result, genQueenSuccessors(b)...);
  result = append(result, genBishopSuccessors(b)...);
  result = append(result, genKnightSuccessors(b)...);
  result = append(result, genRookSuccessors(b)...);
  return result;
}

func GenLegalSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {
	successors := GenSucessors(b)
	return successors;
}

func GenAllCheckedSquares(b *boardstate.BoardState, color int8) []int8 {
	var moves []*boardstate.Move;
  moves = append(moves, genAllPawnMoves(b, color, true)...);
  moves = append(moves, genAllKingMoves(b, color)...);
  moves = append(moves, genAllQueenMoves(b, color)...);
  moves = append(moves, genAllBishopMoves(b, color)...);
  moves = append(moves, genAllKnightMoves(b, color)...);
  moves = append(moves, genAllRookMoves(b, color)...);

	var result []int8
	for j := range(moves) {
		result = append(result, moves[j].Dst())
	}

	return result





}
