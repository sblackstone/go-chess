package movegenerator

import (
	"sort"

	"github.com/sblackstone/go-chess/bitops"
	"github.com/sblackstone/go-chess/boardstate"

	"reflect"
	"testing"
)

func genSortedBoardLocationsGeneric(turn int8, piece int8, result []*boardstate.BoardState) []int8 {
	var locations []int8
	for i := range result {
		locations = append(locations, result[i].FindPieces(turn, piece)...)
	}
	sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
	return locations
}

func testSuccessorsHelper(t *testing.T, b *boardstate.BoardState, pieceType int8, expected []int8) {
	successors := genSuccessorsForPiece(b, pieceType)
	locations := genSortedBoardLocationsGeneric(b.GetTurn(), pieceType, successors)
	if !reflect.DeepEqual(locations, expected) {
		t.Errorf("Expected\n%v\nto be\n%v", locations, expected)
	}
}

func testAttacksHelper(t *testing.T, b *boardstate.BoardState, pieceType int8, expected []int8) {
	attacks := genAttacksForPiece(b, b.GetTurn(), pieceType)

	var locations []int8
	var i int8
	for i = 0; i < 64; i++ {
		if bitops.TestBit(attacks, i) {
			locations = append(locations, i)
		}
	}

	if !reflect.DeepEqual(locations, expected) {
		t.Errorf("Expected\n%v\nto be\n%v", locations, expected)
	}
}
