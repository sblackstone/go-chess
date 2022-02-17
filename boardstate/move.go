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
