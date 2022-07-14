package boardstate

import "testing"

func TestSquareToAlgebraic(t *testing.T) {
	sq1 := SquareToAlgebraic(26)
	if sq1 != "c4" {
		t.Errorf("Expected %v to be %v", sq1, "c4")
	}

	sq2 := SquareToAlgebraic(0)
	if sq2 != "a1" {
		t.Errorf("Expected %v to be %v", sq2, "a1")
	}

	sq3 := SquareToAlgebraic(63)
	if sq3 != "h8" {
		t.Errorf("Expected %v to be %v", sq3, "h8")
	}

	sq4 := SquareToAlgebraic(53)
	if sq4 != "f7" {
		t.Errorf("Expected %v to be %v", sq4, "f7")
	}

}

func TestAlgebraicToSquare(t *testing.T) {

	val, err := AlgebraicToSquare("c4")
	if err != nil || val != 26 {
		t.Errorf("Expected c4 to be 26, not %v %v", val, err)
	}

	val2, err2 := AlgebraicToSquare("H8")
	if err2 != nil || val2 != 63 {
		t.Errorf("Expected H8 to be 63, not %v %v", val2, err2)
	}

	val3, err3 := AlgebraicToSquare("K1")
	if err3 == nil {
		t.Errorf("Expected k1 to give error, not %v %v", val3, err3)
	}

	val4, err4 := AlgebraicToSquare("a9")
	if err4 == nil {
		t.Errorf("Expected a9 to give error, not %v %v", val4, err4)
	}

	val5, err5 := AlgebraicToSquare("g3")
	if err5 != nil || val5 != 22 {
		t.Errorf("Expected g3 to be 22, not %v %v", val5, err5)
	}

	val6, err6 := AlgebraicToSquare("")
	if err6 == nil {
		t.Errorf("Expected empty string to give error, not %v %v", val6, err6)
	}

	val7, err7 := AlgebraicToSquare("AX")
	if err7 == nil {
		t.Errorf("Expected empty string to give error, not %v %v", val7, err7)
	}

}
