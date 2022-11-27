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

func TestZorbistTurns(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey2 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

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

func TestZorbistPieces(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetSquare(10, WHITE, ROOK)
	zorbistKey2 := b.GetZorbistKey()
	b.SetSquare(10, EMPTY, EMPTY)
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

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

func TestZorbistCastling(t *testing.T) {
	b := Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetCastleRights(WHITE, CASTLE_LONG, false)
	zorbistKey2 := b.GetZorbistKey()
	b.SetCastleRights(WHITE, CASTLE_LONG, true)
	zorbistKey3 := b.GetZorbistKey()

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
