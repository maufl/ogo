package openflow10

import (
	"encoding/hex"
	"github.com/kylelemons/godebug/pretty"
	"reflect"
	"testing"
)

var rawEchoRequest string = "0102000800123456"

func TestEchoRequestMarshaling(t *testing.T) {
	t.Parallel()

	echo := NewEchoRequest()
	echo.Header.Xid = 0x123456
	data, err := echo.MarshalBinary()
	if err != nil {
		t.Fatalf("Unable to marshal echo message: %s", err)
	}
	newEcho := NewEchoRequest()
	err = newEcho.UnmarshalBinary(data)
	if err != nil {
		t.Fatalf("Unable to unmarshal echo message: %s", err)
	}
	if !reflect.DeepEqual(echo, newEcho) {
		t.Errorf("Unmarshal does not reproduce original echo message\n", pretty.Compare(echo, newEcho))
	}
	preset, err := hex.DecodeString(rawEchoRequest)
	if err != nil {
		t.Fatalf("Unable to decode testcase: %s", err)
	}
	err = newEcho.UnmarshalBinary(preset)
	if err != nil {
		t.Fatalf("Unable to unmarshal testcase: %s", err)
	}
	if !reflect.DeepEqual(echo, newEcho) {
		t.Errorf("Unmarshal testcase does not produce expected echo message\n", pretty.Compare(echo, newEcho))
	}
}
