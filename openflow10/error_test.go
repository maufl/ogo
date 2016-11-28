package openflow10

import (
	"encoding/hex"
	"github.com/kylelemons/godebug/pretty"
	"reflect"
	"testing"
)

var rawError string = "0101000c9876543200010002"

func TestErrorMarshaling(t *testing.T) {
	t.Parallel()

	error := NewError()
	error.Header.Xid = 0x98765432
	error.Type = ET_BAD_REQUEST
	error.Code = BRC_BAD_STAT

	data, err := error.MarshalBinary()
	if err != nil {
		t.Fatalf("Unable to marshal error message: %s", err)
	}
	newError := NewError()
	err = newError.UnmarshalBinary(data)
	if err != nil {
		t.Fatalf("Unable to unmarshal error message: %s", err)
	}
	if !reflect.DeepEqual(error, newError) {
		t.Errorf("Unmarshaled error message does not reproduce original error message\n", pretty.Compare(error, newError))
	}

	preset, err := hex.DecodeString(rawError)
	if err != nil {
		t.Fatalf("Unable to decode testcase: %s", err)
	}
	err = newError.UnmarshalBinary(preset)
	if err != nil {
		t.Fatalf("Unable to unmarshal testcase: %s", err)
	}
	if !reflect.DeepEqual(error, newError) {
		t.Errorf("Unmarshaling testcase does not produce expected error message\n", pretty.Compare(error, newError))
	}
}
