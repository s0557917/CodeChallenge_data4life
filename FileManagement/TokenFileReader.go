package FileManagement

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/s0557917/CodingChallenge_data4life/DatabaseManagement"
)

//This function takes care of reading the generated token file line by line
//and passig them to the database manager
func ReadAndSaveTokens() map[string]int {
	file, err := os.Open("./tokens.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//A new scanner gets created
	scanner := bufio.NewScanner(file)

	//A connection to the database is created and the relevant setup is done
	DatabaseManagement.ConnectToDatabase()

	fmt.Println(" - Writing to DB")

	frequencyMap := make(map[string]int)

	//The scanner can then be iterated to read each line in the file.
	//Each line/token gets passed to a function which adds it to the transaction
	for scanner.Scan() {
		DatabaseManagement.WriteTokensToDatabase(scanner.Text())
		frequencyMap[scanner.Text()] += 1
	}

	for key, value := range frequencyMap {
		if value <= 1 {
			delete(frequencyMap, key)
		}
	}

	//When all tokens have been read, the transaction is commited and the connection is closed
	DatabaseManagement.CloseDatabaseConnection()
	fmt.Println(" - Finished writing to DB")

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file! %s", err)
	}

	//At the end the connection to the file gets closed
	file.Close()
	return frequencyMap
}
