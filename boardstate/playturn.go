package boardstate

import (
	//"github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func (b * BoardState) updateCastlingRights(src int8, dst int8) {
	if (src == 4) {
		b.SetCastleRights(WHITE, CASTLE_LONG,  false)
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}

	if (src == 60) {
		b.SetCastleRights(BLACK, CASTLE_LONG,  false)
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}

	if (src == 7 || dst == 7) {
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}
	if (src == 0 || dst == 0) {
		b.SetCastleRights(WHITE, CASTLE_LONG, false)
	}
	if (src == 63 || dst == 63) {
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}
	if (src == 56 || dst == 56) {
		b.SetCastleRights(BLACK, CASTLE_LONG, false)
	}
}


func (b *BoardState) handleEnpassant(src int8, dst int8) {
  diff := dst - src

  // Flags
  if (diff == 16) {
    b.SetEnpassant(dst - 8)
	} else if (diff == -16) {
		b.SetEnpassant(dst + 8)
  } else {
    b.ClearEnpassant()
  }

  if (diff == 7 || diff == 9 || diff == -7 || diff == -9) {
    targetPiece := b.PieceOfSquare(dst)
    if targetPiece == EMPTY {
      if dst > src {
        b.SetSquare(dst - 8, EMPTY, EMPTY)
      } else {
        b.SetSquare(dst + 8, EMPTY, EMPTY)
      }
    }
  }
}

func (b *BoardState) handleCastling(src int8, dst int8) {
  if (src-dst == 2) {
    b.MovePiece(src-4, src-1)
  }
  if (src-dst == -2) {
    b.MovePiece(src+3, src+1)
  }
}

func (b *BoardState) PlayTurn(src int8, dst int8, promotePiece int8) {
	  color := b.ColorOfSquare(src)
		piece := b.PieceOfSquare(src)
		dstPiece := b.PieceOfSquare(dst)
		dstColor := b.ColorOfSquare(dst)

		b.movesData = append(b.movesData, MoveData{
			src: src,
			dst: dst,
			srcColor: color,
			srcPiece: piece,
			dstPiece: dstPiece,
			dstColor: dstColor,
			preMoveEnpassantSquare: b.enpassantSquare,
			preMoveMeta: b.meta,
			preHalfMoves: b.halfMoves,
		})

		// Set or Clear enpassant flag
		if (piece == PAWN) {
      b.handleEnpassant(src, dst)
		} else {
			b.ClearEnpassant()
		}

		b.updateCastlingRights(src, dst)

		b.MovePiece(src, dst)

		// Handle castling
		if (piece == KING) {
      b.handleCastling(src, dst)
		}

		// Handle Piece promotion
		if promotePiece != EMPTY {
			b.SetSquare(dst, b.GetTurn(), promotePiece)
		}


		if b.GetTurn() == BLACK {
			b.IncrementFullMoves()
		}

		if (piece == PAWN || dstColor != EMPTY) {
			b.SetHalfMoves(0)
		} else {
			b.IncrementHalfMoves()
		}

		b.ToggleTurn()

}

func (b *BoardState) PopTurn() {
	idx := len(b.movesData) - 1
  md := b.movesData[idx]
	b.movesData = b.movesData[:idx]

	b.MovePiece(md.dst, md.src)

	b.SetSquare(md.dst, md.dstColor, md.dstPiece)

	b.meta = md.preMoveMeta
	b.enpassantSquare = md.preMoveEnpassantSquare
	b.halfMoves = md.preHalfMoves
	if b.turn == WHITE {
		b.fullMoves -= 1
	}

	b.ToggleTurn()


}

func (b *BoardState) PlayTurnFromMove(m *Move) {
	b.PlayTurn(m.src, m.dst, m.promotePiece)
}

func (b *BoardState) CopyPlayTurnFromMove(m *Move) *BoardState {
	return b.CopyPlayTurn(m.src, m.dst, m.promotePiece)
}


func (b *BoardState) CopyPlayTurn(src int8, dst int8, promotePiece int8) *BoardState{
	bCopy := b.Copy()
	bCopy.PlayTurn(src, dst, promotePiece)
	return bCopy
}
