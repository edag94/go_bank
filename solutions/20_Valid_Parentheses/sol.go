package Valid_Parentheses

func isValid(s string) bool {
	stack := make([]string, 0)
	strMap := map[string]string{"[": "]", "{": "}", "(": ")"}
	for i := range s {
		char := string(s[i])
		_, contains := strMap[char]
		if contains {
			// must be forward
			stack = append(stack, char)
		} else {
			// must be backwards
			if len(stack) == 0 {
				// unbalanced end
				return false
			} else {
				lastElt := len(stack) - 1
				top := stack[lastElt]
				stack = stack[:lastElt]
				// if not match
				if strMap[top] != char {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}
