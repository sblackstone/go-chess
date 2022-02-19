package fen


import (
  "testing"
  "fmt"
  "github.com/sblackstone/go-chess/boardstate"
  "reflect"
)

func TestParseTurn(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil || b.GetTurn() != boardstate.WHITE) {
    t.Errorf("Expected turn to be WHITE after import: %v %v", b.GetTurn(), err)
  }

  testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1"
  b2, err2 := FromFEN(testStr2)
  if (err != nil || b2.GetTurn() != boardstate.BLACK) {
    t.Errorf("Expected turn to be WHITE after import: %v %v", b2.GetTurn(), err2)
  }
}

func TestParseEnpassant(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil || b.GetEnpassant() != boardstate.NO_ENPASSANT) {
    t.Errorf("Expected GetEnpassant to be NO_ENPASSANT after import: %v %v", b.GetEnpassant(), err)
  }

  testStr2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq c4 0 1"
  b2, err2 := FromFEN(testStr2)
  if (err2 != nil || b2.GetEnpassant() != 26) {
    t.Errorf("Expected GetEnpassant to be 26 after import: %v %v", b2.GetEnpassant(), err2)
  }
}

func TestParseCastlingNone(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil) {
    t.Errorf("Expected valid import: %v", err)
  }
  if (b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG)) {
      t.Errorf("expected not to have  WHITE LONG")
  }
  if (b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected not to have  WHITE SHORT")
  }
  if (b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG)) {
      t.Errorf("expected not to have  BLACK LONG")
  }
  if (b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected not to have  BLACK SHORT")
  }
}


func TestParseCastlingShortOnly(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kk - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil) {
    t.Errorf("Expected valid import: %v", err)
  }
  if (b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG)) {
      t.Errorf("expected not to have  WHITE LONG")
  }
  if (!b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected TO have  WHITE SHORT")
  }
  if (b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG)) {
      t.Errorf("expected not to have  BLACK LONG")
  }
  if (!b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected TO have  BLACK SHORT")
  }
}


func TestParseCastlingLongOnly(t *testing.T) {
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Qq - 0 1"
  b, err := FromFEN(testStr)
  if (err != nil) {
    t.Errorf("Expected valid import: %v", err)
  }
  if (!b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_LONG)) {
      t.Errorf("expected TO have  WHITE LONG")
  }
  if (b.HasCastleRights(boardstate.WHITE, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected NOT TO have  WHITE SHORT")
  }
  if (!b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_LONG)) {
      t.Errorf("expected TO have  BLACK LONG")
  }
  if (b.HasCastleRights(boardstate.BLACK, boardstate.CASTLE_SHORT)) {
      t.Errorf("expected NOT have  BLACK SHORT")
  }
}



func TestErrortCases(t *testing.T) {
  t.Errorf("complete fen parser tests for error cases")
}


func TestAAdditionalBoardCases(t *testing.T) {
  t.Errorf("complete fen parser test other board configurations for import")
}

func TestFENParserDefaultBoard(t *testing.T) {
  correct := boardstate.Initial()
  testStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b, err := FromFEN(testStr)

  if err != nil {
    fmt.Printf("%v\n", err)
  }

  if !reflect.DeepEqual(correct,b) {
    t.Errorf("Expected import of default board to match the default board")
  }

}

func TestParseMoveCounters(t *testing.T) {
  testStr := "p6P/1p4P1/2p2P2/3pP3/3Pp3/2P2p2/1P4p1/P6p w - - 25 26"
  b, err := FromFEN(testStr)
  if err != nil {
    fmt.Printf("%v\n", err)
  }
  if (b.GetHalfMoves() != 25) {
    t.Errorf("Expected half moves to be 25")
  }
  if (b.GetFullMoves() != 26) {
    t.Errorf("Expected half moves to be 26")
  }


}
