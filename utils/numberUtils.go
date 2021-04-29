package utils

import (
	"math"
	"strconv"

	"github.com/kataras/golog"
)

//StringToFloat - it convert numeric string to float value
func StringToFloat(value string) (float64, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		golog.Errorf("Error when converting string to integer.")
	}
	floatValue := float64(intValue) / 100
	return floatValue, err
}

//StringToFloatString - it convert numeric string to string with float value.
func StringToFloatString(value string) (string, error) {
	floatValue, err := StringToFloat(value)
	// to convert a float number to a string
	floatString := strconv.FormatFloat(floatValue, 'f', 2, 64)
	return floatString, err
}

//Round - it rounds number
func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

//RoundUp -
func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

//RoundDown -
func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

// func main() {
// 	value := "000072346"
// 	floatValue, _ := StringToDecimal(value)
// 	fmt.Printf("Float Value: %0.2f \n", floatValue)
// }
