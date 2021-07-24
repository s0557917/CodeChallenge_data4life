package main

import (
	"github.com/s0557917/CodingChallenge_data4life/FileManagement"
)

//The main method calls a function which creates the token file
//and then one that reads the generated file and saves it's contents to the database
func main() {
	FileManagement.CreateAndWriteTokens(10000000, 7)
	FileManagement.ReadAndSaveTokens()

	// The next code segment does the same as line 11 but with the difference that it saves and
	//prints the frequency map. It has been disabled in order to not clutter the terminal.

	// var frequencyMap map[string]int = FileManagement.ReadAndSaveTokens()
	// for key, value := range frequencyMap {
	// 	fmt.Println("Token: ", key, " -Value: ", value)
	// }
}
