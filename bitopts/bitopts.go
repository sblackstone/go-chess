package bitopts
import ("fmt")
func SetBit(n uint64, pos uint8) uint64 {
  n |= (1 << pos)
  return n
}

func ClearBit(n uint64, pos uint8) uint64 {
    var mask uint64 = ^(1 << pos)
    n &= mask
    return n
}

func TestBit(n uint64, pos uint8) bool {
    return n & (1 << pos) > 0;
}

func FlipBit(n uint64, pos uint8) uint64 {
  n ^= (1 << pos)
  return n
}

func GridToLinear(i uint8, j uint8) uint8 {
	return i*8 + j
}


func Print(n uint64, highlight uint8) {
  var i,j uint8;
  for i = 0; i < 8; i++ {
    for j = 0; j < 8; j++ {
      pos := GridToLinear(i,j)
      if (pos == highlight) {
        fmt.Printf(" * ");
      } else if (TestBit(n, pos)) {
        fmt.Printf(" X ");
      } else {
        fmt.Printf(" - ");
      }
    }
    fmt.Println()
  }

}

/*
func (b *BoardState) Print() {
	pieces := make([][]string, 2)
	pieces[BLACK] = []string{"♖", "♘", "♗", "♕", "♔", "♙"};
	pieces[WHITE] = []string{"♜", "♞", "♝", "♛", "♚", "♟"};
	var i, j uint8
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			color := b.ColorOfSquare(gridToLinear(i, j))
			if color == EMPTY {
				fmt.Printf(" - ")
			} else {
				piece := b.PieceOfSquare(gridToLinear(i, j))
				fmt.Printf(" %s ", pieces[color][piece])
			}
		}
		fmt.Println()
	}
}
*/
