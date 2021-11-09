package Sum_of_Digits_of_String_After_Convert

import (
	"math"
	"strconv"
)

/*
	Create a few different functions

	1. a string convert to digits
	2. a transform func, which takes a number and adds together
	3. a transform loop, which takes a k and calls transform on a loop

	Then combine these two in the top level functions
*/

func getLucky(s string, k int) int {
	digits := stringToDigits(s)
	return transformKTimes(digits, k)
}

func stringToDigits(s string) int {
	digitStr := ""
	for i := 0; i < len(s); i++ {
		numOfAlpha := int(s[i]) % 32
		numOfAlphaStr := strconv.Itoa(numOfAlpha)
		digitStr += numOfAlphaStr
	}
	digits, _ := strconv.Atoi(digitStr)
	return digits
}

func transform(n int) int {
	numberOfDigits := int(math.Floor(math.Log10(float64(n)))) + 1
	sum := 0
	for i := 0; i < numberOfDigits; i++ {
		temp := n
		temp = temp / int(math.Pow10(i))
		temp = int(math.Mod(float64(temp), 10))
		sum += temp
	}
	return sum
}

func transformKTimes(n int, k int) int {
	for i := 0; i < k; i++ {
		n = transform(n)
	}
	return n
}
