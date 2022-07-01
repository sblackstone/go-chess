package movegenerator

import (
	"testing"
	//"fmt"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"

)


func TestGenPawnMovesUnderstandsTurn(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(45, boardstate.BLACK, boardstate.PAWN)

	locations := genSortedBoardLocationsPawns(b)
  expected := []int8{35}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

	b.ToggleTurn()

	locations2 := genSortedBoardLocationsPawns(b)
  expected2 := []int8{37}
  if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2, expected2)
  }
}


func TestPregeneratedPawnAttacks(t *testing.T) {
	t.Errorf("Not implemented")
}
