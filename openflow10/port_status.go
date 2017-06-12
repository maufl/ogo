package openflow10

import (
	"fmt"

	"github.com/maufl/openflow/openflowxx"
)

// ofp_port_status
type PortStatus struct {
	*openflowxx.Header
	Reason uint8
	pad    []uint8 // Size 7
	Desc   PhyPort
}

func NewPortStatus() *PortStatus {
	return &PortStatus{
		Header: openflowxx.NewHeader(VERSION, Type_PortStatus),
		pad:    make([]byte, 7),
		Desc:   *NewPhyPort(),
	}
}

func (ps *PortStatus) String() string {
	return fmt.Sprintf("PortStatus{ Reason: %d, Desc: %s }", ps.Reason, ps.Desc)
}

func (p *PortStatus) Len() (n uint16) {
	n = p.Header.Len()
	n += 8
	n += p.Desc.Len()
	return
}

func (s *PortStatus) MarshalBinary() (data []byte, err error) {
	s.Header.Length = s.Len()
	data, err = s.Header.MarshalBinary()

	b := make([]byte, 8)
	n := 0
	b[0] = s.Reason
	n += 1
	copy(b[n:], s.pad)
	data = append(data, b...)

	b, err = s.Desc.MarshalBinary()
	data = append(data, b...)
	return
}

func (s *PortStatus) UnmarshalBinary(data []byte) error {
	err := s.Header.UnmarshalBinary(data)
	n := int(s.Header.Len())

	s.Reason = data[n]
	n += 1
	copy(s.pad, data[n:])
	n += len(s.pad)

	err = s.Desc.UnmarshalBinary(data[n:])
	return err
}

// ofp_port_reason 1.0
const (
	PR_ADD = iota
	PR_DELETE
	PR_MODIFY
)
