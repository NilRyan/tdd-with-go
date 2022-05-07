package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrNumbers ...[]int) []int {
	sums := []int{}
	for _, numbers := range arrNumbers {
		sums = append(sums, Sum(numbers))
	}
	return sums

}
