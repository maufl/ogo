package openflow10

import (
	"encoding/binary"
	"fmt"
	"net"
)

// ofp_phy_port 1.0
type PhyPort struct {
	PortNo uint16
	HWAddr net.HardwareAddr
	Name   []byte // Size 16

	Config uint32
	State  uint32

	Curr       uint32
	Advertised uint32
	Supported  uint32
	Peer       uint32
}

func NewPhyPort() *PhyPort {
	p := new(PhyPort)
	p.HWAddr = make([]byte, ETH_ALEN)
	p.Name = make([]byte, 16)
	return p
}

func (p PhyPort) String() string {
	return fmt.Sprintf("PhyPort{ PortNo: %d, HWAddr: %s, Name: %s, Config: %b, State: %b, Curr: %b, Advertised: %b, Supported: %b, Peer: %b }",
		p.PortNo, p.HWAddr, string(p.Name), p.Config, p.State, p.Curr, p.Advertised, p.Supported, p.Peer)
}

func (p *PhyPort) Len() (n uint16) {
	n += 2
	n += uint16(len(p.HWAddr) + len(p.Name))
	n += 24
	return
}

func (p *PhyPort) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(p.Len()))
	binary.BigEndian.PutUint16(data, p.PortNo)
	n := 2

	copy(data[n:], p.HWAddr)
	n += len(p.HWAddr)
	copy(data[n:], p.Name)
	n += len(p.Name)

	binary.BigEndian.PutUint32(data[n:], p.Config)
	n += 4
	binary.BigEndian.PutUint32(data[n:], p.State)
	n += 4
	binary.BigEndian.PutUint32(data[n:], p.Curr)
	n += 4
	binary.BigEndian.PutUint32(data[n:], p.Advertised)
	n += 4
	binary.BigEndian.PutUint32(data[n:], p.Supported)
	n += 4
	binary.BigEndian.PutUint32(data[n:], p.Peer)
	n += 4
	return
}

func (p *PhyPort) UnmarshalBinary(data []byte) error {
	p.PortNo = binary.BigEndian.Uint16(data)
	n := 2

	copy(p.HWAddr, data[n:n+6])
	n += 6
	copy(p.Name, data[n:n+16])
	n += 16

	p.Config = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.State = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Curr = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Advertised = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Supported = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Peer = binary.BigEndian.Uint32(data[n:])
	n += 4
	return nil
}

const (
	ETH_ALEN          = 6
	MAX_PORT_NAME_LEN = 16
)

// ofp_port_config 1.0
const (
	PC_PORT_DOWN = 1 << 0

	PC_NO_STP  = 1 << 1
	PC_NO_RECV = 1 << 2

	PC_NO_STP_RECV  = 1 << 3
	PC_NO_FLOOD     = 1 << 4
	PC_NO_FWD       = 1 << 5
	PC_NO_PACKET_IN = 1 << 6
)

// ofp_port_state 1.0
const (
	PS_LINK_DOWN = 1 << 0

	PS_STP_LISTEN  = 0 << 8 /* Not learning or relaying frames. */
	PS_STP_LEARN   = 1 << 8 /* Learning but not relaying frames. */
	PS_STP_FORWARD = 2 << 8 /* Learning and relaying frames. */
	PS_STP_BLOCK   = 3 << 8 /* Not part of spanning tree. */
	PS_STP_MASK    = 3 << 8 /* Bit mask for OFPPS_STP_* values. */
)

// ofp_port 1.0
const (
	P_MAX = 0Xff00

	P_IN_PORT = 0xfff8
	P_TABLE   = 0xfff9

	P_NORMAL = 0xfffa
	P_FLOOD  = 0xfffb

	P_ALL        = 0xfffc
	P_CONTROLLER = 0xfffd
	P_LOCAL      = 0xfffe
	P_NONE       = 0xffff
)

// ofp_port_features 1.0
const (
	PF_10MB_HD  = 1 << 0
	PF_10MB_FD  = 1 << 1
	PF_100MB_HD = 1 << 2
	PF_100MB_FD = 1 << 3
	PF_1GB_HD   = 1 << 4
	PF_1GB_FD   = 1 << 5
	PF_10GB_FD  = 1 << 6

	PF_COPPER     = 1 << 7
	PF_FIBER      = 1 << 8
	PF_AUTONEG    = 1 << 9
	PF_PAUSE      = 1 << 10
	PF_PAUSE_ASYM = 1 << 11
)

// END: 10 - 5.2.1
