package stringManipulator

import (
	"strconv"
	"strings"
)

// Compress() compresses strings
func Compress(s string) string {
	// Return s when string is empty
	if len(s) < 1 {
		return s
	}

	// Split string to array for simplified iteration and preservation of special characters
	var stringArr = strings.Split(s, "")
	var compressedString string
	var prev string
	var count int
	for i := 0; i < len(stringArr); i++ {
		char := stringArr[i]
		if char != prev || count >= 9 {
			compressedString += prev
			if count > 1 {
				compressedString += strconv.Itoa(count)
			}
			prev, count = char, 0
		}
		count++
	}
	compressedString += prev
	if count > 1 {
		compressedString += strconv.Itoa(count)
	}

	return compressedString
}

func Unpack(s string) string {
	// Return s when string is empty
	if len(s) < 1 {
		return s
	}

	var stringArr = strings.Split(s, "")
	var unpackedString string
	var prev string
	for i := 0; i < len(stringArr); i++ {
		char := stringArr[i]
		n, err := strconv.Atoi(char)
		if err == nil && prev != "" {
			for j := 0; j < n; j++ {
				unpackedString += prev
			}
			prev = ""
		} else {
			unpackedString += prev
			prev = char
		}
	}
	unpackedString += prev
	return unpackedString
}
