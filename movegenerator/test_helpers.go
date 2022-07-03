package movegenerator

import(
  "sort"
  "github.com/sblackstone/go-chess/boardstate"
)
func genSortedBoardLocationsGeneric(turn int8, piece int8, result[]*boardstate.BoardState) []int8 {
	var locations []int8
  for i := range(result) {
    locations = append(locations, result[i].FindPieces(turn, piece)...)
  }
  sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
  return locations
}

func testSuccessorsHelper(t *testing.T, b *boardstate.BoardState, pieceType int8, expected []int8) {
		successors := genSuccessorsForPiece(b, pieceType)
		locations  := genSortedBoardLocationsGeneric(b.GetTurn(), pieceType, successors)
		if !reflect.DeepEqual(locations, expected) {
	    t.Errorf("Expected %v to be %v", locations, expected)
	  }
}


// DEPRECATED BELOW HERE.
func genSortedBoardLocationsRooks(b *boardstate.BoardState) []int8 {
	return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.ROOK, genRookSuccessors(b))
}

func genSortedBoardLocationsKnights(b *boardstate.BoardState) []int8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KNIGHT, genKnightSuccessors(b))
}

func genSortedBoardLocationsBishops(b *boardstate.BoardState) []int8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.BISHOP, genBishopSuccessors(b))
}

func genSortedBoardLocationsQueens(b *boardstate.BoardState) []int8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.QUEEN, genQueenSuccessors(b))
}

func genSortedBoardLocationsKings(b *boardstate.BoardState) []int8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.KING, genKingSuccessors(b))
}

func genSortedBoardLocationsPawns(b *boardstate.BoardState) []int8 {
  return genSortedBoardLocationsGeneric(b.GetTurn(), boardstate.PAWN, genPawnSuccessors(b))
}

//
// func genSortedCheckedSquares(b *boardstate.BoardState, color int8) uint64 {
//   return GenAllCheckedSquares(b, color)
// }
