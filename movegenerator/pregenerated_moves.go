package movegenerator
import (
  "github.com/sblackstone/go-chess/bitopts"
)

/*
A      B    C     D

      -17  -15
-10              -6

+6               +10
      +15  +17
*/

func genKnightMoves() [64][]uint8 {
  var result [64][]uint8;
  var row,col uint8
  for row = 0; row < 8; row++ {
    for col =0; col < 8; col++ {
      pos := bitopts.GridToLinear(row,col)
      // A
      if col >= 2 {
        if row >= 1 {
          result[pos] = append(result[pos], pos - 10)
        }
        if row <= 6 {
          result[pos] = append(result[pos], pos + 6)
        }
      }

      // B
      if col >= 1 {
        if row >= 2 {
          result[pos] = append(result[pos], pos - 17)

        }
        if row <= 5 {
          result[pos] = append(result[pos], pos + 15)
        }
      }

      // C
      if col <= 6 {
        if row >= 2 {
          result[pos] = append(result[pos], pos - 15)
        }
        if row <= 5 {
          result[pos] = append(result[pos], pos + 17)
        }
      }

      // D
      if col <= 5 {
        if row >= 1 {
          result[pos] = append(result[pos], pos - 6)
        }
        if row <= 6 {
          result[pos] = append(result[pos], pos + 10)
        }
      }


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


      -17  -15
-10              -6

-6               +10
      +15  +17

18-21
26-29
34-37
42-45




*/
