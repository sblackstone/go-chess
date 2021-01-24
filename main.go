package main

import (
  "fmt"
  "github.com/sblackstone/go-chess/boardstate"
)


func main() {
  b := boardstate.BlankBoard();
  fmt.Println(b.GetWhiteSquares());
	fmt.Println("Hello, world.")
}
