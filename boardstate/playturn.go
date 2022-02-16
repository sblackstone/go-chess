package boardstate

import (
	"github.com/sblackstone/go-chess/bitopts"
	//"fmt"
)


func (b * BoardState) updateCastlingRights(src int8) {
	if (src == 3) {
		b.SetCastleRights(WHITE, CASTLE_LONG,  false)
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}

	if (src == 59) {
		b.SetCastleRights(BLACK, CASTLE_LONG,  false)
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}

	if (src == 0) {
		b.SetCastleRights(WHITE, CASTLE_SHORT, false)
	}
	if (src == 7) {
		b.SetCastleRights(WHITE, CASTLE_LONG, false)
	}
	if (src == 56) {
		b.SetCastleRights(BLACK, CASTLE_SHORT, false)
	}
	if (src == 63) {
		b.SetCastleRights(BLACK, CASTLE_LONG, false)
	}
}


func (b *BoardState) handleEnpassant(src int8, dst int8) {
  diff := dst - src

  // Flags
  if (diff == 16 || diff == -16) {
    b.SetEnpassant(bitopts.FileOfSquare(src))
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
    b.MovePiece(src-3, src-1)
  }
  if (src-dst == -2) {
    b.MovePiece(src+4, src+1)
  }
}

func (b *BoardState) PlayTurn(src int8, dst int8, promotePiece int8) {
		// TODO: We can make this only check the pawns for a few ops worth of savings.
		piece := b.PieceOfSquare(src)

		// Set or Clear enpassant flag
		if (piece == PAWN) {
      b.handleEnpassant(src, dst)
		} else {
			b.ClearEnpassant()
		}

		b.updateCastlingRights(src)

		b.MovePiece(src, dst)

		// Handle castling
		if (piece == KING) {
      b.handleCastling(src, dst)
		}

		// Handle Piece promotion
		if promotePiece != EMPTY {
			b.SetSquare(dst, b.GetTurn(), promotePiece)
		}

		b.ToggleTurn()

}

func (b *BoardState) CopyPlayTurn(src int8, dst int8, promotePiece int8) *BoardState{
	bCopy := b.Copy()
	bCopy.PlayTurn(src, dst, promotePiece)
	return bCopy
}
