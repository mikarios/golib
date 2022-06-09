package stringtools

// SplitStringByLimit splits a given string by maximum of given number of characters.
// If number is <=0 it returns the whole string as the first item of the list.
func SplitStringByLimit(value string, limit int) []string {
	if limit <= 0 {
		return []string{value}
	}

	charSlice := make([]rune, 0)

	// Push characters to rune slice
	for _, char := range value {
		charSlice = append(charSlice, char)
	}

	chunkValue := make([]string, 0)

	for len(charSlice) >= 1 {
		// Check if remaining characters are less that requested limit
		if len(charSlice) < limit {
			limit = len(charSlice)
		}

		// Append chunk of value
		chunkValue = append(chunkValue, string(charSlice[:limit]))
		// Discard the elements that were copied over to result
		charSlice = charSlice[limit:]
	}

	return chunkValue
}
