package pipeline

func Double(numbers []int) []int {
	if numbers == nil {
		return nil
	}
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = n * 2
	}
	return result
}
