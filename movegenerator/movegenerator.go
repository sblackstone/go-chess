package movegenerator

import (
	"github.com/sblackstone/go-chess/boardstate"
)












func genSucessors(b *boardstate.BoardState) []*boardstate.BoardState {
  var result []*boardstate.BoardState;
  result = append(result, genPawnMoves(b)...);
  result = append(result, genKingMoves(b)...);
  result = append(result, genQueenMoves(b)...);
  result = append(result, genBishopMoves(b)...);
  result = append(result, genKnightMoves(b)...);
  result = append(result, genRookMoves(b)...);
  return result;
}
