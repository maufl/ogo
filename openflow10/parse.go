package openflow10

import (
	"errors"

	"github.com/maufl/openflow/openflowxx"
)

func Parse(b []byte) (message openflowxx.Message, err error) {
	switch b[1] {
	case Type_Hello:
		message = NewHello()
	case Type_Error:
		message = NewError()
	case Type_EchoRequest:
		message = NewEchoRequest()
	case Type_EchoReply:
		message = NewEchoReply()
	case Type_Vendor:
		message = NewVendor()
	case Type_FeaturesRequest:
		message = NewFeaturesRequest()
	case Type_FeaturesReply:
		message = NewFeaturesReply()
	case Type_GetConfigRequest:
		message = NewGetConfigRequest()
	case Type_GetConfigReply:
		message = NewGetConfigReply()
	case Type_SetConfig:
		message = NewSetConfig()
	case Type_PacketIn:
		message = NewPacketIn()
	case Type_FlowRemoved:
		message = NewFlowRemoved()
	case Type_PortStatus:
		message = NewPortStatus()
	case Type_PacketOut:
		message = NewPacketOut()
	case Type_FlowMod:
		message = NewFlowMod()
	case Type_PortMod:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	case Type_StatsRequest:
		message = NewStatsRequest()
	case Type_StatsReply:
		message = NewStatsReply()
	case Type_BarrierRequest:
		message = new(openflowxx.Header)
	case Type_BarrierReply:
		message = new(openflowxx.Header)
	case Type_QueueGetConfigRequest:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	case Type_QueueGetConfigReply:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	default:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	}
	if message != nil {
		err = message.UnmarshalBinary(b)
	}
	return
}
