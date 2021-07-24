package FileManagement

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
const newLine = '\n'

//The method creates a tokenCount amount of tokens with the length tokenLength
func CreateAndWriteTokens(tokenCount int, tokenLength int) {

	//Using a byte buffer is more performant as it allows for only one write access to the file
	var buffer bytes.Buffer

	//The tokens are randomly created and added to a byte slice by accessing and getting random
	//letters from an alphabet constant. Then a new line is added at the end of the generated token,
	//which is why the tokenLength gets increased by one position. This ensures proper formatting
	//for when the file is read back in. The use of constants is deliberate, as it takes advantage
	//of the automatic optimization most compilers do when passed constants instead of vars.
	for i := 0; i < tokenCount; i++ {
		generatedToken := make([]byte, tokenLength+1)

		for j := 0; j < len(generatedToken); j++ {
			generatedToken[j] = letters[rand.Int63()%int64(len(letters))]
		}
		generatedToken[tokenLength] = newLine
		buffer.Write(generatedToken)
	}

	//The byte buffer is converted into a string, which will then be written into the file
	WriteToFile(buffer.String())
}

//This function writes the passed string into a new tokens.txt file
func WriteToFile(tokens string) {
	//If the doesn't exist it gets created (os.O_CREATE).
	//If the file exists, it gets opened in write mode (os.O_WRONLY) and the new tokens
	//overwrite any old content (os.O_TRUNC)
	file, err := os.OpenFile("./tokens.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error creating the File!: %s", err)
	}

	//A new writer is created, and because all tokens were written into a single string
	//only one write call needs to be made to the file, signficantly reducing the access time.
	datawriter := bufio.NewWriter(file)
	datawriter.WriteString(tokens)

	//The writer is then flushed and the file closed to ensure proper access when reading later
	datawriter.Flush()
	file.Close()
}
