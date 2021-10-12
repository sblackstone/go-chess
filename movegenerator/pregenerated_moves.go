package movegenerator
import (
  "github.com/sblackstone/go-chess/bitopts"
)
func genKnightMoveBitBoards() [64]uint64 {
  var result [64]uint64;

  for y := 18; y <= 42; y+=8 {
    for x := 0; x < 4; x++ {
      pos := uint8(y + x);
      result[pos] = bitopts.SetBit(result[pos], pos);
    }
  }
  return result;
}

/*

56 57 58 59 60 61 62 63
48 49 50 51 52 53 54 55
40 41 42 43 44 45 46 47
32 33 34 35 36 37 38 39
24 25 26 27 28 29 30 31
16 17 18 19 20 21 22 23
08 09 10 11 12 13 14 15
00 01 02 03 04 05 06 07

18-21
26-29
34-37
42-45




*/
