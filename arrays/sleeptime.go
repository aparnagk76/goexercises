package arrays
import "sort"

func GetSleepTime(a[]int) (int,int,int) {
	var sleepTime[]int 
	alen := len(a)
	for i ,j,k :=0,1,0; i < (alen -1) && j <(alen); i,j,k = i+2,j+2,k+1 {
		if a[i] > a[j] {
			sleepTime[k] = 24 - a[i] + a[j]
		} else if a[i] < a[j] {
			sleepTime[k] = 24 - a[j] + a[i]
		}
	}
	sort.Ints(sleepTime)

	sum := 0

	for _, value := range sleepTime{
		sum = sum + value

	}
	klen := len(sleepTime)
	average := sum/klen
	return sleepTime[0],sleepTime[klen-1],average
}