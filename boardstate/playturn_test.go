package boardstate

import "testing"

// import (
// 	"testing"

// 	"github.com/sblackstone/go-chess/fen"
// )

// func TestUnplayTurnPieceMove(t *testing.T) {
// 	b := Initial()
// 	before := fen.ToFEN(b)
// 	b.PlayTurn(8, 24)
// 	b.UnplayTurn()
// 	after := fen.ToFEN(b)
// 	if before != after {
// 		t.Errorf("Expected %v to be %v", after, before)
// 	}
// }

func TestUnplayTurnCaptureMove(t *testing.T) {
	t.Errorf("Unimplmemented")
}

func TestUnplayTurnEnpassantCapture(t *testing.T) {
	t.Errorf("Unimplmemented")
}

func TestUnplayTurnCastling(t *testing.T) {
	t.Errorf("Unimplmemented")
}

func TestUnplayTurnPromotion(t *testing.T) {
	t.Errorf("unimplmeneted")
}
