package movegenerator

import (
	"testing"
)

func TestBlarg(t *testing.T) {
  boards := genKnightMoveBitBoards();
  t.Errorf("%+v\n\n",boards)
}
