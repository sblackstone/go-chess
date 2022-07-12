package movegenerator

import (
	"reflect"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func testCastlingBoard() *boardstate.BoardState {
	b := boardstate.Blank()
	b.SetSquare(0, boardstate.WHITE, boardstate.ROOK)
	b.SetSquare(4, boardstate.WHITE, boardstate.KING)
	b.SetSquare(7, boardstate.WHITE, boardstate.ROOK)

	b.SetSquare(56, boardstate.BLACK, boardstate.ROOK)
	b.SetSquare(60, boardstate.BLACK, boardstate.KING)
	b.SetSquare(63, boardstate.BLACK, boardstate.ROOK)
	return b
}

func TestCastleWhite(t *testing.T) {
	b := testCastlingBoard()
	expected := []int8{2, 3, 5, 6, 11, 12, 13}
	testSuccessorsHelper(t, b, boardstate.KING, expected)

	expectedAttacks := []int8{3, 5, 11, 12, 13}
	testAttacksHelper(t, b, boardstate.KING, expectedAttacks)
}

func TestCastleBlack(t *testing.T) {
	b := testCastlingBoard()
	b.SetTurn(boardstate.BLACK)

	expected := []int8{51, 52, 53, 58, 59, 61, 62}
	testSuccessorsHelper(t, b, boardstate.KING, expected)

	expectedAttacks := []int8{51, 52, 53, 59, 61}
	testAttacksHelper(t, b, boardstate.KING, expectedAttacks)

}

func TestCastleWhiteBlocked(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(3, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(5, boardstate.WHITE, boardstate.KNIGHT)
	expected := []int8{11, 12, 13}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestCastleBlackBlocked(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(59, boardstate.BLACK, boardstate.BISHOP)
	b.SetSquare(61, boardstate.BLACK, boardstate.KNIGHT)
	b.SetTurn(boardstate.BLACK)
	expected := []int8{51, 52, 53}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestCastleWhiteBlocked2(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(2, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(6, boardstate.WHITE, boardstate.KNIGHT)
	expected := []int8{3, 5, 11, 12, 13}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestCastleBlackBlocked2(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(58, boardstate.BLACK, boardstate.BISHOP)
	b.SetSquare(62, boardstate.BLACK, boardstate.KNIGHT)
	b.SetTurn(boardstate.BLACK)
	expected := []int8{51, 52, 53, 59, 61}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestCastleWhiteBlocked3(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(5, boardstate.WHITE, boardstate.BISHOP)
	b.SetSquare(1, boardstate.WHITE, boardstate.KNIGHT)
	expected := []int8{3, 11, 12, 13}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestCastleBlackBlocked3(t *testing.T) {
	b := testCastlingBoard()
	b.SetSquare(61, boardstate.BLACK, boardstate.BISHOP)
	b.SetSquare(57, boardstate.BLACK, boardstate.KNIGHT)
	b.SetTurn(boardstate.BLACK)
	expected := []int8{51, 52, 53, 59}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestGenAllKingMoves(t *testing.T) {
	cases := [][]int8{
		{8, 9, 1},
		{9, 8, 10, 0, 2},
		{10, 9, 11, 1, 3},
		{11, 10, 12, 2, 4},
		{12, 11, 13, 3, 5},
		{13, 12, 14, 4, 6},
		{14, 13, 15, 5, 7},
		{15, 14, 6},
		{0, 1, 16, 17, 9},
		{1, 0, 2, 17, 16, 18, 8, 10},
		{2, 1, 3, 18, 17, 19, 9, 11},
		{3, 2, 4, 19, 18, 20, 10, 12},
		{4, 3, 5, 20, 19, 21, 11, 13},
		{5, 4, 6, 21, 20, 22, 12, 14},
		{6, 5, 7, 22, 21, 23, 13, 15},
		{7, 6, 23, 22, 14},
		{8, 9, 24, 25, 17},
		{9, 8, 10, 25, 24, 26, 16, 18},
		{10, 9, 11, 26, 25, 27, 17, 19},
		{11, 10, 12, 27, 26, 28, 18, 20},
		{12, 11, 13, 28, 27, 29, 19, 21},
		{13, 12, 14, 29, 28, 30, 20, 22},
		{14, 13, 15, 30, 29, 31, 21, 23},
		{15, 14, 31, 30, 22},
		{16, 17, 32, 33, 25},
		{17, 16, 18, 33, 32, 34, 24, 26},
		{18, 17, 19, 34, 33, 35, 25, 27},
		{19, 18, 20, 35, 34, 36, 26, 28},
		{20, 19, 21, 36, 35, 37, 27, 29},
		{21, 20, 22, 37, 36, 38, 28, 30},
		{22, 21, 23, 38, 37, 39, 29, 31},
		{23, 22, 39, 38, 30},
		{24, 25, 40, 41, 33},
		{25, 24, 26, 41, 40, 42, 32, 34},
		{26, 25, 27, 42, 41, 43, 33, 35},
		{27, 26, 28, 43, 42, 44, 34, 36},
		{28, 27, 29, 44, 43, 45, 35, 37},
		{29, 28, 30, 45, 44, 46, 36, 38},
		{30, 29, 31, 46, 45, 47, 37, 39},
		{31, 30, 47, 46, 38},
		{32, 33, 48, 49, 41},
		{33, 32, 34, 49, 48, 50, 40, 42},
		{34, 33, 35, 50, 49, 51, 41, 43},
		{35, 34, 36, 51, 50, 52, 42, 44},
		{36, 35, 37, 52, 51, 53, 43, 45},
		{37, 36, 38, 53, 52, 54, 44, 46},
		{38, 37, 39, 54, 53, 55, 45, 47},
		{39, 38, 55, 54, 46},
		{40, 41, 56, 57, 49},
		{41, 40, 42, 57, 56, 58, 48, 50},
		{42, 41, 43, 58, 57, 59, 49, 51},
		{43, 42, 44, 59, 58, 60, 50, 52},
		{44, 43, 45, 60, 59, 61, 51, 53},
		{45, 44, 46, 61, 60, 62, 52, 54},
		{46, 45, 47, 62, 61, 63, 53, 55},
		{47, 46, 63, 62, 54},
		{48, 49, 57},
		{49, 48, 50, 56, 58},
		{50, 49, 51, 57, 59},
		{51, 50, 52, 58, 60},
		{52, 51, 53, 59, 61},
		{53, 52, 54, 60, 62},
		{54, 53, 55, 61, 63},
		{55, 54, 62},
	}

	boards := getKingMoves()
	var i int8

	for i = 0; i < 64; i++ {
		if !reflect.DeepEqual(boards[i], cases[i]) {
			t.Errorf("Expected %v to be %v", boards[i], cases[i])
		}
	}
}

func TestKingCenterOfBoard(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.KING)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)

	expected := []int8{18, 19, 20, 26, 28, 34, 35, 36}
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestKingCanCaptureOtherPieces(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)

	expected := []int8{18, 19, 20, 26, 28, 34, 35, 36}
	for _, expect := range expected {
		b.SetSquare(expect, boardstate.BLACK, boardstate.PAWN)
	}
	b.SetSquare(27, boardstate.WHITE, boardstate.KING)
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestKingBlockedByOwnPieces(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)

	b.SetSquare(18, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(19, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(20, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(26, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(28, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(34, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(35, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(36, boardstate.WHITE, boardstate.PAWN)

	b.SetSquare(27, boardstate.WHITE, boardstate.KING)
	var expected []int8
	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
}

func TestGenKingMovesKnowsAboutTurns(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)

	b.SetSquare(43, boardstate.WHITE, boardstate.KING)
	b.SetSquare(19, boardstate.BLACK, boardstate.KING)

	expected := []int8{34, 35, 36, 42, 44, 50, 51, 52}

	testSuccessorsHelper(t, b, boardstate.KING, expected)
	testAttacksHelper(t, b, boardstate.KING, expected)
	b.ToggleTurn()
	expectedBlack := []int8{10, 11, 12, 18, 20, 26, 27, 28}
	testSuccessorsHelper(t, b, boardstate.KING, expectedBlack)
	testAttacksHelper(t, b, boardstate.KING, expectedBlack)
}

func testLegalKingSuccessors(t *testing.T, b *boardstate.BoardState, expected []int8) {
	successors := GenLegalSuccessors(b)
	locations := genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, successors)
	if !reflect.DeepEqual(locations, expected) {
		t.Errorf("Expected %v to be %v", locations, expected)
	}
}

func TestKingCantMoveIntoCheck(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)

	b.SetSquare(43, boardstate.BLACK, boardstate.KING)
	b.SetSquare(27, boardstate.WHITE, boardstate.KING)

	expected := []int8{18, 19, 20, 26, 28}
	testLegalKingSuccessors(t, b, expected)

	expectedAttacks := []int8{18, 19, 20, 26, 28, 34, 35, 36}
	testAttacksHelper(t, b, boardstate.KING, expectedAttacks)

	b.ToggleTurn()

	expectedBlack := []int8{42, 44, 50, 51, 52}
	testLegalKingSuccessors(t, b, expectedBlack)

	expectedAttacksBlack := []int8{34, 35, 36, 42, 44, 50, 51, 52}
	testAttacksHelper(t, b, boardstate.KING, expectedAttacksBlack)

}

func TestKingHasToMoveOutOfCheck(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)
	b.SetSquare(27, boardstate.WHITE, boardstate.KING)
	b.SetSquare(45, boardstate.BLACK, boardstate.BISHOP)

	expected := []int8{19, 20, 26, 28, 34, 35}
	testLegalKingSuccessors(t, b, expected)
}

func TestKingCantCastleInCheckWhite(t *testing.T) {
	b := testCastlingBoard()
	expected := []int8{2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 11, 12, 13}
	testLegalKingSuccessors(t, b, expected)
	b.SetSquare(22, boardstate.BLACK, boardstate.BISHOP)

	expectedNoLong := []int8{3, 5, 11, 12}
	testLegalKingSuccessors(t, b, expectedNoLong)
}

func TestKingCantCastleInCheckBlack(t *testing.T) {
	b := testCastlingBoard()
	b.SetTurn(boardstate.BLACK)
	expected := []int8{51, 52, 53, 58, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 62}
	testLegalKingSuccessors(t, b, expected)

	b.SetSquare(42, boardstate.WHITE, boardstate.BISHOP)
	expectedNoLong := []int8{52, 53, 59, 61}
	testLegalKingSuccessors(t, b, expectedNoLong)
}

func TestCantCastleThroughCheckWhite(t *testing.T) {
	b := testCastlingBoard()

	// baseline
	expected := []int8{2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 11, 12, 13}
	testLegalKingSuccessors(t, b, expected)

	expected3Attacked := []int8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 11, 12, 13}
	b.SetSquare(10, boardstate.BLACK, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected3Attacked)
	b.SetSquare(10, boardstate.EMPTY, boardstate.EMPTY)

	expected2Attacked := []int8{3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 11, 12, 13}
	b.SetSquare(9, boardstate.BLACK, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected2Attacked)
	b.SetSquare(9, boardstate.EMPTY, boardstate.EMPTY)

	expected5Attacked := []int8{2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 11, 12, 13}
	b.SetSquare(14, boardstate.BLACK, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected5Attacked)
	b.SetSquare(14, boardstate.EMPTY, boardstate.EMPTY)

	expected6Attacked := []int8{2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 11, 12}
	b.SetSquare(41, boardstate.BLACK, boardstate.BISHOP)
	testLegalKingSuccessors(t, b, expected6Attacked)
	b.SetSquare(41, boardstate.EMPTY, boardstate.EMPTY)

}

func TestCantCastleThroughCheckBlack(t *testing.T) {
	b := testCastlingBoard()
	b.SetTurn(boardstate.BLACK)
	expected := []int8{51, 52, 53, 58, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 62}
	testLegalKingSuccessors(t, b, expected)

	expected59Attacked := []int8{51, 52, 53, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 62}
	b.SetSquare(50, boardstate.WHITE, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected59Attacked)
	b.SetSquare(50, boardstate.EMPTY, boardstate.EMPTY)

	expected58Attacked := []int8{51, 52, 53, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 62}
	b.SetSquare(49, boardstate.WHITE, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected58Attacked)
	b.SetSquare(49, boardstate.EMPTY, boardstate.EMPTY)

	expected61Attacked := []int8{51, 52, 53, 58, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60}
	b.SetSquare(54, boardstate.WHITE, boardstate.PAWN)
	testLegalKingSuccessors(t, b, expected61Attacked)
	b.SetSquare(54, boardstate.EMPTY, boardstate.EMPTY)

	expected62Attacked := []int8{51, 52, 58, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61}
	b.SetSquare(17, boardstate.WHITE, boardstate.BISHOP)
	testLegalKingSuccessors(t, b, expected62Attacked)
	b.SetSquare(17, boardstate.EMPTY, boardstate.EMPTY)

}
