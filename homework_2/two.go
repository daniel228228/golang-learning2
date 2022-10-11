package main

func Intersection(a, b []int) []int {
	var result []int
	for _, v := range a {
		for _, t := range b {
			if v == t {
				result = append(result, v)
				break
			}
		}
	}
	return result
}
