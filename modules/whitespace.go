package modules

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// WhitespaceMerge accepts two filenames, opens the files, and merges them to the supplied output writer.
func WhitespaceMerge(filename1, filename2 string, output io.Writer) {
	if len(filename1) < 1 || len(filename2) < 1 {
		// Missing content, so return.
		return
	}

	mergeFiles(filename1, filename2, output)
}

func mergeFiles(filename1, filename2 string, output io.Writer) {
	data1, err := os.Open(filename1)
	check(err)
	data2, err := os.Open(filename2)
	check(err)

	// Initialize readers for text and whitespace files.
	textReader := bufio.NewScanner(bufio.NewReader(data1))
	// textReader defaults to bufio.ScanLines
	wsReader := bufio.NewScanner(bufio.NewReader(data2))
	wsReader.Split(bufio.ScanLines)

	// Initialize output writer.
	writer := bufio.NewWriter(output)

	// To merge visible code + whitespace code together,
	// we need more whitespace characters than visible tokens.
	// Otherwise, we don't have enough delimiters!
	for wsReader.Scan() {
		textReader.Scan()

		// bufio.ScanLines removes the trailing LF, so add it back to
		// the whitespace input because it is syntactically significant
		wsLine := wsReader.Text() + "\n"

		// While there isn't enough whitespace to delimit
		// text, dump the whitespace and try the next line.
		numTokens := len(scanWords(textReader.Text()))
		numWS := len(scanWhitespace(wsLine))
		for numTokens >= numWS {
			fmt.Fprint(output, wsLine)
			wsReader.Scan()
			wsLine = wsReader.Text() + "\n"
			numWS = len(scanWhitespace(wsLine))
		}

		fmt.Fprint(output, interleave(textReader.Text(), wsLine))
	}

	// Flush the output.
	writer.Flush()
}

// interleave takes two strings containing text and whitespace tokens.
// It returns a string with the interleaved characters.
func interleave(tokens, whitespace string) string {
	wordScanner := bufio.NewScanner(strings.NewReader(tokens))
	wordScanner.Split(bufio.ScanWords)
	wsScanner := bufio.NewScanner(strings.NewReader(whitespace))
	wsScanner.Split(bufio.ScanBytes)

	output := ""
	for wsScanner.Scan() {
		wordScanner.Scan()
		output += wordScanner.Text()
		output += wsScanner.Text()
	}
	return output
}

// scanWhitespace reads a string and tokenizes whitespace.
// It returns a slice of strings containing the whitespace chars.
func scanWhitespace(line string) []string {
	wsSlice := []string{}
	wsScanner := bufio.NewScanner(strings.NewReader(line))
	wsScanner.Split(bufio.ScanBytes)

	for wsScanner.Scan() {
		wsSlice = append(wsSlice, wsScanner.Text())
	}
	return wsSlice
}

// scanWords reads a string an tokenizes the words using whitespace as a delimiter.
// It returns a slice of strings containing the tokens.
func scanWords(line string) []string {
	wordSlice := []string{}
	wordScanner := bufio.NewScanner(strings.NewReader(line))
	wordScanner.Split(bufio.ScanWords)

	for wordScanner.Scan() {
		wordSlice = append(wordSlice, wordScanner.Text())
	}
	return wordSlice
}

// check accepts an error and panics if an error has occurred.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
