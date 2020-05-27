package arrays

func FindTwoNumbers(a []int, b int) (int, int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == b {
				return a[i], a[j]
			}
		}

	}
	return -1, -1
}
