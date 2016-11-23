package openflow10

import (
	"encoding/binary"
	"errors"
	"github.com/maufl/openflow/openflowxx"
)

// The OFPT_HELLO message consists of an OpenFlow header plus a set of variable
// size hello elements. The version field part of the header field (see 7.1)
// must be set to the highest OpenFlow switch protocol version supported by the
// sender (see 6.3.1).  The elements field is a set of hello elements,
// containing optional data to inform the initial handshake of the connection.
// Implementations must ignore (skip) all elements of a Hello message that they
// do not support.
// The version field part of the header field (see 7.1) must be set to the
// highest OpenFlow switch protocol version supported by the sender (see 6.3.1).
// The elements field is a set of hello elements, containing optional data to
// inform the initial handshake of the connection. Implementations must ignore
// (skip) all elements of a Hello message that they do not support.
type Hello struct {
	*openflowxx.Header
	Elements []HelloElem
}

func NewHello() (h *Hello) {
	return &Hello{
		Header:   openflowxx.NewHeader(1, Type_Hello),
		Elements: []HelloElem{NewHelloElemVersionBitmap()},
	}
}

func (h *Hello) Len() (n uint16) {
	n = h.Header.Len()
	for _, e := range h.Elements {
		n += e.Len()
	}
	return
}

func (h *Hello) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(h.Len()))
	bytes := make([]byte, 0)
	next := 0

	h.Header.Length = h.Len()
	bytes, err = h.Header.MarshalBinary()
	copy(data[next:], bytes)
	next += len(bytes)

	for _, e := range h.Elements {
		bytes, err = e.MarshalBinary()
		copy(data[next:], bytes)
		next += len(bytes)
	}
	return
}

func (h *Hello) UnmarshalBinary(data []byte) error {
	next := 0
	err := h.Header.UnmarshalBinary(data[next:])
	next += int(h.Header.Len())

	h.Elements = make([]HelloElem, 0)
	for next < len(data) {
		e := NewHelloElemHeader()
		e.UnmarshalBinary(data[next:])

		switch e.Type {
		case HelloElemType_VersionBitmap:
			v := NewHelloElemVersionBitmap()
			err = v.UnmarshalBinary(data[next:])
			next += int(v.Len())
			h.Elements = append(h.Elements, v)
		}
	}
	return err
}

const (
	reserved = iota
	HelloElemType_VersionBitmap
)

type HelloElem interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
	Len() uint16
	Header() *HelloElemHeader
}

type HelloElemHeader struct {
	Type   uint16
	Length uint16
}

func NewHelloElemHeader() *HelloElemHeader {
	h := new(HelloElemHeader)
	h.Type = HelloElemType_VersionBitmap
	h.Length = 4
	return h
}

func (h *HelloElemHeader) Header() *HelloElemHeader {
	return h
}

func (h *HelloElemHeader) Len() (n uint16) {
	return 4
}

func (h *HelloElemHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint16(data[:2], h.Type)
	binary.BigEndian.PutUint16(data[2:4], h.Length)
	return
}

func (h *HelloElemHeader) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return errors.New("The []byte is too short to unmarshal a full HelloElemHeader.")
	}
	h.Type = binary.BigEndian.Uint16(data[:2])
	h.Length = binary.BigEndian.Uint16(data[2:4])
	return nil
}

type HelloElemVersionBitmap struct {
	HelloElemHeader
	Bitmaps []uint32
}

func NewHelloElemVersionBitmap() *HelloElemVersionBitmap {
	h := new(HelloElemVersionBitmap)
	h.HelloElemHeader = *NewHelloElemHeader()
	h.Bitmaps = make([]uint32, 0)
	// 1001
	// h.Bitmaps = append(h.Bitmaps, uint32(8) | uint32(1))
	h.Bitmaps = append(h.Bitmaps, uint32(1))
	h.Length = h.Length + uint16(len(h.Bitmaps)*4)
	return h
}

func (h *HelloElemVersionBitmap) Header() *HelloElemHeader {
	return &h.HelloElemHeader
}

func (h *HelloElemVersionBitmap) Len() (n uint16) {
	n = h.HelloElemHeader.Len()
	n += uint16(len(h.Bitmaps) * 4)
	return
}

func (h *HelloElemVersionBitmap) MarshalBinary() (data []byte, err error) {
	data = make([]byte, int(h.Len()))
	bytes := make([]byte, 0)
	next := 0

	bytes, err = h.HelloElemHeader.MarshalBinary()
	copy(data[next:], bytes)
	next += len(bytes)

	for _, m := range h.Bitmaps {
		binary.BigEndian.PutUint32(data[next:], m)
		next += 4
	}
	return
}

func (h *HelloElemVersionBitmap) UnmarshalBinary(data []byte) error {
	length := len(data)
	read := 0
	if err := h.HelloElemHeader.UnmarshalBinary(data[:4]); err != nil {
		return err
	}
	read += int(h.HelloElemHeader.Len())

	h.Bitmaps = make([]uint32, 0)
	for read < length {
		h.Bitmaps = append(h.Bitmaps, binary.BigEndian.Uint32(data[read:read+4]))
		read += 4
	}
	return nil
}
