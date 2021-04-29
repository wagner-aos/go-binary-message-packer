package utils

import (
	"encoding/json"

	"github.com/kataras/golog"
)

//JSONPrintObjectIndented - it converts a object to JSON and prints indented.
func JSONPrintObjectIndented(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "  ")
	golog.Info("\n" + string(result))
}
