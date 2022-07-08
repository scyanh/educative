package slidingwindow

// MaxFruitsIntoBasket returns the maximum number of fruits in the basket
/*
Each basket can have only one type of fruit. There is no limit to how many fruit a basket can hold.
You can start with any tree, but you can’t skip a tree once you have started.
You will pick exactly one fruit from every tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

input: Fruit=['A', 'B', 'C', 'A', 'C']
output: 3
Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']
*/
func (slidingWindowChallenge) MaxFruitsIntoBasket(fruits []string) int {
	res := 0
	left := 0 // left pointer

	var m = make(map[string]int)

	for i, el := range fruits {
		m[el]++

		for len(m) > 2 {
			m[fruits[left]]--
			if m[fruits[left]] == 0 {
				delete(m, fruits[left])
			}
			left++
		}

		res = max(res, i+1-left)
	}

	return res
}
