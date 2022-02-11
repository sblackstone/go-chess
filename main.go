package main

import (
  "github.com/sblackstone/go-chess/boardstate"
  "github.com/sblackstone/go-chess/bitopts"
  "fmt"

)


func main() {
  b := boardstate.Initial();
  b.Print(65)
  fmt.Println()
  bitopts.Print(18446462598732840960, 0);
}
