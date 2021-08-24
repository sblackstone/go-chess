package boardstate

/*


Bit
1      Turn                      (White = 0, Black = 1)
2      White Castling Short      ()


//turn       int8 // 1 bit needed

  SET   = BLACK
  UNSET = WHITE

//wcastle    int8 // 2 bits needed
//bcastle    int8 // 2 bits needed
//wpassant   int8 // 8 bits needed
//bpassant   int8 // 8 bits needed

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
