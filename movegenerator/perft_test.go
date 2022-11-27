package movegenerator

import (
	"testing"

	"github.com/sblackstone/go-chess/boardstate"
)

func genPerf(state *boardstate.BoardState, depth int) int {
	successors := GenLegalSuccessors(state)
	result := 0
	if depth == 1 {
		return len(successors)
	} else {
		for _, succ := range successors {
			result += genPerf(succ, depth-1)
		}
	}
	return result
}

func genPerfMakeUnmake(b *boardstate.BoardState, depth int) int {
	moves := GenMoves(b)
	result := 0
	for _, move := range moves {
		b.PlayTurnFromMove(move)
		if !IsInCheck(b, b.EnemyColor()) {
			if depth == 1 {
				result += 1
			} else {
				result += genPerfMakeUnmake(b, depth-1)
			}
		}
		b.UnplayTurn()
	}
	return result
}

func BenchmarkPerf(b *testing.B) {
	board, _ := boardstate.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v := genPerf(board, 6)
		b.ReportMetric(float64(v), "positions")
	}
}

func BenchmarkPerfMakeUnmake(b *testing.B) {
	board, _ := boardstate.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		v := genPerfMakeUnmake(board, 5)
		b.ReportMetric(float64(v), "positions")

	}
}

func TestPerftPositions(t *testing.T) {
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

	for _, pt := range perfTests {
		b, err := boardstate.FromFEN(pt.fen)
		if err != nil {
			t.Errorf("%v\n", err)
		} else {
			result := genPerf(b, pt.depth)
			result2 := genPerfMakeUnmake(b, pt.depth)
			if result != pt.expected {
				t.Errorf("test: %v, genPerf: %v: Expected %v, got %v", pt.name, pt.fen, pt.expected, result)
			}
			if result2 != pt.expected {
				t.Errorf("test: %v, genPerfMakeUnmake: %v: Expected %v, got %v", pt.name, pt.fen, pt.expected, result2)
			}

		}
	}

}
