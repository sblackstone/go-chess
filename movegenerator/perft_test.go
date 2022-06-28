package movegenerator

import (
  "testing"
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/fen"
  //"reflect"
)



func genPerf(state *boardstate.BoardState, depth int) int {
  successors := GenLegalSuccessors(state)
  result := 0
  if (depth == 1) {
    return len(successors)
  } else {
    for _, succ := range(successors) {
      result += genPerf(succ, depth - 1)
    }
  }
  return result
}


func TestPerft(t *testing.T) {
  b := boardstate.Initial()
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)

}


// https://www.chessprogramming.org/Perft_Results

func TestPerftPosition2(t *testing.T) {
  b, err := fen.FromFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 0")

  if err != nil {
    t.Errorf("%v", err)
    return
  }
  b.Print(127)
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)


}

func TestPosition3(t *testing.T) {
  b, err := fen.FromFEN("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0")

  if err != nil {
    t.Errorf("%v", err)
    return
  }
  b.Print(127)
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)


}


func TestPosition4(t *testing.T) {
  b, err := fen.FromFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1")

  if err != nil {
    t.Errorf("%v", err)
    return
  }
  b.Print(127)
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)


}

func TestPosition5(t *testing.T) {
  b, err := fen.FromFEN("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8")

  if err != nil {
    t.Errorf("%v", err)
    return
  }
  b.Print(127)
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)

}


func TestPosition6(t *testing.T) {
  b, err := fen.FromFEN("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10")

  if err != nil {
    t.Errorf("%v", err)
    return
  }
  b.Print(127)
  result := genPerf(b, 3)
  t.Errorf("%v\n", result)

}
