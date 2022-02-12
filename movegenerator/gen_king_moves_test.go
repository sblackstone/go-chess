package movegenerator

import (
	"testing"
  "reflect"
	//"fmt"
	"github.com/sblackstone/go-chess/boardstate"

)

func TestGenAllKingMoves(t *testing.T) {

	cases := [][]uint8{
		{8,9,1},
		{9,8,10,0,2},
		{10,9,11,1,3},
		{11,10,12,2,4},
		{12,11,13,3,5},
		{13,12,14,4,6},
		{14,13,15,5,7},
		{15,14,6},
		{0,1,16,17,9},
		{1,0,2,17,16,18,8,10},
		{2,1,3,18,17,19,9,11},
		{3,2,4,19,18,20,10,12},
		{4,3,5,20,19,21,11,13},
		{5,4,6,21,20,22,12,14},
		{6,5,7,22,21,23,13,15},
		{7,6,23,22,14},
		{8,9,24,25,17},
		{9,8,10,25,24,26,16,18},
		{10,9,11,26,25,27,17,19},
		{11,10,12,27,26,28,18,20},
		{12,11,13,28,27,29,19,21},
		{13,12,14,29,28,30,20,22},
		{14,13,15,30,29,31,21,23},
		{15,14,31,30,22},
		{16,17,32,33,25},
		{17,16,18,33,32,34,24,26},
		{18,17,19,34,33,35,25,27},
		{19,18,20,35,34,36,26,28},
		{20,19,21,36,35,37,27,29},
		{21,20,22,37,36,38,28,30},
		{22,21,23,38,37,39,29,31},
		{23,22,39,38,30},
		{24,25,40,41,33},
		{25,24,26,41,40,42,32,34},
		{26,25,27,42,41,43,33,35},
		{27,26,28,43,42,44,34,36},
		{28,27,29,44,43,45,35,37},
		{29,28,30,45,44,46,36,38},
		{30,29,31,46,45,47,37,39},
		{31,30,47,46,38},
		{32,33,48,49,41},
		{33,32,34,49,48,50,40,42},
		{34,33,35,50,49,51,41,43},
		{35,34,36,51,50,52,42,44},
		{36,35,37,52,51,53,43,45},
		{37,36,38,53,52,54,44,46},
		{38,37,39,54,53,55,45,47},
		{39,38,55,54,46},
		{40,41,56,57,49},
		{41,40,42,57,56,58,48,50},
		{42,41,43,58,57,59,49,51},
		{43,42,44,59,58,60,50,52},
		{44,43,45,60,59,61,51,53},
		{45,44,46,61,60,62,52,54},
		{46,45,47,62,61,63,53,55},
		{47,46,63,62,54},
		{48,49,64,65,57},
		{49,48,50,65,64,66,56,58},
		{50,49,51,66,65,67,57,59},
		{51,50,52,67,66,68,58,60},
		{52,51,53,68,67,69,59,61},
		{53,52,54,69,68,70,60,62},
		{54,53,55,70,69,71,61,63},
		{55,54,71,70,62},
	}

  boards := genAllKingMoves();
	var i uint8;

	for i = 0; i < 64; i++ {
			if !reflect.DeepEqual(boards[i], cases[i]) {
				t.Errorf("Expected %v to be %v", boards[i], cases[i]);
			}
	}
}



func TestKingCenterOfBoard(t *testing.T) {
	b := boardstate.Blank()
  b.SetSquare(27, boardstate.WHITE, boardstate.KING)
  locations := genSortedBoardLocationsKings(b)
  expected := []uint8{18, 19, 20, 26, 28, 34, 35, 36}
  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}


func TestKingCanCaptureOtherPieces(t *testing.T) {
	b := boardstate.Blank()
	expected := []uint8{18, 19, 20, 26, 28, 34, 35, 36}
	for i := range(expected) {
		b.SetSquare(expected[i], boardstate.BLACK, boardstate.PAWN)
	}
  b.SetSquare(27, boardstate.WHITE, boardstate.KING)
  locations := genSortedBoardLocationsKings(b)


  if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}


func TestKingBlockedByOwnPieces(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(18, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(20, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(26, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(28, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(34, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(35, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)

  b.SetSquare(27, boardstate.WHITE, boardstate.KING)
  locations := genSortedBoardLocationsKings(b)
	var expected []uint8

	if !reflect.DeepEqual(locations, expected) {
    t.Errorf("Expected %v to be %v", locations, expected)
  }
}
