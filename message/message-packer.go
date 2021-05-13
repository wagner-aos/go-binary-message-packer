package messagepacker

import (
	"errors"
	"fmt"

	utils "github.com/wagner-aos/go-binary-message-packer/utils"
)

const (
	messageTag = "[MessagePacker]-"
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
func (mp *MessagePack) NewMessagePack(clientName string, headerSize int, headerTypeReceived, headerTypeSent, payloadTypeReceived, payloadTypeSent string) error {
	if len(clientName) == 0 {
		return errors.New(messageTag + "client name cannot be empty")
	}
	if headerSize < 0 {
		return errors.New(messageTag + "headerSize must not have lesser than 0")
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

	mp.HeaderSize = headerSize
	mp.ClientName = clientName
	mp.HeaderTypeReceived = headerTypeReceived
	mp.HeaderTypeSent = headerTypeSent
	mp.PayloadTypeReceived = payloadTypeReceived
	mp.PayloadTypeSent = payloadTypeSent

	return nil
}

//GetDecimalHeader - it returns the decimal header content
func (mp *MessagePack) GetDecimalHeader(header []byte) (int, error) {
	decimalHeader := 0
	var err error
	//HEADER
	switch mp.HeaderTypeReceived {
	case "hex":
		decimalHeader, err = ConvertMessageHeaderHexToInt(string(header))
	case "decimal":
		decimalHeader, err = ConvertMessageHeaderASCIIToInt(string(header))
	}
	return decimalHeader, err
}

//GetDecimalPayloadSize - it returns the length of the payload
func (mp *MessagePack) GetDecimalPayloadSize(payload []byte) int {
	decimalPayloadSize := 0
	//PAYLOAD
	switch mp.PayloadTypeReceived {
	case "hex":
		decimalPayloadSize = len(payload) //TODO - BUG * 2
	case "ascii":
		decimalPayloadSize = len(payload)
	}
	return decimalPayloadSize
}

//ValidateMessageSize - it returns if message payload size matches with header content.
func (mp *MessagePack) ValidateMessageSize(header, payload []byte) (bool, error) {
	isValid := true
	decimalHeaderSize, err := mp.GetDecimalHeader(header)
	if err != nil {
		return false, err
	}
	decimalPayloadSize := mp.GetDecimalPayloadSize(payload)
	if decimalHeaderSize != decimalPayloadSize {
		return false, fmt.Errorf("message size error: HeaderSize = %d, PayloadSize = %d", decimalHeaderSize, decimalPayloadSize)
	}
	return isValid, nil
}

//ConvertMessageHeader - it converts message header according to message pack
func (mp *MessagePack) ConvertMessageHeader(header []byte) ([]byte, error) {
	headerResponse := ""
	var err error
	//HEADER
	switch mp.HeaderTypeReceived {
	case "hex":
		switch mp.HeaderTypeSent {
		case "hex":
			headerResponse = string(header[:mp.HeaderSize])
		case "decimal":
			headerResponse, err = ConvertMessageHeaderHexToDecimalString(string(header[:mp.HeaderSize]))
		}
	case "decimal":
		switch mp.HeaderTypeSent {
		case "hex":
			headerResponse, err = ConvertMessageHeaderDecimalStringToHex(string(header[:mp.HeaderSize]))
		case "decimal":
			headerResponse = string(header[:mp.HeaderSize])
		}
	}
	return []byte(headerResponse), err
}

//ConvertMessagePayload - it converts message payload according to message pack
func (mp *MessagePack) ConvertMessagePayload(payload []byte) ([]byte, error) {
	payloadResponse := ""
	var err error
	//PAYLOAD
	switch mp.PayloadTypeReceived {
	case "hex":
		switch mp.PayloadTypeSent {
		case "hex":
			payloadResponse = string(payload)
		case "ascii":
			payloadResponse, err = utils.ConvertHexToText(string(payload))
		}
	case "ascii":
		switch mp.PayloadTypeSent {
		case "hex":
			payloadResponse = utils.ConvertTextToHexString(payload)
		case "ascii":
			payloadResponse = string(payload)
		}
	}
	return []byte(payloadResponse), err
}
