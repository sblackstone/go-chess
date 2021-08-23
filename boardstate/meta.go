package boardstate

/*


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


func (b *BoardState) getTurn() uint8 {
  if testBit(b.meta, TURN) {
    return BLACK
  } else {
    return WHITE
  }
}

func (b *BoardState) setTurn(color uint8) {
  if color == WHITE {
    b.meta = clearBit(b.meta, TURN)
  } else {
    b.meta = setBit(b.meta, TURN)
  }
}

func (b *BoardState) toggleTurn() {
  b.meta = flipBit(b.meta, TURN)
}


func (b *BoardState) canCastle(side uint8, color uint8) bool {
  bit := 1 + (color * 2) + side
  return testBit(b.meta, bit)
}

func (b *BoardState) setCastleState(side uint8, color uint8, enabled bool) {

  // WHITE = 0, BLACK = 1
  // SHORT = 0 LONG = 1
  // So,
  // WHITE SHORT = 1 + (2*0) + 0 = 1
  // WHITE LONG  = 1 + (2*0) + 1 = 2
  // BLACK SHORT = 1 + (2*1) + 0 = 3
  // BLACK LONG  = 1 + (2*2) + 1 = 4

  bit := 1 + (color * 2) + side
  if enabled {
    b.meta = setBit(b.meta, bit)
  } else {
    b.meta = clearBit(b.meta, bit)
  }
}



/*
func (b *BoardState) clearCastleRight(color uint8, side uint8) {

}
*/
