package boardstate

type PieceLocations = PieceLocationsLinkedList

type PieceLocationsSlice struct {
	pieces [2][6][]int8
}

func (pl *PieceLocationsSlice) Init() {
	var color, piece int8
	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			pl.pieces[color][piece] = make([]int8, 0, 20)
		}
	}

}

func (pl *PieceLocationsSlice) Copy() PieceLocationsSlice {
	var result PieceLocationsSlice
	var color, piece int8
	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			result.pieces[color][piece] = make([]int8, len(pl.pieces[color][piece]))
			copy(result.pieces[color][piece], pl.pieces[color][piece])
		}
	}
	return result
}

func (pl *PieceLocationsSlice) AddPieceLocation(color, piece, location int8) {
	pl.pieces[color][piece] = append(pl.pieces[color][piece], location)
}

// Warning: We do not preserve the slice here..
// Peformance increase is huge!
func removeValue(s []int8, val int8) []int8 {
	for i, v := range s {
		if v == val {
			s[i] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	return s
}

func (pl *PieceLocationsSlice) RemovePieceLocation(color, piece, location int8) {
	pl.pieces[color][piece] = removeValue(pl.pieces[color][piece], location)
}

func (pl *PieceLocationsSlice) GetLocations(color, piece int8) []int8 {
	return pl.pieces[color][piece]
}

func (pl *PieceLocationsSlice) EachLocation(color, piece int8, f func(int8)) {
	for _, l := range pl.pieces[color][piece] {
		f(l)
	}
}
