package messagepacker

import (
	"go-binary-message-packer/utils"
	"strconv"

	"github.com/kataras/golog"
)

/*
This module represent message header that represents the hex size of the iso message that travels on the tcp socket
Its length is 4 bytes and all the messages starts with it.
Exemplo: HEX String: "0249" means the message content size in decimal string is "585"
*/

//ConvertMessageHeaderHexToDecimalString - It converts message HEX header to STR
func ConvertMessageHeaderHexToDecimalString(headerHexStr string) string {
	golog.Infof("HeaderHEX: %s", headerHexStr)
	headerInt := ConvertMessageHeaderHexToInt(headerHexStr)
	headerStr := formatHeaderDecimalString(headerInt)
	golog.Infof("HeaderStr: %s", headerStr)
	return headerStr
}

//ConvertMessageHeaderDecimalStringToHex - It converts message HEX header to STR
func ConvertMessageHeaderDecimalStringToHex(headerDecimalStr string) string {
	golog.Infof("HeaderDecimalStr: %s", headerDecimalStr)
	headerInt := utils.ConvertStringToInt(headerDecimalStr)
	hexHeader := utils.ConvertDecimalToHexStr(headerInt, HeaderHexFormatter)
	golog.Infof("HeaderHex: %s", hexHeader)
	return hexHeader
}

//ConvertMessageHeaderHexToInt - It converts message HEX header to INT
func ConvertMessageHeaderHexToInt(headerHexStr string) int {
	golog.Infof("HeaderHEX: %s", headerHexStr)
	value, err := strconv.ParseInt(headerHexStr, 16, 64)
	if err != nil {
		golog.Error("Error When parsing header.")
		return 0
	}
	golog.Infof("HeaderInt: %d", value)
	return int(value)
}

//ConvertMessageHeaderASCIIToInt - It converts message ASCII header to INT
func ConvertMessageHeaderASCIIToInt(headerASCIIStr string) int {
	golog.Infof("HeaderASCII: %s", headerASCIIStr)
	value, err := utils.ConvertTextToInt(headerASCIIStr)
	if err != nil {
		golog.Error("Error When parsing header.")
		return 0
	}
	golog.Infof("HeaderInt: %d", value)
	return int(value)
}

func formatHeaderDecimalString(headerInt int) string {
	headerStr := strconv.Itoa(headerInt)
	if headerInt <= 999 {
		return "0" + headerStr
	}
	return headerStr
}

// //ConvertMtiHexToStr - it converts MTI from HEX to String
// func ConvertMtiHexToStr(mtiHex string) string {
// 	golog.Infof("MtiHEX: %s", mtiHex)
// 	mtiStr := utils.ConvertHexToText(mtiHex)
// 	golog.Infof("MtiStr: %s", mtiStr)
// 	return mtiStr
// }

//ConvertMessageHexToStr - it converts message HEX to String and also return its length
func ConvertMessageHexToStr(messageHexContent string) (string, int) {
	//HEX
	golog.Infof("messageHEX: %s", messageHexContent)
	//Text
	messageStr, _ := utils.ConvertHexToText(messageHexContent)
	golog.Infof("messageStr: %s", messageStr)
	return messageStr, len(messageStr)
}

// //ConvertMessageStrToHex - it converts message String to Hex message and also return its length
// /*Returns:
// messageHex = hexHeader + hexMti + hexContent
// messageSize = len(hexMti+hexContent) / 2
// */
// func ConvertMessageStrToHex(message string) (string, int) {
// 	mtiInt := utils.ConvertTextToInt(message[0:MtiSize])
// 	mtiStr := generateMtiStrFromInt(mtiInt)
// 	content := message[MtiSize:]
// 	hexHeader := utils.ConvertDecimalToHexStr(len(mtiStr)+len(content), "%04x")
// 	hexMti := utils.ConvertTextToHexString([]byte(mtiStr))
// 	hexContent := utils.ConvertTextToHexString([]byte(content))
// 	messageHex := hexHeader + hexMti + hexContent
// 	return messageHex, len(hexMti+hexContent) / 2
// }

// //ConvertMessageStrToHexComponents - it returns hex components from message string
// //Returns: hexHeader, hexMti, hexContent, hexContentSize
// func ConvertMessageStrToHexComponents(message string) (string, string, string, int) {
// 	mtiInt := utils.ConvertTextToInt(message[0:MtiSize])
// 	mtiStr := generateMtiStrFromInt(mtiInt)
// 	content := message[MtiSize:]
// 	hexHeader := utils.ConvertDecimalToHexStr(len(mtiStr+content), "%04x")
// 	hexMti := utils.ConvertTextToHexString([]byte(mtiStr))
// 	hexContent := utils.ConvertTextToHexString([]byte(content))
// 	hexContentSize := len(hexContent)
// 	return hexHeader, hexMti, hexContent, hexContentSize
// }
