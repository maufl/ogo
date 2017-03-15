package openflow10

import (
	_ "log"
	"testing"
)

func TestFlowModMarshaling(t *testing.T) {
	fm := NewFlowMod()
	fm.AddAction(NewActionOutput(1))
	fm.AddAction(NewActionVLANVID(5))
	data, err := fm.MarshalBinary()
	if err != nil {
		t.Errorf("Error while marshaling flow mod: %s", err)
	}
	newFm := NewFlowMod()
	err = newFm.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("Error while unmarshaling flow mod: %s", err)
	}
	if !fm.Equal(newFm) {
		t.Errorf("Flow mod is not equal to unmarshaled self:\n%+v\n%+v", fm, newFm)
	}
}
