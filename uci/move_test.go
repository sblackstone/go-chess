package uci

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func TestMoveFromUCI(t *testing.T) {
	m, err := MoveFromUCI("a2a4")
	if err != nil {
		t.Error(err)
	}
	if m.Src != 8 || m.Dst != 24 || m.PromotePiece != boardstate.EMPTY {
		t.Errorf("Bad Move: %v", m)
	}
}

func TestMoveFromUCIBadDstSquare(t *testing.T) {
	_, err := MoveFromUCI("a2a9")
	if err == nil {
		t.Errorf("Expected an error from the bad dst square")
	}
}

func TestMoveFromUCIBadSrcSquare(t *testing.T) {
	_, err := MoveFromUCI("a9a4")
	if err == nil {
		t.Errorf("Expected an error from the bad dst square")
	}
}

func TestMoveFromUCISrcDstEqual(t *testing.T) {
	_, err := MoveFromUCI("a4a4")
	if err == nil {
		t.Errorf("Expected an error from the bad dst square")
	}
}

func TestMoveFromUCIPromoteInvalid(t *testing.T) {
	_, err := MoveFromUCI("a7a8H")
	if err == nil {
		t.Errorf("Expected an error from the bad dst square")
	}
}

func TestMoveFromUCIPromotion(t *testing.T) {
	var promoTests = []struct {
		str           string
		expectedPromo int8
	}{
		{"a7a8q", boardstate.QUEEN},
		{"a7a8n", boardstate.KNIGHT},
		{"a7a8r", boardstate.ROOK},
		{"a7a8b", boardstate.BISHOP},
	}

	for _, pt := range promoTests {
		m, err := MoveFromUCI(pt.str)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if m.PromotePiece != pt.expectedPromo {
			t.Errorf("Expected %v to have a %v", m, pt.expectedPromo)
		}
	}
}

func TestMoveToUCI(t *testing.T) {
	var moveTests = []struct {
		move     *boardstate.Move
		expected string
	}{
		{&boardstate.Move{Src: 8, Dst: 16, PromotePiece: boardstate.EMPTY}, "a2a3"},
		{&boardstate.Move{Src: 48, Dst: 56, PromotePiece: boardstate.QUEEN}, "a7a8q"},
		{&boardstate.Move{Src: 48, Dst: 56, PromotePiece: boardstate.KNIGHT}, "a7a8n"},
		{&boardstate.Move{Src: 48, Dst: 56, PromotePiece: boardstate.BISHOP}, "a7a8b"},
		{&boardstate.Move{Src: 48, Dst: 56, PromotePiece: boardstate.ROOK}, "a7a8r"},
	}

	for _, mt := range moveTests {
		result := MoveToUCI(mt.move)
		if result != mt.expected {
			t.Errorf("Expected %v to be %v", result, mt.expected)
		}

	}
}
