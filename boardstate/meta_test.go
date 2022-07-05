package boardstate

import (
	"testing"
	//  "fmt"
)

func TestEnemyColor(t *testing.T) {
	b := Initial()
	if b.EnemyColor() != BLACK {
		t.Errorf("Expected EnemyColor to be Black")
	}

	b.ToggleTurn()

	if b.EnemyColor() != WHITE {
		t.Errorf("Expected EnemyColor to be WHITE")
	}
}

func TestEnpassent(t *testing.T) {
	b := Initial()

	for i := 0; i < 8; i++ {
		if b.IsEnpassant(int8(i)) {
			t.Errorf("No enpassant file should be set from initial state")
		}
	}

	for i := 0; i < 8; i++ {
		b.SetEnpassant(int8(i))
		if b.GetEnpassant() != int8(i) {
			t.Errorf("Expected %v to be %v", b.GetEnpassant(), int8(i))
		}
		for j := 0; j < 8; j++ {
			if j != i && b.IsEnpassant(int8(j)) {
				t.Errorf("%v shouldnt be enpassant", i)
			}
			if j == i && !b.IsEnpassant(int8(j)) {
				t.Errorf("%v should be enpassant", i)
			}
		}
	}

	b.SetEnpassant(int8(1))
	b.ClearEnpassant()
	for i := 0; i < 8; i++ {
		if b.IsEnpassant(int8(i)) {
			t.Errorf("No enpassant file should be set after clear enpassant")
		}
	}
}

func TestCastleRights(t *testing.T) {
	b := Initial()

	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}

	if !b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to have CASTLE_LONG")
	}
	if !b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(WHITE, CASTLE_LONG, false)

	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if !b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(WHITE, CASTLE_SHORT, false)

	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to NOT have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(BLACK, CASTLE_LONG, false)

	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to NOT have CASTLE_SHORT")
	}
	if b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(BLACK, CASTLE_SHORT, false)

	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to NOT have CASTLE_SHORT")
	}
	if b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to NOT have BLACK")
	}

	// Start Reversing things

	b.SetCastleRights(BLACK, CASTLE_SHORT, true)
	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to NOT have CASTLE_SHORT")
	}
	if b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(BLACK, CASTLE_LONG, true)
	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to NOT have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(WHITE, CASTLE_SHORT, true)

	if b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to NOT have CASTLE_LONG")
	}
	if !b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	b.SetCastleRights(WHITE, CASTLE_LONG, true)
	if !b.HasCastleRights(WHITE, CASTLE_LONG) {
		t.Errorf("Expected White to have CASTLE_LONG")
	}
	if !b.HasCastleRights(WHITE, CASTLE_SHORT) {
		t.Errorf("Expected White to have CASTLE_SHORT")
	}
	if !b.HasCastleRights(BLACK, CASTLE_LONG) {
		t.Errorf("Expected BLACK to have CASTLE_LONG")
	}
	if !b.HasCastleRights(BLACK, CASTLE_SHORT) {
		t.Errorf("Expected BLACK to have BLACK")
	}

	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}
}

func TestToggleTurn(t *testing.T) {

	b := Initial()

	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}

	b.ToggleTurn()

	if b.GetTurn() != BLACK {
		t.Errorf("Expected it to be black's turn")
	}

	b.ToggleTurn()

	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}

}

func TestSetGetTurn(t *testing.T) {
	b := Initial()

	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}

	b.SetTurn(BLACK)
	if b.GetTurn() != BLACK {
		t.Errorf("Expected it to be black's turn")
	}

	b.SetTurn(WHITE)
	if b.GetTurn() != WHITE {
		t.Errorf("Expected it to be white's turn")
	}
}
