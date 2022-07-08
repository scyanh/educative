package slidingwindow

// LongestSubstringWithKDistinctCharacters finds the longest substring with k distinct characters
// Example 1:
// Input: s = "araaci", k = 2
// Output: 4
// Time Complexity: O(n+n) = O(n)
// Space Complexity: O(k) = O(n)
func (slidingWindowChallenge) LongestSubstringWithKDistinctCharacters(s string, k int) int {
	n := len(s) // length of string
	if n == 0 {
		return 0 // if string is empty, return 0
	}

	left := 0 // left pointer
	res := 0 // result
	m := make(map[rune]int) // map of characters and their counts

	for i, el := range s { // iterate through string
		m[el]++ // increment char in map

		for len(m) > k { // if count is greater than k
			r := rune(s[left]) // convert character to rune
			if _, ok := m[r]; ok { // if character is already in map
				m[r]-- // decrement count
				if m[r] == 0 {
					delete(m, r) // delete character from map
				}
			}
			left++ // move left pointer
		}

		res = max(res, i+1-left) // update result
	}

	return res // return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
