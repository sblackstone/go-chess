package boardstate

type PieceLocations struct {
	pieces [2][6][]int8
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
