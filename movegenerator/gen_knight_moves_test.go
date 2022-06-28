package movegenerator

import (
	"testing"
  "reflect"
	"github.com/sblackstone/go-chess/boardstate"
//	"sort"
//	"github.com/sblackstone/go-chess/bitopts"

)



func TestGenKnightMovesKnowsAboutTurns(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(1,  boardstate.WHITE, boardstate.KNIGHT)
	b.SetSquare(57, boardstate.BLACK, boardstate.KNIGHT)
	moves := genSortedBoardLocationsKnights(b)

	expected := []int8{11,16,18}
	if (!reflect.DeepEqual(moves, expected)) {
		t.Errorf("Expected %v to be %v", moves, expected)
	}

	b.ToggleTurn()
	movesBlack := genSortedBoardLocationsKnights(b)
	expectedBlack := []int8{40,42,51}
	if (!reflect.DeepEqual(movesBlack, expectedBlack)) {
		t.Errorf("Expected %v to be %v", movesBlack, expected)
	}
}


func TestGenKnightMovesBlockedBySelf(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(1,  boardstate.WHITE, boardstate.KNIGHT)
	b.SetSquare(57, boardstate.BLACK, boardstate.KNIGHT)
	b.SetSquare(11, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(16, boardstate.WHITE, boardstate.PAWN)

	moves := genSortedBoardLocationsKnights(b)

	expected := []int8{18}
	if (!reflect.DeepEqual(moves, expected)) {
		t.Errorf("Expected %v to be %v", moves, expected)
	}
}

func TestGenKnightMovesCaptures(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(1,  boardstate.WHITE, boardstate.KNIGHT)
	b.SetSquare(57, boardstate.BLACK, boardstate.KNIGHT)
	b.SetSquare(11, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(16, boardstate.BLACK, boardstate.PAWN)

	moves := genSortedBoardLocationsKnights(b)

	expected := []int8{11,16,18}
	if (!reflect.DeepEqual(moves, expected)) {
		t.Errorf("Expected %v to be %v", moves, expected)
	}
}



func TestGenAllKnightMoves(t *testing.T) {

	cases := [][]int8{
	  {17,10},
	  {16,18,11},
	  {8,17,19,12},
	  {9,18,20,13},
	  {10,19,21,14},
	  {11,20,22,15},
	  {12,21,23},
	  {13,22},
	  {25,2,18},
	  {24,26,3,19},
	  {0,16,25,27,4,20},
	  {1,17,26,28,5,21},
	  {2,18,27,29,6,22},
	  {3,19,28,30,7,23},
	  {4,20,29,31},
	  {5,21,30},
	  {1,33,10,26},
	  {0,32,2,34,11,27},
	  {8,24,1,33,3,35,12,28},
	  {9,25,2,34,4,36,13,29},
	  {10,26,3,35,5,37,14,30},
	  {11,27,4,36,6,38,15,31},
	  {12,28,5,37,7,39},
	  {13,29,6,38},
	  {9,41,18,34},
	  {8,40,10,42,19,35},
	  {16,32,9,41,11,43,20,36},
	  {17,33,10,42,12,44,21,37},
	  {18,34,11,43,13,45,22,38},
	  {19,35,12,44,14,46,23,39},
	  {20,36,13,45,15,47},
	  {21,37,14,46},
	  {17,49,26,42},
	  {16,48,18,50,27,43},
	  {24,40,17,49,19,51,28,44},
	  {25,41,18,50,20,52,29,45},
	  {26,42,19,51,21,53,30,46},
	  {27,43,20,52,22,54,31,47},
	  {28,44,21,53,23,55},
	  {29,45,22,54},
	  {25,57,34,50},
	  {24,56,26,58,35,51},
	  {32,48,25,57,27,59,36,52},
	  {33,49,26,58,28,60,37,53},
	  {34,50,27,59,29,61,38,54},
	  {35,51,28,60,30,62,39,55},
	  {36,52,29,61,31,63},
	  {37,53,30,62},
	  {33,42,58},
	  {32,34,43,59},
	  {40,56,33,35,44,60},
	  {41,57,34,36,45,61},
	  {42,58,35,37,46,62},
	  {43,59,36,38,47,63},
	  {44,60,37,39},
	  {45,61,38},
	  {41,50},
	  {40,42,51},
	  {48,41,43,52},
	  {49,42,44,53},
	  {50,43,45,54},
	  {51,44,46,55},
	  {52,45,47},
	  {53,46},
	}

  boards := getPregeneratedKnightMoves();
	var i int8;

	for i = 0; i < 64; i++ {
			if !reflect.DeepEqual(boards[i], cases[i]) {
				t.Errorf("Expected %v to be %v", boards[i], cases[i]);
			}
	}

/*
	// Print out all the knight bitboards
  var pos int8;
	var tmp uint64
	var i int
  for pos = 0; pos < 64; pos++ {
		tmp = 0
		for i = range(boards[pos]) {
			tmp = bitopts.SetBit(tmp, boards[pos][i])
		}
		fmt.Printf("%v: %v\n", pos, boards[pos]);
		bitopts.Print(tmp, pos)
  }
*/
}
