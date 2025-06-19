package practice

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessASCII(input string, style string) (string, int) {
	maxInput := 1000

	if len(input) == 0 {
		return "", 200
	}

	if input == "\\n" {
		return "\n", 200
	}

	// Character Limit
	if len(input) > maxInput {
		fmt.Printf("Error: input too long (max %d characters)\n", maxInput)
		return "" , 500
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")
	// Ascii Character Restriction
	for _, letter := range input {
		if !((letter >= 32 && letter < 127) || letter == 10) { // 32 to 126 are printable ASCII characters
			fmt.Println("Error: Non Ascii Characters were inputted")
			return "", 400
		}
	}
	// Handle different ASCII art styles
	var filepath string
	switch style {
	case "standard":
		filepath = "asciiart/standard.txt"
	case "shadow":
		filepath = "asciiart/shadow.txt"
	case "thinkertoy":
		filepath = "asciiart/thinkertoy.txt"
	default:
		// filepath = "asciiart/standard.txt"
		return "", 400 // Return error if style is not recognized
	}
	//opening standard file
	inputFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", 500
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	str := handleNewLine(input)

	var final string
	//im looping through the word
	for i := 0; i < len(str); i++ {

		if len(str[i]) < 1 { // if the string is empty, it only prints 1 \n
			final += "\n"
			continue
		}

		char := []rune(str[i])

		var letters [][]string //an array of arrays to hold the letters. Each array contains 8 lines of each letter passed in through the argument.

		for _, x := range char { //we need to find the start and end line of the letter in the file.
			v := int(x)
			start := (v-32)*9 + 1 // changed 2 to 1 so it does not skip on the first line
			end := start + 7
			letters = append(letters, lines[start:end+1]) //end + 1 to include the last line
		}

		// Build the result by printing each row of all characters
		var result string
		for i := 0; i < 8; i++ { //for all 8 lines of each letter
			for _, letter := range letters { //looping through all the lines of each letter in the Letters array
				result += letter[i] + " " // append to result the ith line of each letter in the entire string
			}
			result += "\n" //when you're done appending all letters in the string, move to the next line. i++ and second line of all the letters will append.
		}
		final += result
	}
	fmt.Println(final)
	return final, 200 //return the final result string to the main function
}

func handleNewLine(str string) []string {
	return strings.Split(str, "\n") // split on actual newlines
}
