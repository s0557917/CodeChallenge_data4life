package main

import (
	"fmt"
	"time"

	"github.com/s0557917/CodingChallenge_data4life/FileManagement"
)

func main() {
	start := time.Now()
	FileManagement.CreateAndWriteTokens(100, 7)
	FileManagement.ReadAndSaveTokens()
	fmt.Println(time.Since(start))
}
