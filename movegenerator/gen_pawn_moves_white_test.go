package movegenerator

import (
	"reflect"
	"testing"

	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func TestGenerateChecksOnlyWhite(t *testing.T) {
	b := boardstate.Initial()
	squares := GenAllCheckedSquares(b, boardstate.WHITE)
	expected := []int8{16, 17, 18, 19, 20, 21, 22, 23}
	var actual []int8
	var i int8
	for i = 0; i < 64; i++ {
		if bitopts.TestBit(squares, i) {
			actual = append(actual, i)
		}
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestPushPawnWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)

	expected := []int8{35}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	expectedAttacks := []int8{34, 36}
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

	// Obstructed by SELF
	b.SetSquare(35, boardstate.WHITE, boardstate.QUEEN)

	var expectedObstructed []int8
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedObstructed)
	expectedObstructedAttacks := []int8{34, 36}
	testAttacksHelper(t, b, boardstate.PAWN, expectedObstructedAttacks)

	/// Obstructed by ENEMY
	b.SetSquare(35, boardstate.BLACK, boardstate.QUEEN)
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedObstructed)
	testAttacksHelper(t, b, boardstate.PAWN, expectedObstructedAttacks)

}

func TestPushPawnTwoWhite(t *testing.T) {
	// Setup the inital board with a pawn on 8, expect to be pushable 1 or 2 squares.
	b := boardstate.Blank()
	b.SetSquare(8, boardstate.WHITE, boardstate.PAWN)
	expected := []int8{16, 24}
	expectedAttacks := []int8{17}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

	positions := genPawnSuccessors(b)
	for i := range positions {
		if positions[i].PieceOfSquare(24) == boardstate.PAWN {
			if positions[i].GetEnpassant() != 16 {
				t.Errorf("Expected 16 to be enpassant after double push")
			}
		} else {
			if positions[i].GetEnpassant() == 16 {
				t.Errorf("Expected 16 to NOT be enpassant after double push")
			}

		}
	}

	// Test double push behaviors
	var expected2 []int8
	expected2Attacks := []int8{17}
	b.SetSquare(16, boardstate.WHITE, boardstate.QUEEN)
	testSuccessorsHelper(t, b, boardstate.PAWN, expected2)
	testAttacksHelper(t, b, boardstate.PAWN, expected2Attacks)

	b.SetSquare(16, boardstate.EMPTY, boardstate.EMPTY)
	b.SetSquare(24, boardstate.BLACK, boardstate.QUEEN)

	expected3 := []int8{16}
	expected3Attacks := []int8{17}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected3)
	testAttacksHelper(t, b, boardstate.PAWN, expected3Attacks)

}

func TestCaptureHigherFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(28, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{27, 28}
	expectedAttacks := []int8{26, 28}

	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestCaptureLowerFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(26, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{26, 27}
	expectedAttacks := []int8{26, 28}

	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestCaptureHigherFileWhiteWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(61, boardstate.BLACK, boardstate.KING)

	b.SetSquare(62, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		sum += board.PieceOfSquare(62)
	}
	if sum != 6 {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}

func TestCaptureLowerFileWhiteWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(61, boardstate.BLACK, boardstate.KING)

	b.SetSquare(60, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		sum += board.PieceOfSquare(60)
	}
	if sum != 6 {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}

func TestCaptureNoWarpingCapturesHigherFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(23, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(30, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(32, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{30, 31}
	expectedAttacks := []int8{30}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)
}

func TestCaptureNoWarpingCapturesLowerFileWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(16, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(23, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(25, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{24, 25}
	expectedAttacks := []int8{25}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)
}

func TestCaptureNoSelfCapturesWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(36, boardstate.WHITE, boardstate.QUEEN)

	expected := []int8{35}
	var expectedAttacks []int8
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestPushPawnPromoteWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		sum += board.PieceOfSquare(57)
	}
	if sum != 6 {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}

func TestEnPassantCaptureAsWhiteLowerFile(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	expected := []int8{40, 41}
	expectedAttacks := []int8{40, 42}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)
}

func TestEnPassantCaptureAsWhiteHigherFile(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(53, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())
	// Black pushes two setting up enpassant
	b.PlayTurn(53, 37, boardstate.EMPTY)

	expected := []int8{44, 45}
	expectedAttacks := []int8{43, 45}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)
}

func TestEnPassantCaptureAsWhiteUnavailableAfterAdditionalMove(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(48, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(33, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(14, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(55, boardstate.BLACK, boardstate.QUEEN)

	b.SetTurn(boardstate.BLACK)
	//fmt.Println(b.GetEnpassant())

	// Black pushes two setting up enpassant
	b.PlayTurn(48, 32, boardstate.EMPTY)

	expectedInitial := []int8{40, 41}
	expectedInitialAttacks := []int8{40, 42}
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedInitial)
	testAttacksHelper(t, b, boardstate.PAWN, expectedInitialAttacks)

	// White pushes something else`
	b.PlayTurn(14, 22, boardstate.EMPTY)

	// Black pushes something else
	b.PlayTurn(55, 47, boardstate.EMPTY)

	expectedAfter := []int8{41}
	expectedAfterAttacks := []int8{40, 42}
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedAfter)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAfterAttacks)

}
