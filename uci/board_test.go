package uci

import (
	"fmt"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func TestRemovePositionPrefix(t *testing.T) {
	if RemovePositionPrefix("position blarg") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

	if RemovePositionPrefix("position           blarg") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

	if RemovePositionPrefix("  blarg  ") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

}

func TestBoardFromUCIPosition(t *testing.T) {
	b1 := boardstate.Initial()
	b1.PlayTurn(8, 24, boardstate.EMPTY)
	b1.PlayTurn(55, 39, boardstate.EMPTY)

	b2 := BoardFromUCIPosition("position startpos moves a2a4 h7h5")

	f1, err := b1.ToFEN()
	if err != nil {
		t.Error(err)
	}

	f2, err := b2.ToFEN()
	if err != nil {
		t.Error(err)
	}

	if f1 != f2 {
		t.Errorf("expected %v to be %v", f2, f1)
	}

}

func TestBoardFromUCIPositionNonStart(t *testing.T) {
	b1 := boardstate.Initial()
	b1.PlayTurn(8, 24, boardstate.EMPTY)
	b1.PlayTurn(55, 39, boardstate.EMPTY)

	f0, err := b1.ToFEN()
	if err != nil {
		t.Error(err)
	}

	posStr := fmt.Sprintf("position %s moves c2c4", f0)
	b1.PlayTurn(10, 26, boardstate.EMPTY)
	b2 := BoardFromUCIPosition(posStr)

	f1, err := b1.ToFEN()

	if err != nil {
		t.Error(err)
	}
	f2, err := b2.ToFEN()
	if err != nil {
		t.Error(err)
	}

	if f1 != f2 {
		t.Errorf("Expected %v to be %v", f2, f1)
	}

}
