package boardstate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoreCastling(t *testing.T) {
	testCase := "8/4k3/8/8/8/8/r6r/R3K2R w KQ - 0 1"
	expected := "8/4k3/8/8/8/8/R6r/4K2R b K - 0 1"
	b, _ := FromFEN(testCase)

	b.PlayTurn(0, 8, EMPTY)

	newFen, _ := b.ToFEN(true)

	if newFen != expected {
		t.Errorf("Initial  %v\n", testCase)
		t.Errorf("Expected %v\n", expected)
		t.Errorf("Actual   %v\n", newFen)

	}

}

func TestAllFileOffsets(t *testing.T) {
	var i int8
	testStr := "P7/1P6/2P5/3P4/4P3/5P2/6P1/7P w - - 0 1"
	b, err := FromFEN(testStr)
	if err != nil {
		t.Errorf("Unexpected error with parsing fen: %v", err)
	}
	for i = 7; i <= 56; i += 7 {
		if b.PieceOfSquare(i) != PAWN || b.ColorOfSquare(i) != WHITE {
			t.Errorf("Expected a pawn on (%v)\n", i)
		}
	}

}

func TestInvalidHalfMoveClock(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c9 ABCD 1"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad half move clock string")
	}
}

func TestInvalidFullMoveClock(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c9 0 ABCD"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad full move clock string")
	}
}

func TestInvalidEnpassantString(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c9 0 1"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad enpassant string")
	}
}

func TestBadTurnString(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR X KQkq - 0 1"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad castling string")
	}
}

func TestBadCharInFenString(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPXPPP/RNBQKBNR w - - 0 1"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad castling string")
	}
}

func TestMalShapedFenString(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBN w - - 0 1"
	_, err := FromFEN(testStr)
	if err == nil {
		t.Errorf("Expected error with bad castling string")
	}
}

func TestParseTurn(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b, err := FromFEN(testStr)
	if err != nil || b.GetTurn() != WHITE {
		t.Errorf("Expected turn to be WHITE after import: %v %v", b.GetTurn(), err)
	}

	testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1"
	b2, err2 := FromFEN(testStr2)
	if err != nil || b2.GetTurn() != BLACK {
		t.Errorf("Expected turn to be WHITE after import: %v %v", b2.GetTurn(), err2)
	}
}

func TestParseEnpassant(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b, err := FromFEN(testStr)
	if err != nil || b.GetEnpassant() != NO_ENPASSANT {
		t.Errorf("Expected GetEnpassant to be NO_ENPASSANT after import: %v %v", b.GetEnpassant(), err)
	}

	testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c4 0 1"
	b2, err2 := FromFEN(testStr2)
	if err2 != nil || b2.GetEnpassant() != 26 {
		t.Errorf("Expected GetEnpassant to be 26 after import: %v %v", b2.GetEnpassant(), err2)
	}
}

func TestParseCastlingNone(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1"
	b, err := FromFEN(testStr)
	if err != nil {
		t.Errorf("Expected valid import: %v", err)
	}
	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("expected not to have  WHITE LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("expected not to have  WHITE SHORT")
	}
	if b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("expected not to have  BLACK LONG")
	}
	if b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("expected not to have  BLACK SHORT")
	}
}

func TestParseCastlingShortOnly(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kk - 0 1"
	b, err := FromFEN(testStr)
	if err != nil {
		t.Errorf("Expected valid import: %v", err)
	}
	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("expected not to have  WHITE LONG")
	}
	if !b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("expected TO have  WHITE SHORT")
	}
	if b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("expected not to have  BLACK LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("expected TO have  BLACK SHORT")
	}
}

func TestParseCastlingLongOnly(t *testing.T) {
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qq - 0 1"
	b, err := FromFEN(testStr)
	if err != nil {
		t.Errorf("Expected valid import: %v", err)
	}
	if !b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("expected TO have  WHITE LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("expected NOT TO have  WHITE SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("expected TO have  BLACK LONG")
	}
	if b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("expected NOT have  BLACK SHORT")
	}
}

func TestFENParserDefaultBoard(t *testing.T) {
	correct := Initial()
	testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b, err := FromFEN(testStr)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if !reflect.DeepEqual(correct, b) {
		t.Errorf("Expected import of default board to match the default board")
	}

}

func TestParseMoveCounters(t *testing.T) {
	testStr := "p6P/1p4P1/2p2P2/3pP3/3Pp3/2P2p2/1P4p1/P6p w - - 25 26"
	b, err := FromFEN(testStr)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if b.GetHalfMoves() != 25 {
		t.Errorf("Expected half moves to be 25")
	}
	if b.GetFullMoves() != 26 {
		t.Errorf("Expected half moves to be 26")
	}

}
