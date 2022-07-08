package uci

import "testing"

func TestRemovePositionPrefix(t *testing.T) {
	if RemovePositionPrefix("position blarg") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

	if RemovePositionPrefix("position           blarg") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

	if RemovePositionPrefix("  blarg  ") != "blarg" {
		t.Errorf("expected %v to be %v", RemovePositionPrefix("position blarg"), "blarg")
	}

}
