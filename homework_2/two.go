package main

func Intersection(a, b []int) []int {
	var result []int
	for i, v := range a {
		for j, t := range b {
			if v == t {
				result = append(result, v)
				i++
				break
			} else if v != t {
				j++
			}
		}
	}
	return result
}
