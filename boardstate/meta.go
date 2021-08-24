package boardstate

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

const (
  TURN = iota
)

const (
  CASTLE_SHORT = iota
  CASTLE_LONG
)


func (b *BoardState) GetTurn() uint8 {
  if testBit(b.meta, TURN) {
    return BLACK
  } else {
    return WHITE
  }
}

func (b *BoardState) SetTurn(color uint8) {
  if color == WHITE {
    b.meta = clearBit(b.meta, TURN)
  } else {
    b.meta = setBit(b.meta, TURN)
  }
}

func (b *BoardState) ToggleTurn() {
  b.meta = flipBit(b.meta, TURN)
}

func castleBit(color uint8, side uint8) uint8 {
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
  b.meta = b.meta & 0b1111111111111111111111111111111111111111111111111110000000011111
}

// SetEnpassant takes a file 0-7 and saves the enpassant state.
func (b *BoardState) SetEnpassant(file uint8) {
  b.ClearEnpassant()
  b.meta = setBit(b.meta, 5 + file)
}

// SetEnpassant takes a file 0-7 and saves the enpassant state.
func (b *BoardState) IsEnpassant(file uint8) bool {
  return testBit(b.meta, 5 + file)
}



func (b *BoardState) HasCastleRights(color uint8, side uint8) bool {
  return !testBit(b.meta, castleBit(color, side))
}

func (b *BoardState) SetCastleRights(color uint8, side uint8, enabled bool) {
  bit := castleBit(color, side)
  if enabled {
    b.meta = clearBit(b.meta, bit)
  } else {
    b.meta = setBit(b.meta, bit)
  }
}
