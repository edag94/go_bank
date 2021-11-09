package Sum_of_Digits_of_String_After_Convert

import (
	"strconv"
)

/*
	Create a few different functions

	1. a string convert to digits
	2. a transform func, which takes a number and adds together
	3. a transform loop, which takes a k and calls transform on a loop

	Then combine these two in the top level functions

	Keep everything as a string, to prevent overflow
*/

func getLucky(s string, k int) int {
	digits := stringToDigits(s)
	return transformKTimes(digits, k)
}

func stringToDigits(s string) string {
	digitStr := ""
	for i := 0; i < len(s); i++ {
		numOfAlpha := int(s[i]) % 32
		numOfAlphaStr := strconv.Itoa(numOfAlpha)
		digitStr += numOfAlphaStr
	}
	return digitStr
}

func transform(s string) string {
	sum := 0
	for _, c := range s {
		num, _ := strconv.Atoi(string(c))
		sum += num
	}
	return strconv.Itoa(sum)
}

func transformKTimes(s string, k int) int {
	for i := 0; i < k; i++ {
		s = transform(s)
	}
	num, _ := strconv.Atoi(s)
	return num
}
