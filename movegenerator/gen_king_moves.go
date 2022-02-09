package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func genAllKingMoves() [64][]uint8 {
  var result [64][]uint8;
  var rank,file uint8
  for rank = 7; rank < 8; rank-- {
    for file =0; file < 8; file++ {
      pos := bitopts.RankFileToSquare(rank,file)
			if (rank >= 1) {
				result[pos] = append(result[pos], pos - 8);
				if (file > 0) {
					result[pos] = append(result[pos], pos - 9);
				}
				if (file < 7) {
					result[pos] = append(result[pos], pos - 7);
				}
			}

			if (rank <= 7) {
				result[pos] = append(result[pos], pos + 8);
				if (file > 0) {
					result[pos] = append(result[pos], pos + 7);
				}
				if (file < 7) {
					result[pos] = append(result[pos], pos + 9);
				}
			}


			if (file >= 1) {
				result[pos] = append(result[pos], pos - 1);
			}
			if (file <= 6) {
				result[pos] = append(result[pos], pos + 1);
			}


		}
  }
  return result;
}

func genKingMoves(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
  return result;
}

/*
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

*/