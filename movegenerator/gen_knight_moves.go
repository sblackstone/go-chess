package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
)

/*
00 01 02 03 04 05 06 07
08 09 10 11 12 13 14 15
16 17 18 19 20 21 22 23
24 25 26 27 28 29 30 31
32 33 34 35 36 37 38 39
40 41 42 43 44 45 46 47
48 49 50 51 52 53 54 55
56 57 58 59 60 61 62 63


 A      B    C    D   <-- These labels match code sections...

      -17  -15
-10              -6

+6               +10
      +15  +17

*/


func genAllKnightMoves() [64][]uint8 {
  var result [64][]uint8;
  var rank,file uint8
  for rank = 0; rank < 8; rank++ {
    for file =0; file < 8; file++ {
      pos := bitopts.RankFileToSquare(rank,file)
      // A
      if file >= 2 {
        if rank >= 1 {
          result[pos] = append(result[pos], pos - 10)
        }
        if rank <= 6 {
          result[pos] = append(result[pos], pos + 6)
        }
      }

      // B
      if file >= 1 {
        if rank >= 2 {
          result[pos] = append(result[pos], pos - 17)

        }
        if rank <= 5 {
          result[pos] = append(result[pos], pos + 15)
        }
      }

      // C
      if file <= 6 {
        if rank >= 2 {
          result[pos] = append(result[pos], pos - 15)
        }
        if rank <= 5 {
          result[pos] = append(result[pos], pos + 17)
        }
      }

      // D
      if file <= 5 {
        if rank >= 1 {
          result[pos] = append(result[pos], pos - 6)
        }
        if rank <= 6 {
          result[pos] = append(result[pos], pos + 10)
        }
      }


    }
  }
  return result;
}

func genSingleKnightMoves(b *boardstate.BoardState, knightPos uint8) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	allKnightMoves := genAllKnightMoves(); // TODO: THIS MUST BE MEMOIZED SOMEHOW.
	for i := range(allKnightMoves[knightPos]) {
		move := allKnightMoves[knightPos][i];
		if b.ColorOfSquare(move) != b.GetTurn() {
			result = append(result, b.CopyPlayTurn(knightPos, move, boardstate.EMPTY))
		}
	}
	return result
}



func genKnightMoves(b *boardstate.BoardState) []*boardstate.BoardState {
	var result []*boardstate.BoardState;
	knightPositions := b.FindPieces(b.GetTurn(), boardstate.KNIGHT)
	//fmt.Printf("%v\n", rookPositions)
	for i := 0; i < len(knightPositions); i++ {
		result = append(result, genSingleKnightMoves(b, knightPositions[i])...)
	}

  return result;
}
