package boardstate

import (
	"reflect"
	"testing"
)

func TestToFenEnPassant(t *testing.T) {
	b := Blank()
	b.SetSquare(25, BLACK, PAWN)
	b.SetSquare(8, WHITE, PAWN)
	b.PlayTurn(8, 24, EMPTY)
	b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	b.SetCastleRights(BLACK, CASTLE_LONG, false)
	b.SetFullMoves(25)
	b.SetHalfMoves(50)
	boardFen, _ := b.ToFEN(true)
	b2, _ := FromFEN(boardFen)
	if b.GetEnpassant() != b2.GetEnpassant() {
		t.Errorf("Expected FEN to work with enpassant")
	}
	if b.GetTurn() != b2.GetTurn() {
		t.Errorf("Expected FEN Turn")
	}
	if b.GetHalfMoves() != b2.GetHalfMoves() {
		t.Errorf("expected FEN half moves match")
	}
	if b.GetFullMoves() != b2.GetFullMoves() {
		t.Errorf("expected FEN full moves match")
	}

	if !reflect.DeepEqual(b.castleData, b2.castleData) {
		t.Errorf("Expected castling match")
	}

}

func TestToFEN(t *testing.T) {
	b := Initial()
	correctStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	str, _ := b.ToFEN(true)
	if str != correctStr {
		t.Errorf("Expected %v to be %v", str, correctStr)
	}
}

func TestToFENNoClocks(t *testing.T) {
	b := Initial()
	correctStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"
	str, _ := b.ToFEN(false)
	if str != correctStr {
		t.Errorf("Expected %v to be %v", str, correctStr)
	}
}
