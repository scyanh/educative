package slidingwindow

// LongestSubstringWithDistinctCharacters the longest substring with distinct characters
// Example 1:
// input: "abbbb"
// output: 2
// intput: "aabccbb"
// output: 3
// Time Complexity: O(n)
// Space Complexity: O(1) // only one map is used to store only alphabetical characters
func (slidingWindowChallenge) LongestSubstringWithDistinctCharacters(s string) int {
	n := len(s) // length of string
	if n == 0 {
		return 0 // if string is empty, return 0
	}

	left := 0               // left pointer
	res := 0                // result
	m := make(map[rune]int) // map of characters and their counts

	for i, el := range s { // iterate through string
		if m[el] > 0 { // if character is already in map
			m[el]++ // increment char in map
		} else {
			m[el] = 1 // else add new character to map
		}

		for m[el] > 1 { // if character right is already in map
			rLeft := rune(s[left]) // convert left character to rune
			m[rLeft]--             // decrement count
			if m[rLeft] == 0 {
				delete(m, rLeft) // delete character from map
			}
			left++ // move left pointer
		}

		res = max(res, i+1-left) // update result
	}

	return res // return result
}
