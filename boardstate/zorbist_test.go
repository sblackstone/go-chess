package boardstate_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
	"github.com/sblackstone/go-chess/movegenerator"
	"github.com/stretchr/testify/require"
)

func TestZorbistEnpassant(t *testing.T) {
	b := boardstate.Initial()
	b.SetEnpassant(5)
	zorbistKey1 := b.GetZorbistKey()
	b.SetEnpassant(6)
	zorbistKey2 := b.GetZorbistKey()
	b.SetEnpassant(5)
	zorbistKey3 := b.GetZorbistKey()
	// fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistTurns(t *testing.T) {
	b := boardstate.Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey2 := b.GetZorbistKey()
	b.ToggleTurn()
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPieces(t *testing.T) {
	b := boardstate.Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetSquare(10, boardstate.WHITE, boardstate.ROOK)
	zorbistKey2 := b.GetZorbistKey()
	b.SetSquare(10, boardstate.EMPTY, boardstate.EMPTY)
	zorbistKey3 := b.GetZorbistKey()
	fmt.Printf("%+v %+v %+v\n", zorbistKey1, zorbistKey2, zorbistKey3)

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistCastling(t *testing.T) {
	b := boardstate.Blank()
	zorbistKey1 := b.GetZorbistKey()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	zorbistKey2 := b.GetZorbistKey()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, true)
	zorbistKey3 := b.GetZorbistKey()

	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPlayTurnUnplayTurn(t *testing.T) {
	b := boardstate.Initial()
	zorbistKey1 := b.GetZorbistKey()
	b.PlayTurn(8, 16, boardstate.EMPTY)
	zorbistKey2 := b.GetZorbistKey()
	b.UnplayTurn()
	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func TestZorbistPlayTurnUnplayTurnCastling(t *testing.T) {
	b := boardstate.Initial()
	b.PlayTurn(4+8, 4+8+8, boardstate.EMPTY)
	b.PlayTurn(55, 55-8, boardstate.EMPTY)
	b.PlayTurn(6, 6+16+1, boardstate.EMPTY) // Black knight out of the way
	b.PlayTurn(54, 54-8, boardstate.EMPTY)
	b.PlayTurn(5, 5+7, boardstate.EMPTY)
	b.PlayTurn(53, 53-8, boardstate.EMPTY)
	zorbistKey1 := b.GetZorbistKey()
	b.PlayTurn(4, 6, boardstate.EMPTY) // White castle short
	zorbistKey2 := b.GetZorbistKey()
	b.UnplayTurn()
	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)

}

func TestZorbistPlayTurnUnplayTurnEnpassant(t *testing.T) {
	b := boardstate.Initial()
	b.PlayTurn(8, 24, boardstate.EMPTY)
	b.PlayTurn(55, 39, boardstate.EMPTY)
	b.PlayTurn(24, 32, boardstate.EMPTY)

	zorbistKey1 := b.GetZorbistKey()
	b.PrintFen(false)
	// Ready for en-passant push.
	b.PlayTurn(49, 33, boardstate.EMPTY)
	zorbistKey2 := b.GetZorbistKey()
	b.PrintFen(false)

	// Undo turn.
	b.UnplayTurn()
	b.PrintFen(false)

	zorbistKey3 := b.GetZorbistKey()
	checkZorbistBeforeAfter(t, zorbistKey1, zorbistKey2, zorbistKey3)
}

func checkZorbistBeforeAfter(t *testing.T, zorbistKey1, zorbistKey2, zorbistKey3 uint64) {
	if zorbistKey1 == zorbistKey2 {
		t.Errorf("Expected zobristKey1 to not equal zorbistKey2")
	}

	if zorbistKey2 == zorbistKey3 {
		t.Errorf("Expected zobristKey2 to not equal zorbistKey3")
	}

	if zorbistKey3 != zorbistKey1 {
		t.Errorf("Expected zobristKey1 to equal zorbistKey3")
	}
}

func TestZorbistLongGame(t *testing.T) {
	// Zorbist sanity check over a long random game.
	for i := 0; i < 25; i++ {
		b := boardstate.Initial()
		zorbistKeys := make([]uint64, 0, 200)
		for i := 0; i < 100; i++ {
			moves := make([]*boardstate.Move, 0, 1000)
			movegenerator.GenMovesInto(b, &moves)
			if len(moves) == 0 {
				break
			}
			randomIndex := rand.Intn(len(moves))
			b.PlayTurn(moves[randomIndex].Src, moves[randomIndex].Dst, moves[randomIndex].PromotePiece)
			zorbistKeys = append(zorbistKeys, b.GetZorbistKey())
		}

		for i := len(zorbistKeys) - 1; i > 0; i-- {
			b.UnplayTurn()
			require.Equal(t, b.GetZorbistKey(), zorbistKeys[i-1])
		}

	}

}
