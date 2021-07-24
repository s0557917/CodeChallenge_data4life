package FileManagement

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/s0557917/CodingChallenge_data4life/DatabaseManagement"
)

func ReadAndSaveTokens() {
	file, err := os.Open("./tokens.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	DatabaseManagement.ConnectToDatabase()

	fmt.Println(" - Writing to DB")

	for scanner.Scan() {
		DatabaseManagement.WriteTokensToDatabase(scanner.Text())
	}

	DatabaseManagement.CloseDatabaseConnection()
	fmt.Println(" - Finished writing to DB")

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file! %s", err)
	}

	file.Close()
}
