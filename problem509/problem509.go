package problem509

// func fib(N int) int {
// 	if N < 2 {
// 		return N
// 	}

// 	return fib(N-1) + fib(N-2)
// }

func fib(N int) int {
	if N < 2 {
		return N
	}
	sum, a, b := 0, 0, 1
	for N > 1 {
		sum = a + b
		a = b
		b = sum
		N--
	}
	return sum
}
