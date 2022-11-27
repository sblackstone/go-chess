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
	b.UpdateZorbistKey(zorbistTurns[b.turn])
	b.turn = color
	b.UpdateZorbistKey(zorbistTurns[b.turn])
}

func (b *BoardState) ToggleTurn() {
	b.SetTurn(b.turn ^ 1)
}

func (b *BoardState) ClearEnpassant() {
	if b.enpassantSquare != NO_ENPASSANT {
		b.UpdateZorbistKey(zorbistEnpassant[b.enpassantSquare])
	}
	b.enpassantSquare = NO_ENPASSANT
}

func (b *BoardState) GetEnpassant() int8 {
	return b.enpassantSquare
}

// SetEnpassant takes a pos and saves the enpassant state.
func (b *BoardState) SetEnpassant(pos int8) {
	if b.enpassantSquare != NO_ENPASSANT {
		b.UpdateZorbistKey(zorbistEnpassant[b.enpassantSquare])
	}
	b.enpassantSquare = pos
	if pos != NO_ENPASSANT {
		b.UpdateZorbistKey(zorbistEnpassant[pos])
	}
}

// IsEnpassant takes a pos and returns the enpassant state.
func (b *BoardState) IsEnpassant(pos int8) bool {
	return b.enpassantSquare == pos
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
	if (b.castleData[color][side] && !enabled) || (!b.castleData[color][side] && enabled) {
		b.UpdateZorbistKey(zorbistCastling[color][side])
	}
	b.castleData[color][side] = enabled
}
