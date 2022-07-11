package boardstate

type MoveList struct {
	Head *MoveListEntry
}

type MoveListEntry struct {
	Move *Move
	Next *MoveListEntry
}

func (m *MoveList) AddMove(move *Move) {
	if m.Head == nil {
		m.Head = &MoveListEntry{Move: move}
	} else {
		m.Head = &MoveListEntry{Move: move, Next: m.Head}
	}
}
