package boardstate

import (
	"testing"
)

func TestMissingGenerator(t *testing.T) {
	t.Errorf("enpassnt on export")
	t.Errorf("turn on export")
	t.Errorf("halfmove on export")
	t.Errorf("fullmove on export")
	t.Errorf("caslting on export")
}

func TestToFEN(t *testing.T) {
	b := Initial()
	correctStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	str, _ := b.ToFEN()
	if str != correctStr {
		t.Errorf("Expected %v to be %v", str, correctStr)
	}
}
