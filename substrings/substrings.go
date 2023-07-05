package substrings

import "strings"

// FindIndices returns a slice of 1-based indices where the subtext is found in textToSearch
// Matching is case-insensitive, returns nil if no matches are found
// Matching algorithm runs in O(n*m) time, where n is the length of textToSearch and m is the length of subtext
// Only works for ASCII characters
func FindIndices(textToSearch, subtext string) []int {
	if len(subtext) > len(textToSearch) {
		//  nothing to do, early exit
		return nil
	}
	if len(textToSearch) == 0 {
		//  nothing to do, early exit
		return nil
	}
	if len(subtext) == 0 {
		//  nothing to do, early exit
		return nil
	}
	var indices []int // no idea how many indices we'll need
	// lowercase both textToSearch and subtext. This is already optimised for ASCII
	textToSearch = strings.ToLower(textToSearch)
	subtext = strings.ToLower(subtext)

	subTextLen := len(subtext)

	// we only need to check up to the last possible index where the subtext could be found
	for i := 0; i < len(textToSearch)-subTextLen+1; i++ {
		// check if the 1st character of the subtext matches the current character
		if textToSearch[i] == subtext[0] { // note that we are comparing bytes and not runes, hence ASCII only
			for j := 0; j < subTextLen; j++ {
				if textToSearch[i+j] != subtext[j] {
					// if the characters don't match, stop checking
					break
				}
				if j == subTextLen-1 {
					// we have reached the end of the substring, so
					//everything must have matched. Add the index to our indices slice
					indices = append(indices, i+1) // add one to make result 1-based
				}
			}
		}
	}
	return indices
}
