package main

func Intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem]++
		}
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem]--
			result = append(result, elem)
		}
	}
	return result
}
