package arrays

func GetCount(a []int, b int, c int) (int, int) {
	coolerCount := 0
	heaterCount := 0
	for i := 0; i < len(a); i++ {
		if a[i] < b {
			heaterCount = heaterCount + 1
		} else if a[i] > c {
			coolerCount = coolerCount + 1
		}

	}
	return  heaterCount,coolerCount

}
