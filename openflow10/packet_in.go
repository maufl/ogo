package openflow10

import (
	"encoding/binary"
	"github.com/maufl/openflow/openflowxx"
)

// ofp_packet_in 1.0
type PacketIn struct {
	*openflowxx.Header
	BufferId uint32
	TotalLen uint16
	InPort   uint16
	Reason   uint8
	pad      uint8
	Data     openflowxx.Buffer
}

func NewPacketIn() *PacketIn {
	return &PacketIn{
		Header:   openflowxx.NewHeader(VERSION, Type_PacketIn),
		BufferId: 0xffffffff,
		InPort:   P_NONE,
	}
}

func (p *PacketIn) Len() (n uint16) {
	n += p.Header.Len()
	n += 10
	n += p.Data.Len()
	return
}

func (p *PacketIn) MarshalBinary() (data []byte, err error) {
	data, err = p.Header.MarshalBinary()

	b := make([]byte, 10)
	n := 0
	binary.BigEndian.PutUint32(b, p.BufferId)
	n += 4
	binary.BigEndian.PutUint16(b[n:], p.TotalLen)
	n += 2
	binary.BigEndian.PutUint16(b[n:], p.InPort)
	n += 2
	b[n] = p.Reason
	n += 1
	b[n] = p.pad
	n += 1
	data = append(data, b...)

	b, err = p.Data.MarshalBinary()
	data = append(data, b...)
	return
}

func (p *PacketIn) UnmarshalBinary(data []byte) error {
	err := p.Header.UnmarshalBinary(data)
	n := p.Header.Len()

	p.BufferId = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.TotalLen = binary.BigEndian.Uint16(data[n:])
	n += 2
	p.InPort = binary.BigEndian.Uint16(data[n:])
	n += 2
	p.Reason = data[n]
	n += 1
	p.pad = data[n]
	n += 1

	err = p.Data.UnmarshalBinary(data[n:])
	return err
}

// ofp_packet_in_reason 1.0
const (
	R_NO_MATCH = iota
	R_ACTION
)
