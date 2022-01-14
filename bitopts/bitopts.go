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

func RankFileToSquare(rank uint8, file uint8) uint8 {
	return rank*8 + file
}

func SquareToRankFile(n uint8) (uint8, uint8) {
  return n / 8, n % 8
}


func Print(n uint64, highlight uint8) {
  var i,j uint8;
  for i = 0; i < 8; i++ {
    for j = 0; j < 8; j++ {
      pos := RankFileToSquare(i,j)
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
