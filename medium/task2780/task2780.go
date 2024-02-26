package task2780

func MinimumIndex(nums []int) int {
	if len(nums) == 1 {
		return -1
	}

	m := make(map[int]int)
	for _, el := range nums {
		m[el] += 1
	}
	dominant := 0
	for k, v := range m {
		if !isDominant(v, len(nums)-1) {
			delete(m, k)
		} else {
			dominant = k
		}

	}
	m2 := make(map[int]int)

	for i, el := range nums {
		m2[el] += 1
		if el == dominant && isDominant(m2[el], i+1) && isDominant(m[el]-1, len(nums)-i-1) {
			return i
		}
	}

	return -1
}

func isDominant(freq int, len int) bool {
	if freq*2 > len {
		return true
	}

	return false
}
