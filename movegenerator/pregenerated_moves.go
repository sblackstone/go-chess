package movegenerator
import (
  "github.com/sblackstone/go-chess/bitopts"
)
func genKnightMoveBitBoards() [64]uint64 {
  var result [64]uint64;

  for y := 18; y <= 42; y+=8 {
    for x := 0; x < 4; x++ {
      pos := uint8(y + x);
      result[pos] = bitopts.SetBit(result[pos], pos + 10);
      result[pos] = bitopts.SetBit(result[pos], pos - 10);
      result[pos] = bitopts.SetBit(result[pos], pos + 17);
      result[pos] = bitopts.SetBit(result[pos], pos - 17);
      result[pos] = bitopts.SetBit(result[pos], pos - 15);
      result[pos] = bitopts.SetBit(result[pos], pos + 15);
      result[pos] = bitopts.SetBit(result[pos], pos -  6);
      result[pos] = bitopts.SetBit(result[pos], pos +  6);
    }
  }
  return result;
}

/*

00 01 02 03 04 05 06 07
08 09 10 11 12 13 14 15
16 17 18 19 20 21 22 23
24 25 26 27 28 29 30 31
32 33 34 35 36 37 38 39
40 41 42 43 44 45 46 47
48 49 50 51 52 53 54 55
56 57 58 59 60 61 62 63

18:
8, 1, 3, 12, 28, 35, 33, 24

+- 10: 8, 28
+- 17: 1, 35
+- 15: 3, 33
+- 6:  12, 24




18-21
26-29
34-37
42-45




*/
