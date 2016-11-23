package openflow10

import (
	"encoding/binary"
	"errors"
	"github.com/maufl/openflow/openflowxx"
)

// ofp_vendor_header 1.0
type Vendor struct {
	*openflowxx.Header
	Vendor uint32
}

func NewVendor() *Vendor {
	return &Vendor{
		Header: openflowxx.NewHeader(VERSION, Type_Vendor),
	}
}

func (v *Vendor) Len() (n uint16) {
	return v.Header.Len() + 4
}

func (v *Vendor) MarshalBinary() (data []byte, err error) {
	data, err = v.Header.MarshalBinary()

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(data[:4], v.Vendor)

	data = append(data, b...)
	return
}

func (v *Vendor) UnmarshalBinary(data []byte) error {
	if len(data) < int(v.Len()) {
		return errors.New("The []byte the wrong size to unmarshal an " +
			"Vendor message.")
	}
	v.Header.UnmarshalBinary(data)
	n := int(v.Header.Len())
	v.Vendor = binary.BigEndian.Uint32(data[n:])
	return nil
}
