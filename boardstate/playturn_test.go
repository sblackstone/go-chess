package boardstate

import (
	"testing"
)

func TestKingCaptured(t *testing.T) {
	b := Blank()
	b.SetSquare(8, BLACK, KING)
	b.SetSquare(16, WHITE, QUEEN)
	oldKingPos := b.GetKingPos(BLACK)

	if oldKingPos != 8 {
		t.Errorf("expected kingPos to be 8")
	}
	before, _ := b.ToFEN()
	b.PlayTurn(16, 8, EMPTY)

	midKingPos := b.GetKingPos(BLACK)
	if midKingPos != NO_KING {
		t.Errorf("Expected NO_KING")
	}
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
	afterKingPos := b.GetKingPos(BLACK)
	if afterKingPos != 8 {
		t.Errorf("Expected afterKingPos to be 8")
	}

}

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

func TestUnplayTurnCastlingWhiteLong(t *testing.T) {
	b := testCastlingBoard()
	before, _ := b.ToFEN()
	b.PlayTurn(4, 2, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnCastlingWhiteShort(t *testing.T) {
	b := testCastlingBoard()
	before, _ := b.ToFEN()
	b.PlayTurn(4, 6, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnCastlingBlackLong(t *testing.T) {
	b := testCastlingBoard()
	b.ToggleTurn()
	before, _ := b.ToFEN()
	b.PlayTurn(60, 58, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnCastlingBlackShort(t *testing.T) {
	b := testCastlingBoard()
	b.ToggleTurn()
	before, _ := b.ToFEN()
	b.PlayTurn(60, 62, EMPTY)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnPromotionWithCaptureWhite(t *testing.T) {
	b := Blank()
	b.SetSquare(48, WHITE, PAWN)
	b.SetSquare(57, BLACK, QUEEN)
	before, _ := b.ToFEN()
	b.PlayTurn(48, 57, KNIGHT)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}
}

func TestUnplayTurnPromotionWithCaptureBlack(t *testing.T) {
	b := Blank()
	b.SetSquare(9, BLACK, PAWN)
	b.SetSquare(0, WHITE, QUEEN)
	b.ToggleTurn()

	before, _ := b.ToFEN()
	b.PlayTurn(9, 0, KNIGHT)
	b.UnplayTurn()
	after, _ := b.ToFEN()
	if before != after {
		t.Errorf("Expected %v to be %v", after, before)
	}

}
