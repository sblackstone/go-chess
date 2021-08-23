package boardstate

/*
//turn       int8 // 1 bit needed
//wcastle    int8 // 2 bits needed
//bcastle    int8 // 2 bits needed
//wpassant   int8 // 8 bits needed
//bpassant   int8 // 8 bits needed

*/

const (
  TURN = iota
  WCASTLE_SHORT
  WCASTLE_LONG
  BCASTLE_SHORT
  BCASTLE_LONG
)

func (b *BoardState) getTurn() uint8 {
  if testBit(b.meta, TURN) {
    return WHITE
  } else {
    return BLACK
  }
}

func (b *BoardState) setTurn(color uint8) {
  if color == BLACK {
    b.meta = clearBit(b.meta, TURN)
  } else {
    b.meta = setBit(b.meta, TURN)
  }
}

func (b *BoardState) toggleTurn() {
  b.meta = flipBit(b.meta, TURN)
}

/*
func (b *BoardState) clearCastleRight(color uint8, side uint8) {

}
*/
