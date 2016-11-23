package openflow10

import (
	"github.com/maufl/openflow/openflowxx"
)

// Echo request/reply messages can be sent from either the
// switch or the controller, and must return an echo reply. They
// can be used to indicate the latency, bandwidth, and/or
// liveness of a controller-switch connection.

type EchoRequest struct {
	*openflowxx.Header
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{openflowxx.NewHeader(VERSION, Type_EchoRequest)}
}

// Echo request/reply messages can be sent from either the
// switch or the controller, and must return an echo reply. They
// can be used to indicate the latency, bandwidth, and/or
// liveness of a controller-switch connection.

type EchoReply struct {
	*openflowxx.Header
}

func NewEchoReply() *EchoReply {
	return &EchoReply{openflowxx.NewHeader(VERSION, Type_EchoReply)}
}
