package movegenerator

import (
	"testing"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
//	"fmt"
//	"sort"

//  "fmt"
//	"github.com/sblackstone/go-chess/bitopts"

)


func TestGenQueenMovesUnderstandsTurn(t *testing.T) {
	t.Errorf("TODO\n")
}




func TestGenQueenMovesMiddleOfBoard(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.QUEEN)
  locations := genSortedBoardLocationsQueens(b)
  expected := []uint8{0, 3, 6, 9, 11, 13, 18, 19, 20, 24, 25, 26, 28, 29, 30, 31, 34, 35, 36, 41, 43, 45, 48, 51, 54, 59, 63}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	for i := range(expected) {
		b.SetSquare(expected[i], boardstate.WHITE, boardstate.QUEEN)
	}


}
