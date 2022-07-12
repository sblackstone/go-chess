package boardstate

func (b *BoardState) updateCastlingRights(src int8, dst int8) {
	if src == 4 {
		b.SetCastleRights(WHITE, CASTLE_LONG, false)
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}

	if src == 60 {
		b.SetCastleRights(BLACK, CASTLE_LONG, false)
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}

	if src == 7 || dst == 7 {
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}
	if src == 0 || dst == 0 {
		b.SetCastleRights(WHITE, CASTLE_LONG, false)
	}
	if src == 63 || dst == 63 {
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}
	if src == 56 || dst == 56 {
		b.SetCastleRights(BLACK, CASTLE_LONG, false)
	}
}

func (b *BoardState) handleEnpassant(src int8, dst int8) {
	diff := dst - src

	// Flags
	if diff == 16 {
		b.SetEnpassant(dst - 8)
	} else if diff == -16 {
		b.SetEnpassant(dst + 8)
	} else {
		b.ClearEnpassant()
	}

	if diff == 7 || diff == 9 || diff == -7 || diff == -9 {
		targetPiece := b.PieceOfSquare(dst)
		if targetPiece == EMPTY {
			if dst > src {
				b.SetSquare(dst-8, EMPTY, EMPTY)
			} else {
				b.SetSquare(dst+8, EMPTY, EMPTY)
			}
		}
	}
}

func (b *BoardState) handleCastling(src int8, dst int8) {
	if src-dst == 2 {
		b.MovePiece(src-4, src-1)
	}
	if src-dst == -2 {
		b.MovePiece(src+3, src+1)
	}
}

func (b *BoardState) handleUncastling(src int8, dst int8) {
	if src-dst == 2 {
		b.MovePiece(src-1, src-4)
	}
	if src-dst == -2 {
		b.MovePiece(src+1, src+3)
	}
}

func (b *BoardState) UnplayTurn() {
	msd := b.moveStack
	b.moveStack = msd.prev
	b.enpassantSquare = msd.enpassantSquare
	b.castleData = msd.castleData
	b.halfMoves = msd.halfMoves
	b.MovePiece(msd.dst, msd.src)

	turnAfterMove := b.GetTurn()

	// Deal with un-promotion.
	// TODO: This is an expensive way to do this.
	if msd.srcPiece == PAWN {
		b.SetSquare(msd.src, b.EnemyColor(), PAWN)
		if msd.dst == msd.enpassantSquare {
			if msd.dst > msd.src {
				b.SetSquare(msd.dst-8, turnAfterMove, PAWN)
			} else {
				b.SetSquare(msd.dst+8, turnAfterMove, PAWN)

			}
		}
	}

	if msd.dstPiece != EMPTY {
		b.SetSquare(msd.dst, turnAfterMove, msd.dstPiece)
	}

	if b.GetTurn() == WHITE {
		b.fullMoves -= 1
	}

	if msd.srcPiece == KING {
		b.handleUncastling(msd.src, msd.dst)
	}
	// handle uncastling...

	b.ToggleTurn()
}

func (b *BoardState) PlayTurn(src int8, dst int8, promotePiece int8) {
	// color := b.ColorOfSquare(src)
	piece := b.PieceOfSquare(src)
	dstPiece := b.PieceOfSquare(dst)
	dstColor := b.ColorOfSquare(dst)

	msd := &MoveStackData{
		src:             src,
		srcPiece:        piece,
		dst:             dst,
		dstPiece:        dstPiece,
		enpassantSquare: b.enpassantSquare,
		castleData:      b.castleData,
		halfMoves:       b.halfMoves,
		prev:            b.moveStack,
	}
	b.moveStack = msd

	// Set or Clear enpassant flag
	if piece == PAWN {
		b.handleEnpassant(src, dst)
	} else {
		b.ClearEnpassant()
	}

	b.updateCastlingRights(src, dst)

	b.MovePiece(src, dst)

	// Handle castling
	if piece == KING {
		b.handleCastling(src, dst)
	}

	// Handle Piece promotion
	if promotePiece != EMPTY {
		b.SetSquare(dst, b.GetTurn(), promotePiece)
	}

	if b.GetTurn() == BLACK {
		b.fullMoves += 1
	}

	if piece == PAWN || dstColor != EMPTY {
		b.halfMoves = 0
	} else {
		b.halfMoves += 1
	}

	b.ToggleTurn()

}

func (b *BoardState) PlayTurnFromMove(m *Move) {
	b.PlayTurn(m.Src, m.Dst, m.PromotePiece)
}

func (b *BoardState) CopyPlayTurnFromMove(m *Move) *BoardState {
	return b.CopyPlayTurn(m.Src, m.Dst, m.PromotePiece)
}

func (b *BoardState) CopyPlayTurn(src int8, dst int8, promotePiece int8) *BoardState {
	bCopy := b.Copy()
	bCopy.PlayTurn(src, dst, promotePiece)
	return bCopy
}
