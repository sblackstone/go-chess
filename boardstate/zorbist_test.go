package boardstate

import (
	"fmt"
	"testing"
)

func TestZorbistEnpassant(t *testing.T) {
	b := Initial()
	b.SetEnpassant(5)
	zorbistKey1 := b.GetZorbistKey()
	b.SetEnpassant(6)
	zorbistKey2 := b.GetZorbistKey()
	b.SetEnpassant(5)
	zorbistKey3 := b.GetZorbistKey()
	// fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistTurns(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey2 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPieces(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetSquare(10, WHITE, ROOK)
	zorbistKey2 := b.GetZorbistKey()
	b.SetSquare(10, EMPTY, EMPTY)
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistCastling(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetCastleRights(WHITE, CASTLE_LONG, false)
	zorbistKey2 := b.GetZorbistKey()
	b.SetCastleRights(WHITE, CASTLE_LONG, true)
	zorbistKey3 := b.GetZorbistKey()

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPlayTurnUnplayTurn(t *testing.T) {
	b := Initial()
	zorbistKey1 := b.GetZorbistKey()
	b.PlayTurn(8, 16, EMPTY)
	zorbistKey2 := b.GetZorbistKey()
	b.UnplayTurn()
	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistPlayTurnUnplayTurnCastling(t *testing.T) {
	b := Initial()
	b.PlayTurn(4+8, 4+8+8, EMPTY)
	b.PlayTurn(55, 55-8, EMPTY)
	b.PlayTurn(6, 6+16+1, EMPTY) // Black knight out of the way
	b.PlayTurn(54, 54-8, EMPTY)
	b.PlayTurn(5, 5+7, EMPTY)
	b.PlayTurn(53, 53-8, EMPTY)
	zorbistKey1 := b.GetZorbistKey()
	b.PlayTurn(4, 6, EMPTY) // White castle short
	zorbistKey2 := b.GetZorbistKey()
	b.UnplayTurn()
	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPlayTurnUnplayTurnEnpassant(t *testing.T) {
	b := Initial()
	b.PlayTurn(8, 24, EMPTY)
	b.PlayTurn(55, 39, EMPTY)
	b.PlayTurn(24, 32, EMPTY)

	zorbistKey1 := b.GetZorbistKey()
	b.PrintFen(false)
	// Ready for en-passant push.
	b.PlayTurn(49, 33, EMPTY)
	zorbistKey2 := b.GetZorbistKey()
	b.PrintFen(false)

	// Undo turn.
	b.UnplayTurn()
	b.PrintFen(false)

	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func checkZorbistBeforeAfter(t *testing.T, zorbistKey1, zorbistKey2, zorbistKey3 uint64) {
	if zorbistKey1 == zorbistKey2 {
		t.Errorf("Expected zobristKey1 to not equal zorbistKey2")
	}

	if zorbistKey2 == zorbistKey3 {
		t.Errorf("Expected zobristKey2 to not equal zorbistKey3")
	}

	if zorbistKey3 != zorbistKey1 {
		t.Errorf("Expected zobristKey1 to equal zorbistKey3")
	}
}
