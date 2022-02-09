package bitopts
import (
  "fmt"
  "math/bits"
//  "strconv"
)

func SetBit(n uint64, pos uint8) uint64 {
  n |= (1 << pos)
  return n
}

func FindTwoPiecePositions(n uint64) []uint8 {
  var result []uint8
  fmt.Printf("%064b\n", n)
  trailing := bits.TrailingZeros64(n)
  leading  := bits.LeadingZeros64(n)

  fmt.Printf("Trailing %v\n", trailing)
  fmt.Printf("Leading %v\n", leading)

  if (trailing == 64 && leading == 64) {
    return result
  }

  if (trailing < 64) {
    result = append(result, uint8(trailing))
  }

  if (leading < 64) {
    leadingPos := 64 - leading - 1
    if (leadingPos != trailing) {
      result = append(result, uint8(leadingPos))
    }
  }


  //result = append(result, uint8(trailing))
  //result = append(result, uint8(leading))
  return result
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
  var rank,file uint8;
  for rank = 7; rank < 8; rank-- {
    for file = 0; file < 8; file++ {
      pos := RankFileToSquare(rank,file)
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
