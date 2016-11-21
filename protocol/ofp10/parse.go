package ofp10

import (
	"errors"
	
	"github.com/maufl/openflow/protocol/ofpxx"
	"github.com/maufl/openflow/protocol/util"
)

func Parse(b []byte) (message util.Message, err error) {
	switch b[1] {
	case Type_Hello:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_Error:
		message = new(ErrorMsg)
		err = message.UnmarshalBinary(b)
	case Type_EchoRequest:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_EchoReply:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_Vendor:
		message = new(VendorHeader)
		err = message.UnmarshalBinary(b)
	case Type_FeaturesRequest:
		message = NewFeaturesRequest()
		err = message.UnmarshalBinary(b)
	case Type_FeaturesReply:
		message = NewFeaturesReply()
		err = message.UnmarshalBinary(b)
	case Type_GetConfigRequest:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_GetConfigReply:
		message = new(SwitchConfig)
		err = message.UnmarshalBinary(b)
	case Type_SetConfig:
		message = NewSetConfig()
		err = message.UnmarshalBinary(b)
	case Type_PacketIn:
		message = new(PacketIn)
		err = message.UnmarshalBinary(b)
	case Type_FlowRemoved:
		message = NewFlowRemoved()
		err = message.UnmarshalBinary(b)
	case Type_PortStatus:
		message = new(PortStatus)
		err = message.UnmarshalBinary(b)
	case Type_PacketOut:
		message = new(PacketOut)
		err = message.UnmarshalBinary(b)
	case Type_FlowMod:
		message = NewFlowMod()
		err = message.UnmarshalBinary(b)
	case Type_PortMod:
		break
	case Type_StatsRequest:
		message = new(StatsRequest)
		err = message.UnmarshalBinary(b)
	case Type_StatsReply:
		message = new(StatsReply)
		err = message.UnmarshalBinary(b)
	case Type_BarrierRequest:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_BarrierReply:
		message = new(ofpxx.Header)
		err = message.UnmarshalBinary(b)
	case Type_QueueGetConfigRequest:
		break
	case Type_QueueGetConfigReply:
		break
	default:
		err = errors.New("An unknown v1.0 packet type was received. Parse function will discard data.")
	}
	return
}
