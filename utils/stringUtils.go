package utils

import (
	"strings"
	"unicode"

	"github.com/kataras/golog"
)

const EMPTY = ""

var (
	SentinelChars = []string{
		";",
		"B",
		"D",
		"^",
		"=",
		"1011",
		"1101",
	}
)

// StringBreakingLine - It splits a string adding '\n'
func StringBreakingLine(sizeChunk int, field string) string {
	if len(field) > sizeChunk {
		chunkParts := len(field) / sizeChunk
		if len(field)%sizeChunk > 0 {
			chunkParts++
		}
		//log.Tracer("PARTS: %d", chunkParts)

		fieldSlice := []byte(field)
		stringArray := []string{}
		stringResult := string("")

		initStr, endStr := 0, 0
		for i := 0; i < chunkParts; i++ {

			if (endStr + sizeChunk) <= len(field) {
				endStr = endStr + sizeChunk
				//log.Tracer("\nentrou aqui: %d - %d  ", initStr, endStr)
			} else {
				endStr = len(field) - 1
				//log.Tracer("\nentrou max: %d - %d  ", initStr, endStr)
			}

			stringArray = append(stringArray, string(fieldSlice[initStr:endStr]))

			print(string(fieldSlice[initStr:endStr]))

			initStr = endStr
			//log.Tracer("\nIndice: %d", i)
		}

		for _, s := range stringArray {
			stringResult = stringResult + s + "\n"
		}

		return stringResult
	}
	return field
}

//GetSubstring - it returns substring with runes
func GetSubstring(value string, initalPosition int, finalPosition int) string {
	valueSize := len(value)
	if finalPosition < 0 {
		golog.Info("InitialPosition cannot be less than 0")
	}
	if finalPosition > valueSize {
		golog.Infof("FinalPosition cannot be greater than %d", valueSize)
	}
	// Part A: take substring of first word with runes.
	// ... This handles any kind of rune in the string.
	runes := []rune(value)
	// ... Convert back into a string from rune slice.
	return string(runes[initalPosition:finalPosition])
}

//Contains -
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

//Concat - it concats a list of strings
//Example: utils.Concat("string1","string2","string3")
func Concat(strlist ...string) string {
	result := ""
	for _, item := range strlist {
		result = result + item
	}
	return result
}

//GetSubstringIndex - it returns a substring from a given string, and its position.
func GetSubstringIndex(stringValue string, item string, returnSize int) (string, int, bool) {
	i := strings.Index(stringValue, item)
	if i >= 0 {
		if returnSize > len(item) {
			return stringValue[i : i+returnSize], i, true
		}
		return item, i, true
	}
	return string(""), i, false
}

// IsSentinels validate characters is sentinels in specification
func IsSentinels(attribute string) bool {
	if len(attribute) == 0 {
		return false
	}
	attr := strings.TrimFunc(attribute, unicode.IsSpace)
	for _, indicate := range SentinelChars {
		if strings.Contains(attr, indicate) {
			if indicate == "B" && len(attr) > 1 {
				continue
			}
			return true
		}
	}
	return false
}

func IsSpace(value string) bool {
	for _, v := range value {
		if unicode.IsSpace(v) {
			return true
		}
	}
	return false
}

func SubString(s string, i int) (string, int, int) {
	s = GetSubstring(s, 0, i)
	if IsSentinels(s) {
		for _, indicate := range SentinelChars {
			idx := strings.Index(s, indicate)
			if idx >= 0 {
				s = GetSubstring(s, 0, idx)

				if idx == 0 {
					idx++
				}
				i = idx
				break
			}
		}
	}
	return s, i, len(s)
}

func LeftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}
