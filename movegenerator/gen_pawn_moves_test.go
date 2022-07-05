package movegenerator

import (
	"testing"
	//"fmt"
	//"reflect"
	"github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"
)

func TestGenPawnMovesUnderstandsTurn(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(27, boardstate.WHITE, boardstate.PAWN)
	b.SetSquare(45, boardstate.BLACK, boardstate.PAWN)

	expected := []int8{35}
	testSuccessorsHelper(t, b, boardstate.PAWN, expected)
	expectedAttacks := []int8{34, 36}
	testAttacksHelper(t, b, boardstate.PAWN, expectedAttacks)

	b.ToggleTurn()

	expectedBlack := []int8{37}
	testSuccessorsHelper(t, b, boardstate.PAWN, expectedBlack)
	expectedBlackAttacks := []int8{36, 38}
	testAttacksHelper(t, b, boardstate.PAWN, expectedBlackAttacks)
}

// func TestPregeneratedPawnAttacks(t *testing.T) {
// 	t.Errorf("Not implemented")
// }
