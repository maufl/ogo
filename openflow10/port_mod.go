package openflow10

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/maufl/openflow/openflowxx"
)

// ofp_port_mod 1.0
type PortMod struct {
	openflowxx.Header
	PortNo uint16
	HWAddr []byte

	Config    uint32
	Mask      uint32
	Advertise uint32
	pad       []byte // Size 4
}

func NewPortMod() *PortMod {
	return &PortMod{
		Header: *openflowxx.NewHeader(VERSION, Type_PortMod),
		HWAddr: make([]byte, ETH_ALEN),
		pad:    make([]byte, 4),
	}
}

func (pm *PortMod) String() string {
	return fmt.Sprintf("PortMod{ %s, PortNo: %d, HWAddr: %x, Config: %b, Mask: %b, Advertise: %b }", pm.Header, pm.HWAddr, pm.Config, pm.Mask, pm.Advertise)
}

func (p *PortMod) Len() (n uint16) {
	return p.Header.Len() + 2 + ETH_ALEN + 16
}

func (p *PortMod) MarshalBinary() (data []byte, err error) {
	p.Header.Length = p.Len()
	data, err = p.Header.MarshalBinary()

	b := make([]byte, 24)
	n := 0
	binary.BigEndian.PutUint16(b[n:], p.PortNo)
	n += 2
	copy(b[n:], p.HWAddr)
	n += ETH_ALEN
	binary.BigEndian.PutUint32(b[n:], p.Config)
	n += 4
	binary.BigEndian.PutUint32(b[n:], p.Mask)
	n += 4
	binary.BigEndian.PutUint32(b[n:], p.Advertise)
	n += 4
	copy(b[n:], p.pad)
	n += 4
	data = append(data, b...)
	return
}

func (p *PortMod) UnmarshalBinary(data []byte) error {
	if len(data) < int(p.Len()) {
		return errors.New("Insufficent data to unmarshal port mod message")
	}
	err := p.Header.UnmarshalBinary(data)
	n := int(p.Header.Len())

	p.PortNo = binary.BigEndian.Uint16(data[n:])
	n += 2
	copy(p.HWAddr, data[n:])
	n += ETH_ALEN
	p.Config = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Mask = binary.BigEndian.Uint32(data[n:])
	n += 4
	p.Advertise = binary.BigEndian.Uint32(data[n:])
	n += 4
	copy(p.pad, data[n:])
	n += len(p.pad)
	return err
}
