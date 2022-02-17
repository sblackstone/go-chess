package bitopts

import (
	"testing"
	//  "github.com/sblackstone/go-chess/boardstate"
)


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




}

func TestPrint(t *testing.T) {
	var val uint64 = 5;
	Print(val,52)
}

func TestFindTwoPiecePositions(t *testing.T) {
	  var val uint64;

		val = 0
		posOfZero := FindTwoPiecePositions(val)
		if len(posOfZero) != 0 {
			t.Errorf("Expected %v to be []", posOfZero)
		}

		val = SetBit(val, 1)
		//Print(val,64)

		posOfOne := FindTwoPiecePositions(val)
		if len(posOfOne) != 1 || posOfOne[0] != 1 {
			t.Errorf("Expected %v to be [1]", posOfOne)
		}

		val = SetBit(val, 2)
		//Print(val,64)

		posOfThree := FindTwoPiecePositions(val)
		if len(posOfThree) != 2 || posOfThree[0] != 1 || posOfThree[1] != 2 {
			t.Errorf("Expected %v to be [1,2]", posOfThree)
		}

		val = SetBit(0, 17)
		val = SetBit(val, 19)
		//Print(val,64)
		posOfComplex := FindTwoPiecePositions(val)
		if len(posOfComplex) != 2 || posOfComplex[0] != 17 || posOfComplex[1] != 19 {
			t.Errorf("Expected %v to be [1,2]", posOfComplex)
		}

}

func TestRankFileToSquare(t *testing.T) {
	cases := [][3]int8{
		{0,0,0},
		{7,7,63},
		{7,3,59},
		{3,7,31},
	}

	for i := range(cases) {
		v := RankFileToSquare(cases[i][0],cases[i][1])
		if (v != cases[i][2]) {
				t.Errorf("Expected (%v,%v) to be %v, got %v", cases[i][0],cases[i][1], cases[i][2], v);
		}
	}
}

func TestSquareToRankFile(t *testing.T) {
	cases := [][3]int8{
		{0,0,0},
		{7,7,63},
		{7,3,59},
		{3,7,31},
	}

	for i := range(cases) {
		row,col := SquareToRankFile(cases[i][2])
		if (row != cases[i][0] || col != cases[i][1]) {
				t.Errorf("Expected %v to be (%v,%v), got (%v,%v)", cases[i][2], cases[i][0], cases[i][1], row,col);
		}
	}
}



func TestFlipBit(t *testing.T) {
	var v uint64;

	v = 0b101

	v = FlipBit(v, 1)

	if (v != 0b111) {
		t.Errorf("Expected 101 ^ 010 = 111 (7)")
	}

	v = FlipBit(v, 1)

	if (v != 0b101) {
		t.Errorf("Expected 111 ^ 010 = 101 (5)")
	}

}

func TestSetBit(t *testing.T) {
	var v uint64
	v = SetBit(v, 0)
	if v != 1 {
		t.Errorf("%d should have been 1", v)
	}
	v = SetBit(v, 1)
	if v != 3 {
		t.Errorf("%d should have been 3", v)
	}
}

func TestClearBit(t *testing.T) {
	var v uint64 = 7
	v = ClearBit(v, 0)
	if v != 6 {
		t.Errorf("%d should have been 6", v)
	}
	v = ClearBit(v, 1)
	if v != 4 {
		t.Errorf("%d should have been 4", v)
	}
}

func TestTestBit(t *testing.T) {
	var v uint64 = 5
	if !TestBit(v, 0) {
		t.Errorf("Expected first bit to be true")
	}
	if TestBit(v, 1) {
		t.Errorf("Expected first bit to be false")
	}
	if !TestBit(v, 2) {
		t.Errorf("Expected first bit to be true")
	}

}
