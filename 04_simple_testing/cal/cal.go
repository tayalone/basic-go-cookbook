package cal

func SumOfFirst(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}
