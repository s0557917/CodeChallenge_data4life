package main

import (
	"fmt"
	"time"

	"github.com/s0557917/CodingChallenge_data4life/fileManagement"
)

func main() {
	start := time.Now()
	fileManagement.CreateAndWriteTokens(10000000, 7)
	fileManagement.ReadTokens()
	fmt.Println(time.Since(start))
}
