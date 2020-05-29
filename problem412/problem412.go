package problem412

import "strconv"

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return result
}
