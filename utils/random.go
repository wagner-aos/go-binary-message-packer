package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Random - generates a a number between min and max integers
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 0
	TOTAL := 0
	if len(os.Args) > 3 {
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	} else {
		fmt.Println("Usage:", os.Args[0], "MIX MAX TOTAL")
		return
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < TOTAL; i++ {
		myrand := Random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
