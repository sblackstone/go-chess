package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
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
		if !bitopts.TestBit(attacks, kingPos[0]) {
			result = append(result, successors[i])
		}
	}
	return result;
}

func GenAllCheckedSquares(b *boardstate.BoardState, color int8) uint64 {
	var result uint64;
	pawnMoves := genAllPawnMoves(b, color, true)
	kingMoves := genAllKingMoves(b, color, true)
	queenMoves := genAllQueenMoves(b, color)
	bishopMoves := genAllBishopMoves(b, color)
	knightMoves := genAllKnightMoves(b, color)
	rookMoves := genAllRookMoves(b, color)

	for _, m := range(pawnMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}


	for _, m := range(kingMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}

	for _, m := range(queenMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}

	for _, m := range(bishopMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}

	for _, m := range(knightMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}

	for _, m := range(rookMoves) {
		result = bitopts.SetBit(result, m.Dst())
	}




	return result;





}
