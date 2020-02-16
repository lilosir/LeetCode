package problem1281

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for n > 0 {
		last := n % 10
		product = product * last
		sum = sum + last
		n = n / 10
	}
	return product - sum
}
