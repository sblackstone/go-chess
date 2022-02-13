package movegenerator

import (
	"testing"
	//"fmt"
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

	// Obstructed by SELF
	b.SetSquare(35, boardstate.WHITE, boardstate.QUEEN)
	var expected2 []uint8
	locations2 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations2, expected2) {
    t.Errorf("Expected %v to be %v", locations2, expected2)
  }

	/// Obstructed by ENEMY
	b.SetSquare(35, boardstate.BLACK, boardstate.QUEEN)
	var expected3 []uint8
	locations3 := genSortedBoardLocationsPawns(b)

  if !reflect.DeepEqual(locations3, expected3) {
    t.Errorf("Expected %v to be %v", locations3, expected3)
  }



}

func TestPushPawnTwoWhite(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(8, boardstate.WHITE, boardstate.PAWN)
  locations := genSortedBoardLocationsPawns(b)
  expected := []uint8{16,24}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}


func TestPushPawnPromoteWhite(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	boards := genPawnMoves(b)
	var sum uint8
	for i := range(boards) {
		sum += boards[i].PieceOfSquare(57)
	}
	if (sum != 6) {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}
