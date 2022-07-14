package uci

import (
	"fmt"

	"github.com/sblackstone/go-chess/boardstate"
)

func MoveFromUCI(uciStr string) (*boardstate.Move, error) {
	if len(uciStr) > 5 || len(uciStr) < 4 {
		return nil, fmt.Errorf("invalid UCI str: %s", uciStr)
	}
	src, err := boardstate.AlgebraicToSquare(uciStr[0:2])
	if err != nil {
		return nil, err
	}
	dst, err := boardstate.AlgebraicToSquare(uciStr[2:4])
	if err != nil {
		return nil, err
	}

	if src == dst {
		return nil, fmt.Errorf("src and dst were both %v", src)
	}
	promotePiece := int8(boardstate.EMPTY)
	if len(uciStr) == 5 {
		switch uciStr[4:] {
		case "r":
			promotePiece = boardstate.ROOK
		case "n":
			promotePiece = boardstate.KNIGHT
		case "b":
			promotePiece = boardstate.BISHOP
		case "q":
			promotePiece = boardstate.QUEEN
		default:
			return nil, fmt.Errorf("invalid UCI str: %s", uciStr)
		}
	}

	return &boardstate.Move{Src: src, Dst: dst, PromotePiece: promotePiece}, nil

}

func MoveToUCI(move *boardstate.Move) string {
	src := boardstate.SquareToAlgebraic(move.Src)
	dst := boardstate.SquareToAlgebraic(move.Dst)

	promotionPiece := ""

	switch move.PromotePiece {
	case boardstate.QUEEN:
		promotionPiece = "q"
	case boardstate.KNIGHT:
		promotionPiece = "n"
	case boardstate.BISHOP:
		promotionPiece = "b"
	case boardstate.ROOK:
		promotionPiece = "r"
	}

	return src + dst + promotionPiece

}
