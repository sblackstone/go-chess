package movegenerator

import (
	"testing"
  "reflect"
  "github.com/sblackstone/go-chess/boardstate"
//	"sort"

//  "fmt"
//	"github.com/sblackstone/go-chess/bitopts"

)


func TestGenRookMovesUnderstandsTurn(t *testing.T) {
  b := boardstate.Blank()

  b.SetSquare(56, boardstate.WHITE, boardstate.ROOK)
  b.SetSquare(7,  boardstate.BLACK, boardstate.ROOK)

  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{0,8,16,24,32,40,48,57,58,59,60,61,62,63}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

  b.ToggleTurn()
  locationsBlack := genSortedBoardLocationsRooks(b)
  expectedBlack := []uint8{0,1,2,3,4,5,6,15,23,31,39,47,55,63}
  if !reflect.DeepEqual(locationsBlack, expectedBlack) {
    t.Errorf("Expected %v to be %v", locationsBlack, expectedBlack)
  }


}

func TestGenRookNWCorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(56, boardstate.WHITE, boardstate.ROOK)
  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{0,8,16,24,32,40,48,57,58,59,60,61,62,63}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenRookSECorner(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(7, boardstate.WHITE, boardstate.ROOK)
  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{0,1,2,3,4,5,6,15,23,31,39,47,55,63}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenRookBlockedBySelf(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
  b.SetSquare(25, boardstate.WHITE, boardstate.PAWN)
  b.SetSquare(30, boardstate.WHITE, boardstate.QUEEN)
  b.SetSquare(51, boardstate.WHITE, boardstate.KNIGHT)
  b.SetSquare(11, boardstate.WHITE, boardstate.BISHOP)
  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{19, 26, 28, 29, 35, 43}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}

func TestGenRookStopsAtCapture(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
  b.SetSquare(25, boardstate.BLACK, boardstate.PAWN)
  b.SetSquare(30, boardstate.BLACK, boardstate.QUEEN)
  b.SetSquare(51, boardstate.BLACK, boardstate.KNIGHT)
  b.SetSquare(11, boardstate.BLACK, boardstate.BISHOP)
  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{11, 19, 25, 26, 28, 29, 30, 35, 43, 51}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

}

func TestTwoRooksOnBoard(t *testing.T) {
}



func TestGenRookMovesMiddleOfBoard(t *testing.T) {
  b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.ROOK)
  locations := genSortedBoardLocationsRooks(b)
  expected := []uint8{3, 11, 19, 24, 25, 26, 28, 29, 30, 31, 35, 43, 51, 59}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }

}
