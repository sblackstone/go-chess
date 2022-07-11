package boardstate

type MoveList struct {
	head *MoveListEntry
}

type MoveListEntry struct {
	move *Move
	next *MoveListEntry
}

func (m *MoveList) AddMove(move *Move) {
	if m.head == nil {
		m.head = &MoveListEntry{move: move}
	} else {
		m.head = &MoveListEntry{move: move, next: m.head}
	}
}
