package movegenerator

import (
	"testing"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
)


/* Limited tests since we rely on the rook and bishop code to generate queen moves */


func TestGenPawnMovesUnderstandsTurn(t *testing.T) {

}


func TestPushPawnWhite(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
  locations := genSortedBoardLocationsPawns(b)
  expected := []uint8{35}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	for i := range(expected) {
		b.SetSquare(expected[i], boardstate.WHITE, boardstate.QUEEN)
	}


}
