package movegenerator

import (
	"reflect"
	"testing"

	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func TestGenerateChecksOnlyBlack(t *testing.T) {
	b := boardstate.Initial()
	squares := GenAllCheckedSquares(b, boardstate.BLACK)
	expected := []int8{40, 41, 42, 43, 44, 45, 46, 47}
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

func TestPushPawnBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(42, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{34}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	expectedAttacks := []int8{33, 35}
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

	// Obstructed by SELF
	b.SetSquare(34, boardstate.BLACK, boardstate.QUEEN)

	var expectedObstructed []int8
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedObstructed)
	expectedObstructedAttacks := []int8{33, 35}
	testAttacksHelper(t, b, boardstate.PAWN, expectedObstructedAttacks)

	/// Obstructed by ENEMY
	b.SetSquare(34, boardstate.WHITE, boardstate.QUEEN)
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedObstructed)
	testAttacksHelper(t, b, boardstate.PAWN, expectedObstructedAttacks)

}

func TestPushPawnTwoBlack(t *testing.T) {
	// Setup the inital board with a pawn on 8, expect to be pushable 1 or 2 squares.
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(49, boardstate.BLACK, boardstate.PAWN)
	expected := []int8{33, 41}
	expectedAttacks := []int8{40, 42}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

	positions := genPawnSuccessors(b)
	for _, pos := range positions {
		if pos.PieceOfSquare(33) == boardstate.PAWN {
			if pos.GetEnpassant() != 41 {
				t.Errorf("Expected 41 to be enpassant after double push, got %v", pos.GetEnpassant())
			}
		} else {
			if pos.GetEnpassant() == 41 {
				t.Errorf("Expected 41 to NOT be enpassant after double push, got %v", pos.GetEnpassant())
			}

		}
	}

	// Test double-push obstruction behaviors.
	var expected2 []int8
	expected2Attacks := []int8{40, 42}
	b.SetSquare(41, boardstate.WHITE, boardstate.QUEEN)
	testSuccessorsHelper(t, b, boardstate.PAWN, expected2)
	testAttacksHelper(t, b, boardstate.PAWN, expected2Attacks)

	b.SetSquare(41, boardstate.EMPTY, boardstate.EMPTY)
	b.SetSquare(33, boardstate.BLACK, boardstate.QUEEN)

	expected3 := []int8{41}
	expected3Attacks := []int8{40, 42}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected3)
	testAttacksHelper(t, b, boardstate.PAWN, expected3Attacks)

}

func TestCaptureHigherFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(43, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{35, 36}
	expectedAttacks := []int8{34, 36}

	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)
}

func TestCaptureLowerFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(34, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(43, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{34, 35}
	expectedAttacks := []int8{34, 36}

	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestCaptureHigherFileBlackWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(14, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(7, boardstate.WHITE, boardstate.KING)
	b.SetSquare(6, boardstate.WHITE, boardstate.QUEEN)

	//b.SetSquare(62, boardstate.BLACK, boardstate.QUEEN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		//boards[i].Print(255)
		sum += board.PieceOfSquare(7)
	}
	if sum != 6 {
		t.Errorf("Expected square 7 to have rook,knight,bishop or queen, sum was %v", sum)
	}
}

func TestCaptureLowerFileBlackWithPromotion(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(10, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(1, boardstate.WHITE, boardstate.KING)

	b.SetSquare(2, boardstate.WHITE, boardstate.QUEEN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		sum += board.PieceOfSquare(1)
	}
	if sum != 6 {
		t.Errorf("Expected square 6 to have rook,knight,bishop or queen")
	}
}

func TestCaptureNoWarpingCapturesHigherFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(31, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(22, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(24, boardstate.WHITE, boardstate.PAWN)

	expected := []int8{22, 23}
	expectedAttacks := []int8{22}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestCaptureNoWarpingCapturesLowerFileBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(24, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(17, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(15, boardstate.WHITE, boardstate.PAWN)

	expected := []int8{16, 17}
	expectedAttacks := []int8{17}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestCaptureNoSelfCapturesBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(36, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(27, boardstate.BLACK, boardstate.QUEEN)
	b.SetSquare(29, boardstate.BLACK, boardstate.QUEEN)

	expected := []int8{28}
	var expectedAttacks []int8
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestPushPawnPromoteBlack(t *testing.T) {
	b := boardstate.Blank()
	b.ToggleTurn()
	b.SetSquare(11, boardstate.BLACK, boardstate.PAWN)
	boards := genPawnSuccessors(b)
	var sum int8
	for _, board := range boards {
		sum += board.PieceOfSquare(3)
	}
	if sum != 6 {
		t.Errorf("Expected square 3 to have rook,knight,bishop or queen")
	}
}

func TestEnPassantCaptureAsBlackLowerFile(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(26, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(9, boardstate.WHITE, boardstate.PAWN)

	// White pushes two setting up enpassant
	b.PlayTurn(9, 25, boardstate.EMPTY)

	expected := []int8{17, 18}
	expectedAttacks := []int8{17, 19}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestEnPassantCaptureAsBlackHigherFile(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(29, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(14, boardstate.WHITE, boardstate.PAWN)

	// White pushes two setting up enpassant
	b.PlayTurn(14, 30, boardstate.EMPTY)

	expected := []int8{21, 22}
	expectedAttacks := []int8{20, 22}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

}

func TestEnPassantCaptureAsBlackUnavailableAfterAdditionalMove(t *testing.T) {

	b := boardstate.Blank()
	b.SetSquare(30, boardstate.BLACK, boardstate.PAWN)
	b.SetSquare(13, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(9, boardstate.WHITE, boardstate.QUEEN)
	b.SetSquare(57, boardstate.BLACK, boardstate.QUEEN)

	// White pushes two setting up enpassant
	b.PlayTurn(13, 29, boardstate.EMPTY)

	expectedInitial := []int8{21, 22}
	expectedInitialAttacks := []int8{21, 23}
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedInitial)
	testAttacksHelper(t, b, boardstate.PAWN, expectedInitialAttacks)

	// Block pushes something else`
	b.PlayTurn(57, 58, boardstate.EMPTY)

	// White pushes something else
	b.PlayTurn(9, 10, boardstate.EMPTY)

	expectedAfter := []int8{22}
	expectedAfterAttacks := []int8{21, 23}
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedAfter)
	testAttacksHelper(t, b, boardstate.PAWN, expectedAfterAttacks)

}
