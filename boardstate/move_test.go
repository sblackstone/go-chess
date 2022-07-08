package boardstate

import "testing"

func TestMoveFromUCI(t *testing.T) {
	m, err := MoveFromUCI("a2a4")
	if err != nil {
		t.Error(err)
	}
	if m.Src != 8 || m.Dst != 24 || m.PromotePiece != EMPTY {
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
		{"a7a8q", QUEEN},
		{"a7a8n", KNIGHT},
		{"a7a8r", ROOK},
		{"a7a8b", BISHOP},
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

/*
	var perfTests = []struct {
		name     string
		fen      string
		depth    int
		expected int
	}{
		{"Initial Position", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", 5, 4865609},
		{"Position 2", "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 0", 3, 97862},
		{"Position 3", "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0", 5, 674624},
		{"Position 4A", "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", 3, 9467},
		{"Position 4B", "r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1", 3, 9467},
		{"Position 5", "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8", 3, 62379},
		{"Position 6", "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10", 3, 89890},
	}
*/
