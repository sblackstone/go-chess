package movegenerator

import(
  "sort"
  "github.com/sblackstone/go-chess/boardstate"
)
func genSortedBoardLocationsGeneric(turn uint8, piece uint8, result[]*boardstate.BoardState) []uint8 {
	var locations []uint8
  for i := range(result) {
    locations = append(locations, result[i].FindPieces(turn, piece)...)
  }
  sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
  return locations
}

func genSortedBoardLocationsRooks(b *boardstate.BoardState) []uint8 {
	return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.ROOK, genRookMoves(b))
}

func genSortedBoardLocationsKnights(b *boardstate.BoardState) []uint8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KNIGHT, genKnightMoves(b))
}

func genSortedBoardLocationsBishops(b *boardstate.BoardState) []uint8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.BISHOP, genBishopMoves(b))
}

func genSortedBoardLocationsQueens(b *boardstate.BoardState) []uint8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.QUEEN, genQueenMoves(b))
}

func genSortedBoardLocationsKings(b *boardstate.BoardState) []uint8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, genKingMoves(b))
}

func genSortedBoardLocationsPawns(b *boardstate.BoardState) []uint8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.PAWN, genPawnMoves(b))
}
