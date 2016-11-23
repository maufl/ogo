package openflow10

import (
	"encoding/binary"
	"github.com/maufl/openflow/openflowxx"
)

// When the controller wishes to send a packet out through the
// datapath, it uses the OFPT_PACKET_OUT message: The buffer_id
// is the same given in the ofp_packet_in message. If the
// buffer_id is -1, then the packet data is included in the data
// array. If OFPP_TABLE is specified as the output port of an
// action, the in_port in the packet_out message is used in the
// flow table lookup.
type PacketOut struct {
	*openflowxx.Header
	BufferId   uint32
	InPort     uint16
	ActionsLen uint16
	Actions    []Action
	Data       *openflowxx.Buffer
}

func NewPacketOut() *PacketOut {
	return &PacketOut{
		Header:   openflowxx.NewHeader(VERSION, Type_PacketOut),
		BufferId: 0xffffffff,
		InPort:   P_NONE,
		Actions:  make([]Action, 0),
	}
}

func (p *PacketOut) AddAction(act Action) {
	p.Actions = append(p.Actions, act)
	p.ActionsLen += act.Len()
}

func (p *PacketOut) Len() (n uint16) {
	n += p.Header.Len()
	n += 8
	n += p.ActionsLen
	for _, a := range p.Actions {
		n += a.Len()
	}
	n += p.Data.Len()
	//if n < 72 { return 72 }
	return
}

func (p *PacketOut) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(p.Len()))
	b := make([]byte, 0)
	n := 0

	p.Header.Length = p.Len()
	b, err = p.Header.MarshalBinary()
	copy(data[n:], b)
	n += len(b)

	binary.BigEndian.PutUint32(data[n:], p.BufferId)
	n += 4
	binary.BigEndian.PutUint16(data[n:], p.InPort)
	n += 2
	binary.BigEndian.PutUint16(data[n:], p.ActionsLen)
	n += 2

	for _, a := range p.Actions {
		b, err = a.MarshalBinary()
		copy(data[n:], b)
		n += len(b)
	}

	b, err = p.Data.MarshalBinary()
	copy(data[n:], b)
	n += len(b)
	return
}

func (p *PacketOut) UnmarshalBinary(data []byte) error {
	err := p.Header.UnmarshalBinary(data)
	n := p.Header.Len()

	p.BufferId = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.InPort = binary.BigEndian.Uint16(data[n:])
	n += 2
	p.ActionsLen = binary.BigEndian.Uint16(data[n:])
	n += 2

	actionsUpperBound := n + p.ActionsLen
	for n < actionsUpperBound {
		a, err := DecodeAction(data[n:])
		if err != nil {
			return err
		}
		p.Actions = append(p.Actions, a)
		n += a.Len()
	}

	p.Data = openflowxx.NewBuffer(data[n:])
	return err
}
