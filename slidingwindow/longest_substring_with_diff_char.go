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
	if s == "" {
		return 0
	}

	left := 0
	res := 0
	m := make(map[rune]int)

	for i, el := range s {
		if _, ok := m[el]; ok {
			// this is tricky; in the current window, we will not have any 'rightChar' after its previous index
			// and if 'windowStart' is already ahead of the last index of 'rightChar', we'll keep 'windowStart'
			left=max(left, m[el]+1) // left is the index of the first occurrence of the current character
		}

		m[el] = i
		res = max(res, i-left+1)// i-left+1 is the length of the substring
	}

	return res
}
