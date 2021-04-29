package messagepacker

import (
	"errors"
	"fmt"

	"go-binary-message-packer/utils"
)

const (
	messageTag  = "[MessagePacker]-"
	HEADER_SIZE = 4
)

/* MessagePack - it defines the type of encoding for message (header + payload)
Parameters:
	ClientName          string  (cannot be empty)
	HeaderSize          int     fixed 4
	HeaderTypeReceived  string  (hex | decimal)
	HeaderTypeSent      string  (hex | decimal)
	PayloadTypeReceived string  (hex | ascii)
	PayloadTypeSent     string  (hex | ascii)

Use the method NewMessagePack in order to create a new object.

*/
type MessagePack struct {
	ClientName          string
	HeaderSize          int
	HeaderTypeReceived  string
	HeaderTypeSent      string
	PayloadTypeReceived string
	PayloadTypeSent     string
}

//NewMessagePack - Creates an object to define inbound / outbound message protocol.
func (mp *MessagePack) NewMessagePack(clientName, headerTypeReceived, headerTypeSent, payloadTypeReceived, payloadTypeSent string) error {
	if len(clientName) == 0 {
		return errors.New(messageTag + "client name cannot be empty")
	}
	if headerTypeReceived != "hex" && headerTypeReceived != "decimal" {
		return errors.New(messageTag + "headerTypeReceived must contain 'hex' or 'decimal' type")
	}
	if headerTypeSent != "hex" && headerTypeSent != "decimal" {
		return errors.New(messageTag + "headerTypeSent must contain 'hex' or 'decimal' type")
	}
	if payloadTypeReceived != "hex" && payloadTypeReceived != "ascii" {
		return errors.New(messageTag + "payloadTypeReceived must contain 'hex' or 'ascii' type")
	}
	if payloadTypeSent != "hex" && payloadTypeSent != "ascii" {
		return errors.New(messageTag + "payloadTypeSent must contain 'hex' or 'ascii' type")
	}

	mp.HeaderSize = HEADER_SIZE
	mp.ClientName = clientName
	mp.HeaderTypeReceived = headerTypeReceived
	mp.HeaderTypeSent = headerTypeSent
	mp.PayloadTypeReceived = payloadTypeReceived
	mp.PayloadTypeSent = payloadTypeSent

	return nil
}

//GetDecimalHeader - it returns the decimal header content
func (mp *MessagePack) GetDecimalHeader(header []byte) int {
	decimalHeader := 0
	//HEADER
	switch mp.HeaderTypeReceived {
	case "hex":
		decimalHeader = ConvertMessageHeaderHexToInt(string(header))
	case "decimal":
		decimalHeader = ConvertMessageHeaderASCIIToInt(string(header))
	}
	return decimalHeader
}

//GetDecimalPayloadSize - it returns the length of the payload
func (mp *MessagePack) GetDecimalPayloadSize(payload []byte) int {
	decimalPayloadSize := 0
	//PAYLOAD
	switch mp.PayloadTypeReceived {
	case "hex":
		decimalPayloadSize = len(payload) * 2
	case "ascii":
		decimalPayloadSize = len(payload)
	}
	return decimalPayloadSize
}

//ValidateMessageSize - it returns if message payload size matches with header content.
func (mp *MessagePack) ValidateMessageSize(header, payload []byte) (bool, error) {
	isValid := true
	decimalHeaderSize := mp.GetDecimalHeader(header)
	decimalPayloadSize := mp.GetDecimalPayloadSize(payload)
	if decimalHeaderSize != decimalPayloadSize {
		return false, fmt.Errorf("message size error: HeaderSize = %d, PayloadSize = %d", decimalHeaderSize, decimalPayloadSize)
	}
	return isValid, nil
}

//ConvertMessageHeader - it converts message header according to message pack
func (mp *MessagePack) ConvertMessageHeader(header []byte) []byte {
	headerResponse := ""
	//HEADER
	switch mp.HeaderTypeReceived {
	case "hex":
		switch mp.HeaderTypeSent {
		case "hex":
			headerResponse = string(header)
		case "decimal":
			headerResponse = ConvertMessageHeaderHexToDecimalString(string(header))
		}
	case "decimal":
		switch mp.HeaderTypeSent {
		case "hex":
			headerResponse = ConvertMessageHeaderDecimalStringToHex(string(header))
		case "decimal":
			headerResponse = string(header)
		}
	}
	return []byte(headerResponse)
}

//ConvertMessagePayload - it converts message payload according to message pack
func (mp *MessagePack) ConvertMessagePayload(payload []byte) []byte {
	payloadResponse := ""
	//PAYLOAD
	switch mp.PayloadTypeReceived {
	case "hex":
		switch mp.PayloadTypeSent {
		case "hex":
			payloadResponse = string(payload)
		case "ascii":
			payloadResponse, _ = utils.ConvertHexToText(string(payload))
		}
	case "ascii":
		switch mp.PayloadTypeSent {
		case "hex":
			payloadResponse = utils.ConvertTextToHexString(payload)
		case "ascii":
			payloadResponse = string(payload)
		}
	}
	return []byte(payloadResponse)
}
