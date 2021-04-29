package utils

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConverter(t *testing.T) {

	Convey("Given a text, should be converted to...", t, func() {

		Convey("from hex string to text", func() {

			textHex := "30323030"

			hexStr, _ := ConvertHexToText(textHex)
			convey.Printf("hexStr %s", hexStr)

			So(hexStr, ShouldEqual, "0200")
		})

		Convey("from hex string to decimal", func() {

			hexNumber := "0249"

			decimal := ConvertHexStrToDecimal(hexNumber)
			convey.Printf("decimal %d", decimal)

			So(decimal, ShouldEqual, 585)
		})

		Convey("from decimal to hex string", func() {

			decimal := 585
			formatter := "%04x"

			hexStr := ConvertDecimalToHexStr(decimal, formatter)
			convey.Printf("hexStr %s", hexStr)

			So(hexStr, ShouldEqual, "0249")
		})

	})
}
