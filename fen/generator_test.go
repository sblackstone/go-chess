package fen

import (
  "testing"
  "github.com/sblackstone/go-chess/boardstate"
)

func TestsMissingGenerator(t *testing.T) {
  t.Errorf("enpassnt on export")
  t.Errorf("turn on export")
  t.Errorf("halfmove on export")
  t.Errorf("fullmove on export")
  t.Errorf("caslting on export")
}

func TestToFEN(t *testing.T) {
  b := boardstate.Initial()
  correctStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  str, _ := ToFEN(b)
  if (str != correctStr) {
    t.Errorf("Expected %v to be %v", str, correctStr)
  }
}
