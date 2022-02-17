package boardstate

import("testing")


func TestCreateMove(t *testing.T) {
  m := CreateMove(1,2,3)
  if (m.Src() != 1 || m.Dst() != 2 || m.PromotePiece() != 3) {
    t.Errorf("Create Move didnt setup correctly: %v, %v %v", m.Src(), m.Dst(), m.PromotePiece())
  }
}
