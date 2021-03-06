package openflow10

import (
	"encoding/binary"
	"net"

	"github.com/maufl/openflow/openflowxx"
)

type FeaturesRequest struct {
	*openflowxx.Header
}

// FeaturesRequest constructor
func NewFeaturesRequest() *FeaturesRequest {
	return &FeaturesRequest{openflowxx.NewHeader(VERSION, Type_FeaturesRequest)}
}

type FeaturesReply struct {
	*openflowxx.Header
	DPID         net.HardwareAddr // Size 8
	Buffers      uint32
	Tables       uint8
	pad          []uint8 // Size 3
	Capabilities uint32
	Actions      uint32

	Ports []PhyPort
}

// FeaturesReply constructor
func NewFeaturesReply() *FeaturesReply {
	return &FeaturesReply{
		Header: openflowxx.NewHeader(VERSION, Type_FeaturesReply),
		DPID:   make([]byte, 8),
		pad:    make([]byte, 3),
		Ports:  make([]PhyPort, 0),
	}
}

func (s *FeaturesReply) Len() (n uint16) {
	n = s.Header.Len()
	n += uint16(len(s.DPID))
	n += 16
	for _, p := range s.Ports {
		n += p.Len()
	}
	return
}

func (s *FeaturesReply) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(s.Len()))
	bytes := make([]byte, 0)
	next := 0

	s.Header.Length = s.Len()
	bytes, err = s.Header.MarshalBinary()
	copy(data[next:], bytes)
	next += len(bytes)
	copy(data[next:], s.DPID)
	next += len(s.DPID)
	binary.BigEndian.PutUint32(data[next:], s.Buffers)
	next += 4
	data[next] = s.Tables
	next += 1
	copy(data[next:], s.pad)
	next += len(s.pad)
	binary.BigEndian.PutUint32(data[next:], s.Capabilities)
	next += 4
	binary.BigEndian.PutUint32(data[next:], s.Actions)
	next += 4

	for _, p := range s.Ports {
		bytes, err = p.MarshalBinary()
		if err != nil {
			return
		}
		copy(data[next:], bytes)
		next += len(bytes)
	}
	return
}

func (s *FeaturesReply) UnmarshalBinary(data []byte) error {
	var err error
	next := 0

	err = s.Header.UnmarshalBinary(data[next:])
	next = int(s.Header.Len())
	copy(s.DPID, data[next:])
	next += len(s.DPID)
	s.Buffers = binary.BigEndian.Uint32(data[next:])
	next += 4
	s.Tables = data[next]
	next += 1
	copy(s.pad, data[next:])
	next += len(s.pad)
	s.Capabilities = binary.BigEndian.Uint32(data[next:])
	next += 4
	s.Actions = binary.BigEndian.Uint32(data[next:])
	next += 4

	for next < len(data) {
		p := NewPhyPort()
		err = p.UnmarshalBinary(data[next:])
		s.Ports = append(s.Ports, *p)
		next += int(p.Len())
	}
	return err
}

// ofp_capabilities 1.0
const (
	C_FLOW_STATS   = 1 << 0
	C_TABLE_STATS  = 1 << 1
	C_PORT_STATS   = 1 << 2
	C_STP          = 1 << 3
	C_RESERVED     = 1 << 4
	C_IP_REASM     = 1 << 5
	C_QUEUE_STATS  = 1 << 6
	C_ARP_MATCH_IP = 1 << 7
)
