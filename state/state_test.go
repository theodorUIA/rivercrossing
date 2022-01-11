package state

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "[Kylling rev korn mann ---\\ \\__/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil i ViewState: %s != %s", state, wanted)
	}
}
