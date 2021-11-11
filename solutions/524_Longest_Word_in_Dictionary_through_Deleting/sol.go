package Longest_Word_in_Dictionary_through_Deleting

func findLongestWord(s string, dictionary []string) string {
	sol := ""
	for i := range dictionary {
		word := dictionary[i]
		end := len(word)
		wordIndex := 0
		for sIndex := range s {
			sChar := string(s[sIndex])
			wChar := string(word[wordIndex])
			if sChar == wChar {
				wordIndex++
			}
			if wordIndex >= end {
				// is solution

				if len(word) > len(sol) {
					sol = word
				} else if len(word) == len(sol) {
					if isNewSmaller(sol, word) {
						sol = word
					}
				}
				break
			}
		}
	}
	return sol
}

func isNewSmaller(first string, second string) bool {
	if first == "" {
		return true
	}
	return second < first
}
