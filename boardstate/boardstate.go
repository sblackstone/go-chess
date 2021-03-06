package boardstate


// gridToLinear maps (i,j) -> n
func gridToLinear(i uint8, j uint8) uint8 {
	return i*8 + j
}

// BoardState contains the state of the Board
type BoardState struct {
	colors []uint64
	pieces []uint64
	state  uint32

	//wpassant   int8 // 8 bits needed
	//bpassant   int8 // 8 bits needed
	//turn       int8 // 1 bit needed
	//wcastle    int8 // 2 bits needed
	//bcastle    int8 // 2 bits needed
}



// Blank returns a blank board with no pieces on it
func Blank() *BoardState {
	b := BoardState{}
	b.colors = []uint64{0, 0}
	b.pieces = []uint64{0, 0, 0, 0, 0, 0}
	return &b
}

// Initial returns a blank board with the initial pieces on it
func Initial() *BoardState {
	var j uint8

	b := Blank()

	backFile := []uint8{ROOK, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, ROOK}
	for j = 0; j < 8; j++ {
		b.SetSquare(0, j, WHITE, backFile[j])
		b.SetSquare(7, j, BLACK, backFile[j])

		b.SetSquare(1, j, WHITE, PAWN)
		b.SetSquare(6, j, BLACK, PAWN)
	}
	return b
}

// ColorOfSquare returns WHITE,BLACK, or EMPTY
func (b *BoardState) ColorOfSquare(n uint8) uint8 {
	if testBit(b.colors[WHITE], n) {
		return WHITE
	}
	if testBit(b.colors[BLACK], n) {
		return BLACK
	}
	return EMPTY
}

// PieceOfSquare t
func (b *BoardState) PieceOfSquare(n uint8) uint8 {
	var i uint8
	for i = 0; i < 8; i++ {
		if testBit(b.pieces[i], n) {
			return i
		}
	}
	return 255
}

// SetSquare blah blah blah
func (b *BoardState) SetSquare(i uint8, j uint8, color uint8, piece uint8) {
	n := gridToLinear(i, j)
	b.colors[color] = setBit(b.colors[color], n)
	b.pieces[piece] = setBit(b.pieces[piece], n)
}
