package openflow10

import (
	"testing"
)

func TestActionOutputEqual(t *testing.T) {
	a := NewActionOutput(8)
	b := NewActionOutput(8)
	if !a.Equal(b) {
		t.Errorf("Two output actions with same port are not equal: %s != %s", a, b)
	}
}
