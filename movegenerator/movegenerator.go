package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/bitopts"
	"fmt"
)


func genSuccessorsForPiece(b *boardstate.BoardState, pieceType int8) []*boardstate.BoardState {
		switch pieceType {
		case  boardstate.PAWN:
			return genPawnSuccessors(b)
		case  boardstate.KING:
			return genKingSuccessors(b)
		case  boardstate.QUEEN:
			return genQueenSuccessors(b)
		case  boardstate.BISHOP:
			return genBishopSuccessors(b)
		case  boardstate.KNIGHT:
			return genKnightSuccessors(b)
		case  boardstate.ROOK:
			return genRookSuccessors(b)
		default:
			panic("GenSuccessorsForPiece: Unknown piece type " + fmt.Sprint(pieceType))
		}
}


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
	return genAllKnightAttacks(b, color) |
		     genAllBishopAttacks(b, color) |
				 genAllRookAttacks(b, color)   |
				 genAllQueenAttacks(b, color)  |
				 genAllKingAttacks(b, color)   |
				 genAllPawnAttacks(b, color, true)



}
