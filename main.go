package main

import (
  "github.com/sblackstone/go-chess/boardstate"
)


func main() {
  b := boardstate.Initial();
  b.Print()
}
