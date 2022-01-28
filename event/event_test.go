package event

import (
	"testing"

	"github.com/GlennJoakimB/rivercrossing/state"
)

func TestMoveItem(t *testing.T) {
	state.ViewState(2, 0, 2, 1, 2)

	wanted := "Kylling"
	MoveItem(wanted)

	bItem := state.GetBoatItem()

	if bItem != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", bItem, wanted)
	}
}

func TestCross(t *testing.T) {
	wanted := "noError"
	got := Cross()
	t.Log(got)

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
