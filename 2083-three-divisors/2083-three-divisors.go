func isThree(n int) bool {
	var result int

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			result++
		} else if 3 < result {
			break
		}
	}

	return result == 3
}