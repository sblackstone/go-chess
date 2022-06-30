package boardstate


type Move struct {
  Src            int8
  Dst            int8
  PromotePiece   int8
}

func CreateMove(src int8, dst int8, promotePiece int8) *Move {
  return &Move{
    Src: src,
    Dst: dst,
    PromotePiece: promotePiece,
  }
}
