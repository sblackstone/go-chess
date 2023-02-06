package boardstate

type LocationNode struct {
	next     *LocationNode
	location int8
}

type PieceLocationsLinkedList struct {
	pieces [2][6]*LocationNode
}

func (pl *PieceLocationsLinkedList) Init() {
}

func (pl *PieceLocationsLinkedList) Copy() PieceLocationsLinkedList {
	var result PieceLocationsLinkedList
	var color, piece int8
	for color = WHITE; color <= BLACK; color++ {
		for piece = ROOK; piece <= PAWN; piece++ {
			for cur := pl.pieces[color][piece]; cur != nil; cur = cur.next {
				result.AddPieceLocation(color, piece, cur.location)
			}
		}
	}
	return result
}

func (pl *PieceLocationsLinkedList) AddPieceLocation(color, piece, location int8) {
	newNode := &LocationNode{
		location: location,
	}
	if pl.pieces[color][piece] == nil {
		pl.pieces[color][piece] = newNode
	} else {
		newNode.next = pl.pieces[color][piece]
		pl.pieces[color][piece] = newNode
	}
}

func (pl *PieceLocationsLinkedList) RemovePieceLocation(color, piece, location int8) {
	if pl.pieces[color][piece].location == location {
		pl.pieces[color][piece] = pl.pieces[color][piece].next
	} else {
		for cur := pl.pieces[color][piece]; cur.next != nil; cur = cur.next {
			if cur.next.location == location {
				cur.next = cur.next.next
				break
			}
		}
	}
}

func (pl *PieceLocationsLinkedList) GetLocations(color, piece int8) []int8 {
	var ret []int8
	for cur := pl.pieces[color][piece]; cur != nil; cur = cur.next {
		ret = append(ret, cur.location)
	}
	return ret
}

func (pl *PieceLocationsLinkedList) Each(color int8, piece int8, callback func(int8)) {
	for cur := pl.pieces[color][piece]; cur != nil; cur = cur.next {
		callback(cur.location)
	}
}
