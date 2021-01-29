package boardstate

import (
	"testing"
	//  "github.com/sblackstone/go-chess/boardstate"
)

func TestSetBit(t *testing.T) {
	var v uint64
	v = setBit(v, 0)
	if v != 1 {
		t.Errorf("%d should have been 1", v)
	}
	v = setBit(v, 1)
	if v != 3 {
		t.Errorf("%d should have been 3", v)
	}
}

func TestClearBit(t *testing.T) {
	var v uint64 = 7
	v = clearBit(v, 0)
	if v != 6 {
		t.Errorf("%d should have been 6", v)
	}
	v = clearBit(v, 1)
	if v != 4 {
		t.Errorf("%d should have been 4", v)
	}
}

func TestTestBit(t *testing.T) {
	var v uint64 = 5
	if !testBit(v, 0) {
		t.Errorf("Expected first bit to be true")
	}
	if testBit(v, 1) {
		t.Errorf("Expected first bit to be false")
	}
	if !testBit(v, 2) {
		t.Errorf("Expected first bit to be true")
	}

}
