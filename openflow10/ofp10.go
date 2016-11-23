// OpenFlow Wire Protocol 0x01
// Package openflow10 provides OpenFlow 1.0 structs along with Read
// and Write methods for each.
//
// Struct documentation is taken from the OpenFlow Switch
// Specification Version 1.0.0.
// https://www.opennetworking.org/images/stories/downloads/sdn-resources/onf-specifications/openflow/openflow-spec-v1.0.0.pdf
package openflow10

const (
	VERSION = 1
)

// ofp_type 1.0
const (
	/* Immutable messages. */
	Type_Hello = iota
	Type_Error
	Type_EchoRequest
	Type_EchoReply
	Type_Vendor

	/* Switch configuration messages. */
	Type_FeaturesRequest
	Type_FeaturesReply
	Type_GetConfigRequest
	Type_GetConfigReply
	Type_SetConfig

	/* Asynchronous messages. */
	Type_PacketIn
	Type_FlowRemoved
	Type_PortStatus

	/* Controller command messages. */
	Type_PacketOut
	Type_FlowMod
	Type_PortMod

	/* Statistics messages. */
	Type_StatsRequest
	Type_StatsReply

	/* Barrier messages. */
	Type_BarrierRequest
	Type_BarrierReply

	/* Queue Configuration messages. */
	Type_QueueGetConfigRequest
	Type_QueueGetConfigReply
)
