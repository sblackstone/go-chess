package movegenerator

import (
	"reflect"
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func TestGenLegalMoves(t *testing.T) {
	t.SkipNow()
}

func TestIsInCheckWhite(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(0, boardstate.WHITE, boardstate.KING)
	b.SetSquare(16, boardstate.BLACK, boardstate.QUEEN)

	if !IsInCheck(b, boardstate.WHITE) {
		t.Errorf("Expected WHITE to be in check")
	}

	if IsInCheck(b, boardstate.BLACK) {
		t.Errorf("Expected BLACK to NOT be in check")
	}

}

func TestCountMoves(t *testing.T) {
	b := boardstate.Initial()
	result := GenMovesCount(b)
	if result != 20 {
		t.Errorf("Expected GenMovesCount to return 20 for initial position")
	}
}

func TestIsInCheckBlack(t *testing.T) {
	b := boardstate.Blank()
	b.SetSquare(0, boardstate.BLACK, boardstate.KING)
	b.SetSquare(16, boardstate.WHITE, boardstate.QUEEN)

	if !IsInCheck(b, boardstate.BLACK) {
		t.Errorf("Expected BLACK to be in check")
	}

	if IsInCheck(b, boardstate.WHITE) {
		t.Errorf("Expected WHITE to NOT be in check")
	}

}

func TestCheckEndOfGameFoolsMate(t *testing.T) {
	b := boardstate.Initial()
	if CheckEndOfGame(b) != GAME_STATE_PLAYING {
		t.Errorf("Expected %v to be GAME_STATE_PLAYING", CheckEndOfGame(b))
	}

	b.PlayTurn(13, 21, boardstate.EMPTY) // f3
	b.PlayTurn(52, 44, boardstate.EMPTY) // e6
	b.PlayTurn(14, 30, boardstate.EMPTY) // g4
	b.PlayTurn(59, 31, boardstate.EMPTY)

	if CheckEndOfGame(b) != GAME_STATE_CHECKMATE {
		t.Errorf("Expected %v to be GAME_STATE_CHECKMATE", CheckEndOfGame(b))
	}

}

func TestCheckEndOfGameStalemate(t *testing.T) {
	b := boardstate.Blank()
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG, false)
	b.SetCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT, false)
	b.SetSquare(7, boardstate.WHITE, boardstate.KING)
	b.SetSquare(56, boardstate.BLACK, boardstate.KING)

	if CheckEndOfGame(b) != GAME_STATE_PLAYING {
		t.Errorf("Expected %v to be GAME_STATE_PLAYING", CheckEndOfGame(b))
	}
	b.SetSquare(13, boardstate.BLACK, boardstate.QUEEN)

	if CheckEndOfGame(b) != GAME_STATE_STALEMATE {
		t.Errorf("Expected %v to be GAME_STATE_STALEMATE", CheckEndOfGame(b))
	}

}

func TestGenGenSuccessorsInitialPosition(t *testing.T) {
	b := boardstate.Initial()
	successors := GenSuccessors(b)

	if len(successors) != 20 {
		t.Errorf("Expected initial successors to be 20, got %v", len(successors))
	}

	b.PlayTurn(8, 24, boardstate.EMPTY)

	successors2 := GenSuccessors(b)

	if len(successors2) != 20 {
		t.Errorf("Expected initial successors to be 20, got %v", len(successors2))
		for i := range successors2 {
			successors2[i].Print(127)
		}
	}

}

func TestGenLegalSuccessorsOpposition(t *testing.T) {
	b := boardstate.Blank()
	b.ClearCastling()
	b.SetSquare(43, boardstate.WHITE, boardstate.KING)
	b.SetSquare(27, boardstate.BLACK, boardstate.KING)

	legalWhite := genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, GenLegalSuccessors(b))

	expectedWhite := []int8{42, 44, 50, 51, 52}
	if !reflect.DeepEqual(legalWhite, expectedWhite) {
		t.Errorf("Expected %v to be %v", legalWhite, expectedWhite)
	}

	b.SetTurn(boardstate.BLACK)

	legalBlack := genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, GenLegalSuccessors(b))
	expectedBlack := []int8{18, 19, 20, 26, 28}
	if !reflect.DeepEqual(legalBlack, expectedBlack) {
		t.Errorf("Expected %v to be %v", legalBlack, expectedBlack)
	}

}

// GenMoves is just using all of the underlying infrastructure thats already well tested
// So we just check a simple case to ensure basic functionality
func TestGenMoves(t *testing.T) {
	b := boardstate.Initial()
	moves := GenMoves(b)
	if len(moves) != 20 {
		t.Errorf("Expected Gen Moves to find 20 inital moves")
	}
}

func TestGenMovesInto(t *testing.T) {
	moves := make([]*boardstate.Move, 0, 5)
	b := boardstate.Initial()
	GenMovesInto(b, &moves)
	if len(moves) != 20 {
		t.Errorf("Expected Gen Moves to find 20 inital moves")
	}
}
