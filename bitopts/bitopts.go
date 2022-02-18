package bitopts
import (
  "fmt"
  "math/bits"
  "strings"
  "strconv"
  "errors"
)

func SquareToAlgebraic(pos int8) string {
  rank, file := SquareToRankFile(pos)
  ranks := []string{"1","2","3","4","5","6","7","8"}
  files := []string{"a","b","c","d","e","f","g","h"}
  return files[int(file)] + ranks[int(rank)]
}

func AlgebraicToSquare(algrbraicSquare string) (int8, error) {
  var result int8

  if len(algrbraicSquare) != 2 {
    return -1, errors.New("algrbraicSquare must be a string of len 2 " + fmt.Sprint(result))
  }

  algrbraicSquare = strings.ToLower(algrbraicSquare)
  var parts = strings.Split(algrbraicSquare, "")
  result = int8(algrbraicSquare[0]) - int8('a')

  if (result < 0 || result > 7) {
    return -1, errors.New("Invalid file: " + fmt.Sprint(result))
  }

  rank, err := strconv.Atoi(parts[1])
  if err != nil {
    return -1, err
  }
  rank -= 1

  if rank > 7 {
    return -1, errors.New("Invalid Rank: " + fmt.Sprint(rank))
  }

  result += int8(rank) * 8
  return result,nil
}

func SetBit(n uint64, pos int8) uint64 {
  n |= (1 << pos)
  return n
}

func FindTwoPiecePositions(n uint64) []int8 {
  var result []int8
  trailing := bits.TrailingZeros64(n)

  if (trailing == 64) {
    return result
  }

  leading  := bits.LeadingZeros64(n)

  if (trailing < 64) {
    result = append(result, int8(trailing))
  }

  if (leading < 64) {
    leadingPos := 64 - leading - 1
    if (leadingPos != trailing) {
      result = append(result, int8(leadingPos))
    }
  }

  return result
}

func ClearBit(n uint64, pos int8) uint64 {
    var mask uint64 = ^(1 << pos)
    n &= mask
    return n
}

func TestBit(n uint64, pos int8) bool {
    return n & (1 << pos) > 0;
}

func FlipBit(n uint64, pos int8) uint64 {
  n ^= (1 << pos)
  return n
}

func RankFileToSquare(rank int8, file int8) int8 {
	return rank*8 + file
}

func RankOfSquare(n int8) int8 {
  return n / 8
}

func FileOfSquare(n int8) int8 {
  return n % 8
}


func SquareToRankFile(n int8) (int8, int8) {
  return RankOfSquare(n), FileOfSquare(n)
}


func Print(n uint64, highlight int8) {
  var rank,file int8;
  for rank = 7; rank >= 0; rank-- {
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
