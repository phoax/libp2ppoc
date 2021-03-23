package mymessage

type FCRMessage struct {
	MessageType       int32   `json:"message_type"`
	ProtocolVersion   int32   `json:"protocol_version"`
	ProtocolSupported []int32 `json:"protocol_supported"`
	MessageBody       []byte  `json:"message_body"`
	Signature         string  `json:"message_signature"`
}
