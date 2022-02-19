package movegenerator

import (
  "testing"
  "github.com/sblackstone/go-chess/fen"
  "fmt"

)

func TestLetsTryIt(t *testing.T) {
  b, err := fen.FromFEN("8/ppp3p1/8/8/3p4/8/1ppp2K1/brk2Q1n b - - 12 7")
  if err != nil {
    t.Errorf("Err: %v", err)
  }

  sucessors := GenSucessors(b)

  for i := range(sucessors) {
    fenStr, err := fen.ToFEN(sucessors[i])
    if err != nil {
      t.Errorf("%v", err)
    }
    fmt.Println(fenStr)
  }

}
