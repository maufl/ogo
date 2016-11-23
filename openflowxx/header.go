// Package ofpxx defines OpenFlow message types that are version independent.
package openflowxx

import (
	"encoding/binary"
	"errors"
	"math/rand"
)

// The version specifies the OpenFlow protocol version being
// used. During the current draft phase of the OpenFlow
// Protocol, the most significant bit will be set to indicate an
// experimental version and the lower bits will indicate a
// revision number. The current version is 0x01. The final
// version for a Type 0 switch will be 0x00. The length field
// indicates the total length of the message, so no additional
// framing is used to distinguish one frame from the next.
type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	Xid     uint32
}

func NewHeader(version, typ int) *Header {
	return &Header{
		Version: uint8(version),
		Type:    uint8(typ),
		Length:  8,
		Xid:     rand.Uint32(),
	}
}

func (h *Header) Head() *Header {
	return h
}

func (h *Header) Len() (n uint16) {
	return 8
}

func (h *Header) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8)
	data[0] = h.Version
	data[1] = h.Type
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	binary.BigEndian.PutUint32(data[4:8], h.Xid)
	return
}

func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to unmarshel a full HelloElemHeader.")
	}
	h.Version = data[0]
	h.Type = data[1]
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.Xid = binary.BigEndian.Uint32(data[4:8])
	return nil
}
