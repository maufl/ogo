package openflowxx

import (
	"encoding/hex"
	"testing"
)

func TestHeaderMarshal(t *testing.T) {
	reference := "0100000800000003"
	h := &Header{Version: 1, Type: 0, Length: 8, Xid: 3}
	serialized, err := h.MarshalBinary()
	if err != nil {
		t.Errorf("Header.MarshalBinary returned an error: %s", err)
	}
	str := hex.EncodeToString(serialized)
	if str != reference {
		t.Errorf("Expected header to serialize to %s but received %s", reference, str)
	}
}

func TestHeaderMarshalUnmarshal(t *testing.T) {
	h := NewHeader(1, 2)
	marshaled, err := h.MarshalBinary()
	if err != nil {
		t.Errorf("Header.MarshalBinary returned an error: %s", err)
	}
	unmarshaled := NewHeader(0, 0)
	if err = unmarshaled.UnmarshalBinary(marshaled); err != nil {
		t.Errorf("Header.UnmarshalBinary returned an error: %s", err)
	}
	if h.Version != unmarshaled.Version ||
		h.Type != unmarshaled.Type ||
		h.Length != unmarshaled.Length ||
		h.Xid != unmarshaled.Xid {
		t.Errorf("Expected unmarshald header to be %+v, got %+v", h, unmarshaled)
	}
}
