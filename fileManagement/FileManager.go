package fileManagement

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
const newLine = "\n"

func CreateAndWriteTokens(tokenCount int, tokenLength int) {
	var buffer bytes.Buffer

	for i := 0; i < tokenCount; i++ {
		generatedToken := make([]byte, tokenLength+1)

		for j := 0; j < len(generatedToken); j++ {
			generatedToken[j] = letters[rand.Int63()%int64(len(letters))]
		}
		generatedToken[tokenLength] = newLine[0]
		buffer.Write(generatedToken)
	}

	WriteToFile(buffer.String())
}

func WriteToFile(tokens string) {
	file, err := os.OpenFile("./tokens.txt", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error creating the File!: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	datawriter.WriteString(tokens)

	datawriter.Flush()
	file.Close()
}

func ReadTokens() {
	file, err := os.Open("./tokens.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
