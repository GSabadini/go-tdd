package array

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int  {
	var result int
	
	for _, v := range numbers {
		result += v
	}

	return result
}

// SumAll calculates total of more than one slice of numbers 
func SumAll(numbersOfSum ...[]int) []int {
	var sum []int

	for _, v := range numbersOfSum {
		sum = append(sum, Sum(v))
	}

	return sum
}

// SumAllTails calculates the sums of all but the first number given a collection of slices
func SumAllTails(numbersOfTail ...[]int) []int {
	var sum []int

	for _, v := range numbersOfTail {
		if len(v) == 0 {
			sum = append(sum, 0)
		} else {
			last := v[1:]
			sum = append(sum, Sum(last))
		}
	}

	return sum
}