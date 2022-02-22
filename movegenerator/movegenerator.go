package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func GenSucessors(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
  result = append(result, genPawnSuccessors(b)...);
  result = append(result, genKingSuccessors(b)...);
  result = append(result, genQueenSuccessors(b)...);
  result = append(result, genBishopSuccessors(b)...);
  result = append(result, genKnightSuccessors(b)...);
  result = append(result, genRookSuccessors(b)...);
  return result;
}

/*
func GenAllMoves(b *boardstate.BoardState) []*boardstate.Move {
  var result []*boardstate.Move;
  result = append(result, genAllPawnMoves(b, b.GetTurn())...);
  result = append(result, genAllKingMoves(b, b.GetTurn())...);
  result = append(result, genAllQueenMoves(b, b.GetTurn())...);
  result = append(result, genAllBishopMoves(b, b.GetTurn())...);
  result = append(result, genAllKnightMoves(b, b.GetTurn())...);
  result = append(result, genAllRookMoves(b, b.GetTurn())...);
  return result;
}
*/

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
