package integers

func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	// using make
	// numLen := len(numsToSum)
	// sums := make([]int, numLen)
	// for i, n := range numsToSum {
	// 	sums[i] = Sum(n)
	// }
	// return sums

	// using append
	var sums []int
	for _, n := range numsToSum {
		sums = append(sums, Sum(n))
	}
	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, n := range numsToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			tail := n[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}