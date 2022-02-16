package boardstate

import (
"strings"
	"github.com/sblackstone/go-chess/bitopts"
  "errors"
  "fmt"
)

func FromFEN(fenString string) (*BoardState, error) {
  var rank,file int8
  var j int
  b := Blank()


  fenSlice := strings.Split(fenString, "")
  rank = 7


  addPiece := func(r int8, f int8, color int8, piece int8) {
    b.SetSquare(bitopts.RankFileToSquare(r,f), color, piece)
    file += 1
  }

  for j = range(fenSlice) {
    char := fenSlice[j]
    fmt.Printf("%v %v %v\n", rank, file, char)

    if (char == "p") {
      addPiece(rank, file, BLACK, PAWN)
    }
    if (char == "r") {
      addPiece(rank, file, BLACK, ROOK)
    }

    if (char == "n") {
      addPiece(rank, file, BLACK, KNIGHT)
    }
    if (char == "b") {
      addPiece(rank, file, BLACK, BISHOP)
    }
    if (char == "q") {
      addPiece(rank, file, BLACK, QUEEN)
    }
    if (char == "k") {
      addPiece(rank, file, BLACK, KING)
    }

    if (char == "P") {
      addPiece(rank, file, WHITE, PAWN)
    }
    if (char == "R") {
      addPiece(rank, file, WHITE, ROOK)
    }

    if (char == "N") {
      addPiece(rank, file, WHITE, KNIGHT)
    }
    if (char == "B") {
      addPiece(rank, file, WHITE, BISHOP)
    }
    if (char == "Q") {
      addPiece(rank, file, WHITE, QUEEN)
    }
    if (char == "K") {
      addPiece(rank, file, WHITE, KING)
    }

    if (char == " ") {
      return nil, errors.New("Unexpected ' ' in FEN, expected piece placement")
    }

    if char == "/" {
      rank -= 1
      file = 0
    }

    if rank == 0 && file == 8 {
      break;
    }
  }

  j += 1

  if fenSlice[j] != " " {
    return nil, errors.New("Expected space after board output")
  }

  j += 1

  if fenSlice[j] == "w" {
    b.SetTurn(WHITE)
  } else if fenSlice[j] == "b" {
    b.SetTurn(BLACK)
  } else {
    return nil, errors.New("Expected value for turn: " + fenSlice[j])

  }


  fmt.Println(j)
  fmt.Println(fenSlice[j])
  return b, nil
}
