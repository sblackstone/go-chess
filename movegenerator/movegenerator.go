package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	//"fmt"
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


// We can do much better here, naieve O(n^2) first attempt.
func GenLegalSuccessors(b *boardstate.BoardState) []*boardstate.BoardState {

	currentTurnColor := b.GetTurn()
	oppTurnColor     := b.EnemyColor()

	// Find all the successors for B...
	successors := GenSuccessors(b)


	var result []*boardstate.BoardState;

	// For each candidate successor
	for i := range(successors) {
		// Find the position of the current turn's king after its move.
		kingPos := successors[i].FindPieces(currentTurnColor, boardstate.KING)

		// Find all the checked squares for the opposing side
		attacks := GenAllCheckedSquares(successors[i], oppTurnColor)
		attacked := false
		for j := range(attacks) {
			if attacks[j] == kingPos[0] {
				attacked = true
				break
			}
		}
		if !attacked {
			result = append(result, successors[i])
		}
	}
	return result;
}

func GenAllCheckedSquares(b *boardstate.BoardState, color int8) []int8 {
	var moves []*boardstate.Move;
  moves = append(moves, genAllPawnMoves(b, color, true)...);
  moves = append(moves, genAllKingMoves(b, color, true)...);
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
