package utils

import (
	"log"
	"os/exec"
)

//GenerateUUID -
func GenerateUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
