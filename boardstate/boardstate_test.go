package boardstate

import (
	"testing"
  "fmt"
)


func TestGetColorBitboard(t *testing.T) {
	b := Initial()

	if b.colors[WHITE] != b.GetColorBitboard(WHITE) {
	  t.Errorf("GetColorBitboard wasn't right for WHITE")
	}

	if b.colors[BLACK] != b.GetColorBitboard(BLACK) {
	  t.Errorf("GetColorBitboard wasn't right for BLACK")
	}


}

func TestIncrementHalfMoves(t *testing.T) {
	b := Initial()
	if b.GetHalfMoves() != 0 {
		t.Errorf("Expected halfMoves to be 0, got %v", b.GetHalfMoves())
	}

	b.IncrementHalfMoves()

	if b.GetHalfMoves() != 1 {
		t.Errorf("Expected halfMoves to be 1, got %v", b.GetHalfMoves())
	}
}

func IncrementFullMoves(t *testing.T) {
	b := Initial()
	if b.GetHalfMoves() != 0 {
		t.Errorf("Expected halfMoves to be 0, got %v", b.GetHalfMoves())
	}

	b.IncrementHalfMoves()

	if b.GetHalfMoves() != 1 {
		t.Errorf("Expected halfMoves to be 1, got %v", b.GetHalfMoves())
	}
}

func TestSetGetHalfMoves(t *testing.T) {
	b := Initial()
	b.SetHalfMoves(512)
	if (b.GetHalfMoves() != 512) {
		t.Errorf("Expected half moves to be 512, got %v", b.GetHalfMoves())
	}
}

func TestSetGetFullMoves(t *testing.T) {
	b := Initial()
	b.SetFullMoves(512)
	if (b.GetFullMoves() != 512) {
		t.Errorf("Expected half moves to be 512, got %v", b.GetFullMoves())
	}
}

func TestFullMoveCountProperlyIncrementsOnMoves(t *testing.T) {
	b := Initial()
	if (b.GetFullMoves() != 1) {
		t.Errorf("Expected half moves to be 1, got %v", b.GetFullMoves())
	}
	b.PlayTurn(1,16,EMPTY)
	if (b.GetFullMoves() != 1) {
		t.Errorf("Expected half moves to be 1, got %v", b.GetFullMoves())
	}
	b.PlayTurn(57,40,EMPTY)
	if (b.GetFullMoves() != 2) {
		t.Errorf("Expected half moves to be 2, got %v", b.GetFullMoves())
	}
	b.PlayTurn(11,27,EMPTY)
	if (b.GetFullMoves() != 2) {
		t.Errorf("Expected half moves to be 2, got %v", b.GetFullMoves())
	}
	b.PlayTurn(48,40,EMPTY)
	if (b.GetFullMoves() != 3) {
		t.Errorf("Expected half moves to be 3, got %v", b.GetFullMoves())
	}



}

func TestHalfMoveCountResetsOnPawnMoves(t *testing.T) {
	// Pawn Push
	b := Initial()
	b.PlayTurn(1,16, EMPTY)
	b.PlayTurn(57,40, EMPTY)
	if (b.GetHalfMoves() != 2) {
		t.Errorf("Expected half moves to be 2, got %v", b.GetHalfMoves())
	}
	b.PlayTurn(11, 27, EMPTY)
	if (b.GetHalfMoves() != 0) {
		t.Errorf("Expected half moves to be 0, got %v", b.GetHalfMoves())
	}
}

func TestHalfMoveCountResetsOnCaptureMoves(t *testing.T) {
	// Pawn Push
	b := Initial()
	b.PlayTurn(1,18, EMPTY)
	b.PlayTurn(62,45, EMPTY)
	b.PlayTurn(18,35, EMPTY)

	if (b.GetHalfMoves() != 3) {
		t.Errorf("Expected half moves to be 3, got %v", b.GetHalfMoves())
	}
	b.PlayTurn(35, 45, EMPTY)
	if (b.GetHalfMoves() != 0) {
		t.Errorf("Expected half moves to be 0, got %v", b.GetHalfMoves())
	}
}




func testCastlingBoard() *BoardState {
	b := Blank()
	b.SetSquare(0, WHITE, ROOK)
	b.SetSquare(4, WHITE, KING)
	b.SetSquare(7, WHITE, ROOK)

	b.SetSquare(56, BLACK, ROOK)
	b.SetSquare(60, BLACK, KING)
	b.SetSquare(63, BLACK, ROOK)
	return b
}



func TestGenerateSuccessors(t *testing.T)  {
	b := Blank()
	b.SetSquare(22, WHITE, PAWN)
	m1 := CreateMove(22, 22+8, EMPTY)
	m2 := CreateMove(22, 22+16, EMPTY)

	var moves []*Move;


	moves = append(moves, m1)
	moves = append(moves, m2)

	successors := b.GenerateSuccessors(moves)

	if len(successors) != 2 {
		t.Errorf("Expected successors to be length 2");
	}

}

func TestTakingRooksDisableCastlingWhiteLong(t *testing.T) {
	b := testCastlingBoard();
	b.SetTurn(BLACK)

	if (!b.HasCastleRights(WHITE, CASTLE_LONG)) {
		t.Errorf("Expected WHITE to have CASTLE_LONG before rook move")
	}

	b.PlayTurn(56, 0, EMPTY)

	if (b.HasCastleRights(WHITE, CASTLE_LONG)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_LONG after rook move")
	}
}

func TestTakingRooksDisableCastlingWhiteShort(t *testing.T) {
	b := testCastlingBoard();
	b.SetTurn(BLACK)

	if (!b.HasCastleRights(WHITE, CASTLE_SHORT)) {
		t.Errorf("Expected WHITE to have CASTLE_SHORT before rook move")
	}

	b.PlayTurn(63, 7, EMPTY)

	if (b.HasCastleRights(WHITE, CASTLE_SHORT)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_SHORT after rook move")
	}
}


/////
func TestTakingRooksDisableCastlingBlackLong(t *testing.T) {
	b := testCastlingBoard();

	if (!b.HasCastleRights(BLACK, CASTLE_LONG)) {
		t.Errorf("Expected BLACK to have CASTLE_LONG before rook move")
	}

	b.PlayTurn(0, 56, EMPTY)

	if (b.HasCastleRights(BLACK, CASTLE_LONG)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG after rook move")
	}
}

func TestTakingRooksDisableCastlingBlackShort(t *testing.T) {
	b := testCastlingBoard();

	if (!b.HasCastleRights(BLACK, CASTLE_SHORT)) {
		t.Errorf("Expected BLACK to have CASTLE_SHORT before rook move")
	}

	b.PlayTurn(7, 63, EMPTY)

	if (b.HasCastleRights(BLACK, CASTLE_SHORT)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_SHORT after rook move")
	}
}

/////



func TestMoveKingLoosesRights(t *testing.T) {
	b := testCastlingBoard()
	b.PlayTurn(4, 12, EMPTY)
	if (b.HasCastleRights(WHITE, CASTLE_LONG)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_LONG after king move")
	}

	if (b.HasCastleRights(WHITE, CASTLE_SHORT)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_SHORT after king move")
	}

	b.PlayTurn(60, 52, EMPTY)

	if (b.HasCastleRights(BLACK, CASTLE_LONG)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG after king move")
	}

	if (b.HasCastleRights(BLACK, CASTLE_SHORT)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_SHORT after king move")
	}

}

func TestMoveRookLoosesHalfCastleRights(t *testing.T) {
	b := testCastlingBoard()
	b.PlayTurn(7, 15, EMPTY)
	if (b.HasCastleRights(WHITE, CASTLE_SHORT) || !b.HasCastleRights(WHITE, CASTLE_LONG)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_SHORT and have CASTLE_LONG after rook move")
	}

	b2 := testCastlingBoard()
	b2.PlayTurn(0, 8, EMPTY)
	if (b2.HasCastleRights(WHITE, CASTLE_LONG) || !b2.HasCastleRights(WHITE, CASTLE_SHORT)) {
		t.Errorf("Expected WHITE to NOT have CASTLE_LONG and have CASTLE_SHORT after rook move")
	}

	b3 := testCastlingBoard()
	b3.SetTurn(BLACK)
	b3.PlayTurn(63, 55, EMPTY)
	if (b3.HasCastleRights(BLACK, CASTLE_SHORT) || !b3.HasCastleRights(BLACK, CASTLE_LONG)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_SHORT, but have CASTLE_LONG after rook move")
	}

	b4 := testCastlingBoard()
	b4.SetTurn(BLACK)
	b4.PlayTurn(56, 48, EMPTY)
	if (b4.HasCastleRights(BLACK, CASTLE_LONG) || !b4.HasCastleRights(BLACK, CASTLE_SHORT)) {
		t.Errorf("Expected BLACK to NOT have CASTLE_LONG, but have CASTLE_SHORT after rook move")
	}

}

func TestInitialMoveCounters(t *testing.T) {
	b := Initial()
	if (b.GetFullMoves() != 1) {
		t.Errorf("Expected GetFullMoves to start at 1, got %v", b.GetFullMoves())
	}
	if (b.GetHalfMoves() != 0) {
		t.Errorf("Expected GetHalfMoves to start at 1, got %v", b.GetHalfMoves())
	}

}

func TestInitialCastlingRights(t *testing.T) {
	b := testCastlingBoard()

	if (!b.HasCastleRights(WHITE, CASTLE_LONG)) {
		t.Errorf("Expected WHITE to have CASTLE_LONG initially")
	}

	if (!b.HasCastleRights(WHITE, CASTLE_SHORT)) {
		t.Errorf("Expected WHITE to have CASTLE_SHORT initially")
	}

	if (!b.HasCastleRights(BLACK, CASTLE_LONG)) {
		t.Errorf("Expected BLACK to have CASTLE_LONG initially")
	}

	if (!b.HasCastleRights(BLACK, CASTLE_SHORT)) {
		t.Errorf("Expected BLACK to have CASTLE_SHORT initially")
	}

}

func TestCastleShortWhite(t *testing.T) {
	b := testCastlingBoard()
	b.PlayTurn(4,6, EMPTY)
	if b.PieceOfSquare(6) != KING {
		t.Errorf("Expected KING in sq 6");
	}
	if b.PieceOfSquare(5) != ROOK {
		t.Errorf("Expected ROOK in sq 5");
	}


}

func TestCastleLongWhite(t *testing.T) {
	b := testCastlingBoard()
	b.PlayTurn(4,2, EMPTY)
	if b.PieceOfSquare(2) != KING {
		t.Errorf("Expected KING in sq 2");
	}
	if b.PieceOfSquare(3) != ROOK {
		t.Errorf("Expected ROOK in sq 3");
	}
}

//////////////

func TestCastleShortBlack(t *testing.T) {
	b := testCastlingBoard()
	b.SetTurn(BLACK)
	b.PlayTurn(60,62, EMPTY)
	if b.PieceOfSquare(62) != KING {
		t.Errorf("Expected KING in sq 62");
	}
	if b.PieceOfSquare(61) != ROOK {
		t.Errorf("Expected ROOK in sq 61");
	}
}

func TestCastleLongBlack(t *testing.T) {
	b := testCastlingBoard()
	b.SetTurn(BLACK)
	b.PlayTurn(60,58, EMPTY)
	if b.PieceOfSquare(58) != KING {
		t.Errorf("Expected KING in sq 58");
	}
	if b.PieceOfSquare(59) != ROOK {
		t.Errorf("Expected ROOK in sq 59");
	}
}



func TestCopyPlayMove(t *testing.T) {
	b1 := Initial()
	b2 := b1.CopyPlayTurn(1, 18, EMPTY)
	if (b1 == b2) {
		t.Errorf("Expected b1 to be different than b2")
	}

	if (b2.PieceOfSquare(18) != KNIGHT || b2.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b2.PieceOfSquare(1) != EMPTY || b2.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}

	if (b1.PieceOfSquare(1) == EMPTY || b1.ColorOfSquare(1) == EMPTY) {
		t.Errorf("square 1 on b1 shouldn't be empty")
	}

	if (b1.PieceOfSquare(18) == KNIGHT || b1.ColorOfSquare(18) == WHITE) {
		t.Errorf("square 18 on b1 should be empty")
	}
}


func TestCopyPlayMoveFromTurn(t *testing.T) {
	b1 := Initial()
	b2 := b1.CopyPlayTurnFromMove(&Move{Src: 1, Dst: 18, PromotePiece: EMPTY})
	if (b1 == b2) {
		t.Errorf("Expected b1 to be different than b2")
	}

	if (b2.PieceOfSquare(18) != KNIGHT || b2.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b2.PieceOfSquare(1) != EMPTY || b2.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}

	if (b1.PieceOfSquare(1) == EMPTY || b1.ColorOfSquare(1) == EMPTY) {
		t.Errorf("square 1 on b1 shouldn't be empty")
	}

	if (b1.PieceOfSquare(18) == KNIGHT || b1.ColorOfSquare(18) == WHITE) {
		t.Errorf("square 18 on b1 should be empty")
	}
}


func TestEnpassantAsWhite(t *testing.T) {
	b := Initial()
	if (b.GetEnpassant() == 18) {
		t.Errorf("Did not expect enpassant to be set to 18, got %v", b.GetEnpassant())
	}
	b.PlayTurn(10,26, EMPTY)
	if (b.GetEnpassant() != 18) {
		t.Errorf("Expected enpassant to be set to 18, got %v", b.GetEnpassant())
	}
	b.PlayTurn(57,42, EMPTY)
	if (b.GetEnpassant() == 18) {
		t.Errorf("Expected enpassant to not be set to 18")
	}

}

func TestEnpassantAsBlack(t *testing.T) {
	b := Initial()

	if (b.GetEnpassant() == 45) {
		t.Errorf("Did not expect enpassant to be set to 45")
	}
	b.PlayTurn(53,37, EMPTY)
	if (b.GetEnpassant() != 45) {
		t.Errorf("Expected enpassant to be set to be 45, got %v", b.GetEnpassant())
	}
	b.PlayTurn(9,25, EMPTY)
	if (b.GetEnpassant() == 45) {
		t.Errorf("Did not expect enpassant to be set to 45")
	}

}



func TestEmptySquare(t *testing.T) {
	b := Initial()
	if (b.EmptySquare(10)) {
		t.Errorf("expected square 10 to not be empty\n")
	}
	if (b.EmptySquare(51)) {
		t.Errorf("expected square 51 to not be empty\n")
	}

	if (!b.EmptySquare(27)) {
		t.Errorf("expected square 27 to be empty\n")
	}

}

func TestEmptyOrEnemyOccupiedSquare(t *testing.T) {
	b := Initial()
	if (b.EmptyOrEnemyOccupiedSquare(10)) {
		t.Errorf("expected square 10 to not be empty/enemy occupiped\n")
	}
	if (!b.EmptyOrEnemyOccupiedSquare(18)) {
		t.Errorf("expected square 18 to be empty/enemy occupiped\n")
	}

	if (!b.EmptyOrEnemyOccupiedSquare(50)) {
		t.Errorf("expected square 50 to be empty/enemy occupiped\n")
	}

	b.ToggleTurn()

	if (b.EmptyOrEnemyOccupiedSquare(51)) {
		t.Errorf("expected square 10 to not be empty/enemy occupiped\n")
	}
	if (!b.EmptyOrEnemyOccupiedSquare(43)) {
		t.Errorf("expected square 43 to be empty/enemy occupiped\n")
	}

	if (!b.EmptyOrEnemyOccupiedSquare(10)) {
		t.Errorf("expected square 50 to be empty/enemy occupiped\n")
	}

}

func TestEnemyOccupriedSquare(t *testing.T) {
	b := Initial()
	if (b.EnemyOccupiedSquare(10)) {
		t.Errorf("expected square 10 to not be enemy occupiped\n")
	}
	if (b.EnemyOccupiedSquare(18)) {
		t.Errorf("expected square 18 to not be enemy occupiped\n")
	}

	if (!b.EnemyOccupiedSquare(50)) {
		t.Errorf("expected square 50 to be enemy occupiped\n")
	}

	b.ToggleTurn()

	if (b.EnemyOccupiedSquare(51)) {
		t.Errorf("expected square 10 to not be enemy occupiped\n")
	}
	if (b.EnemyOccupiedSquare(43)) {
		t.Errorf("expected square 43 to not be enemy occupiped\n")
	}

	if (!b.EnemyOccupiedSquare(10)) {
		t.Errorf("expected square 50 to be enemy occupiped\n")
	}



}

func TestCopy(t *testing.T) {
	b1 := Initial()
	b1.PlayTurn(1, 18, EMPTY)
	b1.SetFullMoves(25)
	b1.SetHalfMoves(26)
	b2 := b1.Copy()

	if (b2.GetFullMoves() != 25 || b2.GetHalfMoves() != 26) {
		t.Errorf("Expected full moves to be copied")
	}

	if (b2.PieceOfSquare(18) != KNIGHT || b2.ColorOfSquare(18) != WHITE) {
		t.Errorf("square 18 isn't a white knight")
	}

	if (b2.PieceOfSquare(1) != EMPTY || b2.ColorOfSquare(1) != EMPTY) {
		t.Errorf("square 1 isn't empty")
	}

	if b2.GetTurn() != BLACK {
		t.Errorf("Expected turn after PlayTurn to be black");
	}

	b2.PlayTurn(57, 42, EMPTY)

	if b1.PieceOfSquare(42) != EMPTY || b1.ColorOfSquare(42) != EMPTY || b1.GetTurn() != BLACK {
		t.Errorf("copy is effecting the original")
	}

}




func TestFindPieces(t *testing.T) {
	b := Initial()

	res := b.FindPieces(BLACK, QUEEN)
	if (len(res) != 1 || res[0] != 59) {
		t.Errorf("Expected %v to be [59]\n", res)
	}

	res2 := b.FindPieces(WHITE, QUEEN)
	if (len(res2) != 1 || res2[0] != 3) {
		t.Errorf("Expected %v to be [3]\n", res2)
	}

	res3 := b.FindPieces(BLACK, KNIGHT)
	if (len(res3) != 2 || res3[0] != 57 || res3[1] != 62) {
		t.Errorf("Expected %v to be [57,62]\n", res3)
	}

	res4 := b.FindPieces(WHITE, KNIGHT)
	if (len(res4) != 2 || res4[0] != 1 || res4[1] != 6) {
		t.Errorf("Expected %v to be [1,6]\n", res4)
	}


	res5 := b.FindPieces(WHITE, PAWN)
	if (len(res5) != 8 || res5[0] != 8 || res5[7] != 15) {
		t.Errorf("Expected %v to be [8,9,10,11,12,13,14,15]\n", res5)
	}

	res6 := b.FindPieces(BLACK, PAWN)
	if (len(res6) != 8 || res6[0] != 48 || res6[7] != 55) {
		t.Errorf("Expected %v to be [48,49,50,51,52,53,54,55]\n", res6)
	}

	b2 := Blank()
	res7 := b2.FindPieces(BLACK, PAWN)
	if (len(res7) != 0) {
		t.Errorf("Expected %v to be []\n", res7)
	}

	res8 := b2.FindPieces(WHITE, QUEEN)
	if (len(res8) != 0) {
		t.Errorf("Expected %v to be []\n", res8)
	}


	b3 := Initial()
	b3.SetSquare(57, EMPTY, EMPTY)
	res9 := b3.FindPieces(BLACK, KNIGHT)
	if (len(res9) != 1 || res9[0] != 62) {
		t.Errorf("Expected %v to be [62]\n", res9)
	}


	b4 := Blank()
	b4.SetSquare(55, BLACK, PAWN)
	res10 := b4.FindPieces(BLACK, PAWN)
	if (len(res10) != 1 || res10[0] != 55) {
		t.Errorf("Expected %v to be [55]\n", res10)
	}
}

func TestEnpassantCaptureClearsEnemyPawnBlackLowerFile(t *testing.T) {
	b1 := Blank()
	b1.SetSquare(8, WHITE, PAWN)
	b1.SetSquare(25, BLACK, PAWN)
	//b1.Print(255)
	b1.PlayTurn(8, 24, EMPTY)
	fmt.Println(b1.GetEnpassant())
	//b1.Print(255)
	fmt.Println()

	b1.PlayTurn(25, 16, EMPTY)
	//b1.Print(255)

	if (b1.PieceOfSquare(24) != EMPTY) {
		fmt.Println("Expected 24 to be empty")
	}

}

func TestEnpassantCaptureClearsEnemyPawnBlackHigherFile(t *testing.T) {
	b1 := Blank()
	b1.SetSquare(11, WHITE, PAWN)
	b1.SetSquare(26, BLACK, PAWN)
	//b1.Print(255)
	//fmt.Println()

	b1.PlayTurn(11, 27, EMPTY)
	//fmt.Println(b1.GetEnpassant())
	//b1.Print(255)
	//fmt.Println()

	b1.PlayTurn(26, 19, EMPTY)
	//b1.Print(255)

	if (b1.PieceOfSquare(27) != EMPTY) {
		fmt.Println("Expected 27 to be empty")
	}

}


///////+++++++++++++++++++++++++/
func TestEnpassantCaptureClearsEnemyPawnWhiteLowerFile(t *testing.T) {
	b1 := Blank()
	b1.SetSquare(55, BLACK, PAWN)
	b1.SetSquare(38, WHITE, PAWN)
	b1.SetTurn(BLACK)
	//b1.Print(255)
	b1.PlayTurn(55, 39, EMPTY)
	//fmt.Println(b1.GetEnpassant())
	//b1.Print(255)
	//fmt.Println()

	b1.PlayTurn(38, 47, EMPTY)
	//b1.Print(255)

	if (b1.PieceOfSquare(39) != EMPTY) {
		fmt.Println("Expected 39 to be empty")
	}

}


func TestEnpassantCaptureClearsEnemyPawnWhiteHigherFile(t *testing.T) {
	b1 := Blank()
	b1.SetSquare(49, BLACK, PAWN)
	b1.SetSquare(34, WHITE, PAWN)
	b1.SetTurn(BLACK)
	//b1.Print(255)
	b1.PlayTurn(49, 33, EMPTY)
	//fmt.Println(b1.GetEnpassant())
	//b1.Print(255)
	//fmt.Println()

	b1.PlayTurn(34, 41, EMPTY)
	//b1.Print(255)

	if (b1.PieceOfSquare(33) != EMPTY) {
		fmt.Println("Expected 33 to be empty")
	}

}

//////++++++++++++++++++++++++++/



func TestCreateBlankBoard(t *testing.T) {
  b := Blank();
  if (len(b.colors) != 2) {
    t.Errorf("Colors are not 0,0")
  }
  if (len(b.pieces) != 6) {
    t.Errorf("pieces are not 0,0")
  }
	if (b.GetEnpassant() != NO_ENPASSANT) {
		t.Errorf("Expected an initial board to have NO_ENPASSANT")
	}

}

func TestMovePieceWhiteKnight(t *testing.T) {
	b := Initial()

	if b.PieceOfSquare(1) != KNIGHT {
		t.Errorf("Exepcted piece in square 1 to be KNIGHT")
	}
	if b.ColorOfSquare(1) != WHITE {
		t.Errorf("Exepcted color in square 1 to be WHITE")
	}

	if b.PieceOfSquare(18) != EMPTY {
		t.Errorf("Exepcted piece in square 18 to be empty")
	}
	if b.ColorOfSquare(18) != EMPTY {
		t.Errorf("Exepcted color in square 18 to be empty")
	}

	b.MovePiece(1, 18)

	if b.PieceOfSquare(1) != EMPTY {
		t.Errorf("Exepcted piece in square 1 to be empty")
	}
	if b.ColorOfSquare(1) != EMPTY {
		t.Errorf("Exepcted color in square 1 to be empty")
	}

	if b.PieceOfSquare(18) != KNIGHT {
		t.Errorf("Exepcted piece in square 18 to be KNIGHT")
	}
	if b.ColorOfSquare(18) != WHITE {
		t.Errorf("Exepcted color in square 18 to be WHITE")
	}
}

func TestMovePieceBlackKnight(t *testing.T) {
	b := Initial()

	if b.PieceOfSquare(57) != KNIGHT {
		t.Errorf("Exepcted piece in square 1 to be KNIGHT")
	}
	if b.ColorOfSquare(57) != BLACK {
		t.Errorf("Exepcted color in square 1 to be BLACK")
	}

	if b.PieceOfSquare(42) != EMPTY {
		t.Errorf("Exepcted piece in square 42 to be empty")
	}
	if b.ColorOfSquare(42) != EMPTY {
		t.Errorf("Exepcted color in square 42 to be empty")
	}

	b.MovePiece(57, 42)

	if b.PieceOfSquare(57) != EMPTY {
		t.Errorf("Exepcted piece in square 1 to be empty")
	}
	if b.ColorOfSquare(57) != EMPTY {
		t.Errorf("Exepcted color in square 1 to be empty")
	}

	if b.PieceOfSquare(42) != KNIGHT {
		t.Errorf("Exepcted piece in square 18 to be KNIGHT")
	}
	if b.ColorOfSquare(42) != BLACK {
		t.Errorf("Exepcted color in square 18 to be BLACK")
	}
}


func testInitialBoard(t *testing.T, b *BoardState) {
	// 64bit int versions of the chess board assuming a correct inital state.
  expectedPieces := []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
  expectedColors := []uint64{65535,18446462598732840960};

	if (b.GetEnpassant() != NO_ENPASSANT) {
		t.Errorf("Expected an initial board to have NO_ENPASSANT")
	}

  for i := range(b.pieces) {
    if (b.pieces[i] != expectedPieces[i]) {
      t.Errorf("Initial Board mismatch pieces");
    }
  }

  for i := range(b.colors) {
    if (b.colors[i] != expectedColors[i]) {
      t.Errorf("Initial Board mismatch colors");
    }
  }
}

func TestCreateInitialBoard(t *testing.T) {
	testInitialBoard(t, Initial())
}

func TestCreateInitialBoardManual(t *testing.T) {
	testInitialBoard(t, initialManual())
}


func TestColorOfSquare(t *testing.T) {
  b := Initial()
  if (b.ColorOfSquare(63) != BLACK) {
    t.Errorf("Expected square 63 to be BLACK")
  }
  if (b.ColorOfSquare(0) != WHITE) {
    t.Errorf("Expected square 0 to be WHITE")
  }
  if (b.ColorOfSquare(32) != EMPTY) {
    t.Errorf("Expected square 32 to be EMPTY")
  }
}


func TestPieceOfSquare(t *testing.T) {
  b := Initial()

  expected := [][]int8{
    {0,ROOK},
    {1,KNIGHT},
    {2,BISHOP},
    {3,QUEEN},
    {4,KING},
    {8,PAWN},
		{32,EMPTY},
		{56,ROOK},
    {57,KNIGHT},
    {58,BISHOP},
    {59,QUEEN},
    {60,KING},
    {48,PAWN},
  }

  for _, expect := range(expected) {
    if b.PieceOfSquare(expect[0]) != expect[1] {
        t.Errorf("Expected square %v to be %v", expect[0], expect[1])
    }
  }
}

func TestSetSquare(t *testing.T) {
	b := Initial()

	if (b.ColorOfSquare(0) != WHITE) {
		t.Errorf("Expected color of square 0 to be WHITE (%v), got %v", WHITE, b.ColorOfSquare(0))
	}

	if (b.PieceOfSquare(0) != ROOK) {
		t.Errorf("Expected piece of square 0 to be ROOK (%v), got %v", ROOK, b.PieceOfSquare(0))
	}

	b.SetSquareRankFile(0,0, BLACK, PAWN)

	if (b.ColorOfSquare(0) != BLACK) {
		t.Errorf("Expected color of square 0 to be BLACK")
	}
	if (b.PieceOfSquare(0) != PAWN) {
		t.Errorf("Expected piece of square 0 to be PAWN (%v), got %v", PAWN, b.PieceOfSquare(0))
	}

	b.SetSquareRankFile(0,0, EMPTY, EMPTY)
	if (b.ColorOfSquare(0) != EMPTY) {
		t.Errorf("Expected color of square 0 to be EMPTY")
	}
	if (b.PieceOfSquare(0) != EMPTY) {
		t.Errorf("Expected piece of square 0 to be EMPTY (%v), got %v", EMPTY, b.PieceOfSquare(0))
	}



}
