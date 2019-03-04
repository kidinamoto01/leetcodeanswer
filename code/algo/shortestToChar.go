package algo
import "math"
func shortestToChar(S string, C byte) []int {
	length := len(S)
	ret := make([]int, length)
	prev := math.MinInt64/2
	for i:= 0; i < length; i++ {

		if C == S[i] {
			prev = i
		}


		ret[i] = i-prev
	}

	prev = math.MaxInt64/2

	for i := length-1;i>=0;i--{
		if C == S[i] {
			prev = i
		}

		ret[i] = min(ret[i],prev-i)
	}
	return ret
}

func min(i int,j int) int{
	if i>j{
		return j
	}else{
		return i
	}
}
