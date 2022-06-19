package recursion

type recursionPalindrome struct {
}

func NewRecursionPalindrome() recursionPalindrome {
	return recursionPalindrome{}
}

func (r recursionPalindrome) IsPalindrome(word string) bool {
	lenght := len(word)
	if lenght < 2 {
		return true
	}

	if word[0] != word[lenght-1] {
		return false
	}

	newWord := word[1 : lenght-1]
	return r.IsPalindrome(newWord)
}
