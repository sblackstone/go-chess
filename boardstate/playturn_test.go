package boardstate

import (
	"testing"
)

func TestUnplayTurnPieceMove(t *testing.T) {
	b := Initial()
	before, _ := b.ToFEN()
	b.PlayTurn(8, 24, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnCaptureMove(t *testing.T) {
	b := Blank()
	b.SetSquare(27, WHITE, PAWN)
	b.SetSquare(36, BLACK, QUEEN)
	before, _ := b.ToFEN()
	b.PlayTurn(27, 36, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnEnpassantCaptureWhite(t *testing.T) {
	b := Blank()
	b.SetSquare(32, WHITE, PAWN)
	b.SetSquare(49, BLACK, PAWN)
	b.ToggleTurn()
	b.PlayTurn(49, 33, EMPTY)
	b.Print(127)
	before, _ := b.ToFEN()
	b.PlayTurn(32, 41, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnEnpassantCaptureBlack(t *testing.T) {
	b := Blank()
	b.SetSquare(8, WHITE, PAWN)
	b.SetSquare(25, BLACK, PAWN)
	b.PlayTurn(8, 24, EMPTY)
	before, _ := b.ToFEN()
	b.PlayTurn(25, 16, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnEnpassantDeclinedWhite(t *testing.T) {
	b := Blank()
	b.SetSquare(32, WHITE, PAWN)
	b.SetSquare(49, BLACK, PAWN)
	b.SetSquare(15, WHITE, QUEEN)
	b.ToggleTurn()
	b.PlayTurn(49, 33, EMPTY)
	b.Print(127)
	before, _ := b.ToFEN()
	b.PlayTurn(15, 23, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnEnpassantDeclinedBlack(t *testing.T) {
	b := Blank()
	b.SetSquare(8, WHITE, PAWN)
	b.SetSquare(25, BLACK, PAWN)
	b.SetSquare(55, BLACK, QUEEN)
	b.PlayTurn(8, 24, EMPTY)
	before, _ := b.ToFEN()
	b.PlayTurn(55, 47, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnCastling(t *testing.T) {
	t.Errorf("Unimplmemented")
}

func TestUnplayTurnPromotion(t *testing.T) {
	t.Errorf("unimplmeneted")
}
