package boardstate


type Move struct {
  src            int8
  dst            int8
  promotePiece   int8
}

func CreateMove(src int8, dst int8, promotePiece int8) *Move {
  return &Move{
    src: src,
    dst: dst,
    promotePiece: promotePiece,
  }
}

func (m *Move) Src() int8 {
  return m.src;
}

func (m *Move) Dst() int8 {
  return m.src;
}


func (m *Move) PromotePiece() int8 {
  return m.promotePiece;
}
