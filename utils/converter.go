package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

//ConvertHexToText -
func ConvertHexToText(HexStr string) (string, error) {
	//fmt.Printf("\nHEX: %s", HexStr)
	bs, err := hex.DecodeString(HexStr)
	if err != nil {
		return string(bs), err
	}
	return string(bs), nil
}

//ConvertTextToInt -
func ConvertTextToInt(text string) (int, error) {
	integer, err := strconv.Atoi(text)
	if err != nil {
		return integer, err
	}
	return integer, nil
}

//ConvertIntToString -
func ConvertIntToString(number int) string {
	return strconv.Itoa(number)
}

//ConvertStringToInt -
func ConvertStringToInt(number string) (int, error) {
	result, err := strconv.Atoi(number)
	if err != nil {
		return result, err
	}
	return result, nil
}

//ConvertTextToHexString -
func ConvertTextToHexString(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

//ConvertHexStrToDecimal -
func ConvertHexStrToDecimal(hexNumber string) (int, error) {
	output, err := strconv.ParseUint(hexaNumberToInteger(hexNumber), 16, 64)
	if err != nil {
		return int(output), err
	}
	return int(output), nil
}

func hexaNumberToInteger(hexString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

//ConvertDecimalToHexStr - you can use formater, default = "%02x"
func ConvertDecimalToHexStr(numberDecimal int, formatter string) string {
	output := ""
	defaultFormatter := "%02x"
	if len(formatter) == 0 {
		output = fmt.Sprintf(defaultFormatter, numberDecimal)
	} else {
		output = fmt.Sprintf(formatter, numberDecimal)
	}
	return output
}

func ConvertBinaryToDecimal(binaryString string) (int, error) {
	output, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		return int(output), err
	}
	return int(output), nil
}

func ConvertDecimalToBinary(number int) string {
	return strconv.FormatInt(int64(number), 2)
}

func ConvertIntToBcd(value int) int {
	return (((value / 10) % 10) << 4) | (value % 10)
}

func ConvertBcdToInt(value int) int {
	return (int)((value>>4)*10 + (value & 0x0F))
}
