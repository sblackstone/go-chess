package uci

import (
	"strings"

	"github.com/sblackstone/go-chess/boardstate"
)

func RemovePositionPrefix(positionStr string) string {
	if strings.HasPrefix(positionStr, "position") {
		positionStr = positionStr[8:]
	}
	return strings.Trim(positionStr, " ")

}

func BoardFromUCIPosition(positionStr string) *boardstate.BoardState {
	positionStr = RemovePositionPrefix(positionStr)
	return boardstate.Initial()
}
