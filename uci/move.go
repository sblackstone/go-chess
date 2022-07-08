package uci

import (
	"fmt"

	"github.com/sblackstone/go-chess/bitopts"
	"github.com/sblackstone/go-chess/boardstate"
)

func MoveFromUCI(uciStr string) (*boardstate.Move, error) {
	if len(uciStr) > 5 || len(uciStr) < 4 {
		return nil, fmt.Errorf("invalid UCI str: %s", uciStr)
	}
	src, err := bitopts.AlgebraicToSquare(uciStr[0:2])
	if err != nil {
		return nil, err
	}
	dst, err := bitopts.AlgebraicToSquare(uciStr[2:4])
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
