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

// Initial returns a board with the initial setup.
func Initial() *BoardState {
	b := BoardState{}
	// These constants are pre-calculated for the initial board state.
	b.colors = []uint64{65535, 18446462598732840960}
	b.pieces = []uint64{9295429630892703873, 4755801206503243842, 2594073385365405732, 576460752303423496, 1152921504606846992, 71776119061282560}
	return &b
}

// InitialManual sets up the board manually, only used to calculate the constants for the fast version Initial.
func InitialManual() *BoardState {
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
	return EMPTY
}

// SetSquare blah blah blah
func (b *BoardState) SetSquare(i uint8, j uint8, color uint8, piece uint8) {
	n := gridToLinear(i, j)
	b.colors[color] = setBit(b.colors[color], n)
	b.pieces[piece] = setBit(b.pieces[piece], n)
}
