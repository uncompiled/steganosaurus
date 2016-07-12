package modules

import (
	"testing"
)

// TestScanWords tests whether the internal scanWords method tokenizes properly.
func TestScanWords(t *testing.T) {

	testString := "const figlet = require('figlet');"
	expectedString := [4]string{"const", "figlet", "=", "require('figlet');"}
	output := scanWords(testString)

	for i := range output {
		if len(output) != len(expectedString) {
			// Should be the same length
			t.Fail()
		}

		if output[i] != expectedString[i] {
			// Should be the same value
			t.Fail()
		}
	}
}

// TestScanWhitespace tests whether the internal scanWhitespace method tokenizes properly.
func TestScanWhitespace(t *testing.T) {

	testString := " \t \n"
	expectedString := [4]string{" ", "\t", " ", "\n"}
	output := scanWhitespace(testString)

	for i := range output {
		if len(output) != len(expectedString) {
			// Should be the same length
			t.Fail()
		}

		if output[i] != expectedString[i] {
			// Should be the same value
			t.Fail()
		}
	}
}
