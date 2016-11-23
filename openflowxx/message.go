package openflowxx

type Message interface {
	//encoding.BinaryMarshaler
	//encoding.BinaryUnmarshaler
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
	Head() *Header
	Len() uint16
}
