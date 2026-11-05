package pipeline

func Double(in []int) []int {
	if in == nil {
		return nil
	}
	result := make([]int, len(in))
	for i, n := range in {
		result[i] = n * 2
	}
	return result
}
func FilterMin(in []int, min int) []int {
	if in == nil {
		return nil
	}
	result := make([]int, 0)
	for _, n := range in {
		if n >= min {
			result = append(result, n)
		}
	}
	return result
}
func UniqueStable(in []int) []int {
	if in == nil {
		return nil
	}
	seen := make(map[int]bool)
	result := make([]int, 0)
	for _, n := range in {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}
