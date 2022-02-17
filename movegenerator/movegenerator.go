package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)

func GenSucessors(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
  result = append(result, genPawnMoves(b)...);
  result = append(result, genKingMoves(b)...);
  result = append(result, genQueenSuccessors(b)...);
  result = append(result, genBishopSuccessors(b)...);
  result = append(result, genKnightMoves(b)...);
  result = append(result, genRookSuccessors(b)...);
  return result;
}
