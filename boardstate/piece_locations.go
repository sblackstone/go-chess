package boardstate

type PieceLocations struct {
	pieces [2][6][]int8
}

func (pl *PieceLocations) Copy() PieceLocations {
	var result PieceLocations
	var color, piece int8
	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			result.pieces[color][piece] = make([]int8, len(pl.pieces[color][piece]))
			copy(result.pieces[color][piece], pl.pieces[color][piece])
		}
	}
	return result
}

func (pl *PieceLocations) AddPieceLocation(color, piece, location int8) {
	pl.pieces[color][piece] = append(pl.pieces[color][piece], location)
}

func removeValue(s []int8, val int8) []int8 {
	ret := make([]int8, 0)
	for index, v := range s {
		if v == val {
			ret = append(ret, s[:index]...)
			return append(ret, s[index+1:]...)
		}
	}
	return ret
}

func (pl *PieceLocations) RemovePieceLocation(color, piece, location int8) {
	pl.pieces[color][piece] = removeValue(pl.pieces[color][piece], location)
}

func (pl *PieceLocations) GetLocations(color, piece int8) []int8 {
	return pl.pieces[color][piece]
}
