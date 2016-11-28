package openflow10

import (
	"encoding/hex"
	"github.com/kylelemons/godebug/pretty"
	"reflect"
	"testing"
)

var rawPortMod string = "010f0020aa209b8e00010123456789ab0000000a0000000a0000000f00000000"

func TestPortModMarshaling(t *testing.T) {
	t.Parallel()
	portMod := NewPortMod()
	portMod.Header.Xid = 0xaa209b8e
	portMod.PortNo = 1
	portMod.HWAddr = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}
	portMod.Config = PC_NO_STP | PC_NO_STP_RECV
	portMod.Mask = PC_NO_STP | PC_NO_STP_RECV
	portMod.Advertise = PF_10MB_HD | PF_10MB_FD | PF_100MB_HD | PF_100MB_FD

	data, err := portMod.MarshalBinary()
	if err != nil {
		t.Fatalf("Encountered error while marshaling port mod message: %s", err)
	}
	newPortMod := NewPortMod()
	if err = newPortMod.UnmarshalBinary(data); err != nil {
		t.Errorf("Encountered error while unmarshaling port mod message: %s", err)
	}
	if !reflect.DeepEqual(portMod, newPortMod) {
		t.Errorf("Unmarshaling does not reproduce original port mod message\n%s", pretty.Compare(portMod, newPortMod))
	}
	preset, err := hex.DecodeString(rawPortMod)
	if err != nil {
		t.Fatalf("Unable to decode testcase: %s", err)
	}
	err = newPortMod.UnmarshalBinary(preset)
	if err != nil {
		t.Fatalf("Unable to unmarshal testcase: %s", err)
	}
	if !reflect.DeepEqual(portMod, newPortMod) {
		t.Errorf("Unmarshaling of testcase does not result in expected port mod message\n%s", pretty.Compare(portMod, newPortMod))
	}
}
