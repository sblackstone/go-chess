package movegenerator

import (
	"testing"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
//	"github.com/sblackstone/go-chess/fen"
//	"fmt"
)


/* Limited tests since we rely on the rook and bishop code to generate queen moves */


func TestGenQueenMovesUnderstandsTurn(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(56,  boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(7,   boardstate.BLACK, boardstate.QUEEN)
	moves := genSortedBoardLocationsQueens(b)

	expected := []int8{0, 7, 8, 14, 16, 21, 24, 28, 32, 35, 40, 42, 48, 49, 57, 58, 59, 60, 61, 62, 63}
	if (!reflect.DeepEqual(moves, expected)) {
		t.Errorf("Expected %v to be %v", moves, expected)
	}

	b.ToggleTurn()
	movesBlack := genSortedBoardLocationsQueens(b)
	expectedBlack := []int8{0, 1, 2, 3, 4, 5, 6, 14, 15, 21, 23, 28, 31, 35, 39, 42, 47, 49, 55, 56, 63}
	if (!reflect.DeepEqual(movesBlack, expectedBlack)) {
		t.Errorf("Expected %v to be %v", movesBlack, expected)
	}

}


func TestGenQueenMovesMiddleOfBoard(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.QUEEN)
  locations := genSortedBoardLocationsQueens(b)
  expected := []int8{0, 3, 6, 9, 11, 13, 18, 19, 20, 24, 25, 26, 28, 29, 30, 31, 34, 35, 36, 41, 43, 45, 48, 51, 54, 59, 63}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	for i := range(expected) {
		b.SetSquare(expected[i], boardstate.WHITE, boardstate.QUEEN)
	}
}
