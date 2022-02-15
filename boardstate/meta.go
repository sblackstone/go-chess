package boardstate

import ("github.com/sblackstone/go-chess/bitopts")
/*

// All fields on an initial board should be 0 so no initializiation is necessary.

Bit
0      Turn                      (White = 0, Black = 1)
1      White Castling Short      (Available = 0, Unavailalbe = 1)
2      White Castling Long
3      Black Castling Short
4      Black Castling Long
5-12   Enpassant state after last move  (Set = pawn pushed in that file by opposing color)

*/

const NO_ENPASSANT = 127

const (
  TURN = iota
)

const (
  CASTLE_SHORT = iota
  CASTLE_LONG
)


func (b *BoardState) GetTurn() int8 {
  return b.turn
}

func (b *BoardState) SetTurn(color int8) {
  b.turn = color
}

func (b *BoardState) ToggleTurn() {
  b.turn = b.turn ^ 1
}

func castleBit(color int8, side int8) int8 {
  // WHITE = 0, BLACK = 1
  // SHORT = 0 LONG = 1
  // So,
  // WHITE SHORT = 1 + (2*0) + 0 = 1
  // WHITE LONG  = 1 + (2*0) + 1 = 2
  // BLACK SHORT = 1 + (2*1) + 0 = 3
  // BLACK LONG  = 1 + (2*2) + 1 = 4
  return  1 + (color * 2) + side;
}

func (b *BoardState) ClearEnpassant() {
  b.enpassantFile = NO_ENPASSANT;
}

func (b *BoardState) GetEnpassant() int8 {
  return b.enpassantFile;
}


// SetEnpassant takes a file 0-7 and saves the enpassant state.
func (b *BoardState) SetEnpassant(file int8) {
  b.enpassantFile = file;
}

// IsEnpassant takes a file 0-7 and returns the enpassant state.
func (b *BoardState) IsEnpassant(file int8) bool {
  return b.enpassantFile == file
}

func (b *BoardState) HasCastleRights(color int8, side int8) bool {
  return !bitopts.TestBit(b.meta, castleBit(color, side))
}

func (b *BoardState) SetCastleRights(color int8, side int8, enabled bool) {
  bit := castleBit(color, side)
  if enabled {
    b.meta = bitopts.ClearBit(b.meta, bit)
  } else {
    b.meta = bitopts.SetBit(b.meta, bit)
  }
}
