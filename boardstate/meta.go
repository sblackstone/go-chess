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

const NO_ENPASSANT = 127

const (
	TURN = iota
)

const (
	CASTLE_SHORT = iota
	CASTLE_LONG
)

func (b *BoardState) GetHalfMoves() int {
	return b.halfMoves
}

func (b *BoardState) GetFullMoves() int {
	return b.fullMoves
}

func (b *BoardState) SetFullMoves(movesCount int) {
	b.fullMoves = movesCount
}

func (b *BoardState) SetHalfMoves(movesCount int) {
	b.halfMoves = movesCount
}

func (b *BoardState) GetTurn() int8 {
	return b.turn
}

func (b *BoardState) EnemyColor() int8 {
	return b.turn ^ 1
}

func (b *BoardState) SetTurn(color int8) {
	b.turn = color
}

func (b *BoardState) ToggleTurn() {
	b.turn = b.turn ^ 1
}

func (b *BoardState) ClearEnpassant() {
	b.enpassantSquare = NO_ENPASSANT
}

func (b *BoardState) GetEnpassant() int8 {
	return b.enpassantSquare
}

// SetEnpassant takes a file 0-7 and saves the enpassant state.
func (b *BoardState) SetEnpassant(file int8) {
	b.enpassantSquare = file
}

// IsEnpassant takes a file 0-7 and returns the enpassant state.
func (b *BoardState) IsEnpassant(file int8) bool {
	return b.enpassantSquare == file
}

func (b *BoardState) HasCastleRights(color int8, side int8) bool {
	return b.castleData[color][side]
}

// Only used in testing, removes all castling rights.
func (b *BoardState) ClearCastling() {
	b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	b.SetCastleRights(WHITE, CASTLE_LONG, false)
	b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	b.SetCastleRights(BLACK, CASTLE_LONG, false)
}

func (b *BoardState) EnableAllCastling() {
	b.SetCastleRights(WHITE, CASTLE_SHORT, true)
	b.SetCastleRights(WHITE, CASTLE_LONG, true)
	b.SetCastleRights(BLACK, CASTLE_SHORT, true)
	b.SetCastleRights(BLACK, CASTLE_LONG, true)
}

func (b *BoardState) SetCastleRights(color int8, side int8, enabled bool) {
	b.castleData[color][side] = enabled
}
