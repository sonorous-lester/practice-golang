package array_slice

func Sum(arr []int) int {
	var result int
	for _, val := range arr {
		result += val
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTail(numbers ...[]int) []int {
	var result []int
	for _, number := range  numbers{
		if len(number) != 0 {
			tail := number[1:]
			result = append(result, Sum(tail))
		}else {
			result = append(result, 0)
		}
	}
	return result
}
