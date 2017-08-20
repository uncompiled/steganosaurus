package modules

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const hiddenZero = `​` // Zero-width space
const hiddenOne = `‍`  // Zero-width non-joiner

// ZeroWidthEncode converts an input stream into zero-width code points
func ZeroWidthEncode(input io.Reader, output io.Writer) {
	inputBuffer := bufio.NewScanner(bufio.NewReader(input))
	for inputBuffer.Scan() {
		currentLine := inputBuffer.Text()
		encodedText := encode(stringToBin(currentLine))
		fmt.Fprintln(output, encodedText)
	}
}

// ZeroWidthDecode converts an input stream containing zero-width into plaintext
func ZeroWidthDecode(input io.Reader, output io.Writer) {
	inputBuffer := bufio.NewScanner(bufio.NewReader(input))
	for inputBuffer.Scan() {
		currentLine := inputBuffer.Text()
		hiddenMessage := findHiddenMessage(currentLine)
		if len(hiddenMessage) > 0 {
			fmt.Fprintln(output, binToString(decode(hiddenMessage)))
		} else {
			fmt.Fprintln(output, currentLine)
		}
	}
}

// findHiddenMessage returns a new string containing all of the zero-width codepoints in the string
func findHiddenMessage(input string) (output string) {
	for _, c := range input {
		if string(c) == hiddenOne || string(c) == hiddenZero {
			output = fmt.Sprintf("%s%s", output, string(c))
		}
	}
	return
}

// Decodes hidden text by turning zero-width codepoints into ones and zeroes
func decode(input string) (output string) {
	output = input
	output = strings.Replace(output, hiddenZero, "0", -1)
	output = strings.Replace(output, hiddenOne, "1", -1)
	return output
}

// Encodes plaintext ones and zeros into zero-width codepoints
func encode(input string) (output string) {
	output = input
	output = strings.Replace(output, "0", hiddenZero, -1)
	output = strings.Replace(output, "1", hiddenOne, -1)
	return output
}

// stringToBin converts a string to a binary string representation
func stringToBin(s string) (binString string) {
	for _, c := range s {
		// Build binString by appending the binary representation of each rune in the original string
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return
}

// binToString converts a binary string into codepoints
func binToString(binString string) (s string) {
	for i := range binString {
		if i%8 == 0 {
			// read eight bytes at a time
			binaryCodepoint := binString[i : i+8]
			// we ignore binary strings that can't be parsed
			codepoint, _ := strconv.ParseInt(binaryCodepoint, 2, 32)
			s = fmt.Sprintf("%s%c", s, codepoint)
		}
	}
	return
}
