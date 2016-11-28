package openflow10

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var rawPhyPort string = "00010123456789ab706f72743100000000000000000000000000000a00000300000004a0000004a0000006c0000008a0"

func TestPhyPortMarshaling(t *testing.T) {
	t.Parallel()
	port := NewPhyPort()
	port.PortNo = 1
	port.HWAddr = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}
	copy(port.Name, []byte("port1"))
	port.Config = PC_NO_STP | PC_NO_STP_RECV
	port.State = PS_STP_BLOCK
	port.Curr = PF_COPPER | PF_PAUSE | PF_1GB_FD
	port.Advertised = PF_COPPER | PF_PAUSE | PF_1GB_FD
	port.Supported = PF_COPPER | PF_PAUSE | PF_AUTONEG | PF_10GB_FD
	port.Peer = PF_COPPER | PF_PAUSE_ASYM | PF_1GB_FD

	data, err := port.MarshalBinary()
	if err != nil {
		t.Fatalf("Encountered error while marshaling physical port: %s", err)
	}
	newPort := NewPhyPort()
	if err = newPort.UnmarshalBinary(data); err != nil {
		t.Errorf("Encountered error while unmarshaling physical port: %s", err)
	}
	if !reflect.DeepEqual(port, newPort) {
		t.Errorf("Unmarshaling does not reproduce original port")
	}
	preset, err := hex.DecodeString(rawPhyPort)
	if err != nil {
		t.Fatalf("Unable to decode testcase: %s", err)
	}
	err = newPort.UnmarshalBinary(preset)
	if err != nil {
		t.Fatalf("Unable to unmarshal testcase: %s", err)
	}
	if !reflect.DeepEqual(port, newPort) {
		t.Errorf("Unmarshaling of testcase does not result in expected port")
	}
}
