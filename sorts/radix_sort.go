package sorts

func max(arr []int) int {
	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}

func countSort(arr []int, exp int) []int {
	var digits [10]int
	var output = make([]int, len(arr))

	for _, item := range arr {
		digits[(item/exp)%10]++
	}
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output[digits[(arr[i]/exp)%10]-1] = arr[i]
		digits[(arr[i]/exp)%10]--
	}

	return output
}

func unsignedRadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	maxElement := max(arr)
	for exp := 1; maxElement/exp > 0; exp *= 10 {
		arr = countSort(arr, exp)
	}
	return arr
}

func RadixSort(arr []int) []int {
	var negatives, nonNegatives []int

	for _, item := range arr {
		if item < 0 {
			negatives = append(negatives, -item)
		} else {
			nonNegatives = append(nonNegatives, item)
		}
	}
	negatives = unsignedRadixSort(negatives)

	// Reverse the negative array and restore signs
	for i, j := 0, len(negatives)-1; i <= j; i, j = i+1, j-1 {
		negatives[i], negatives[j] = -negatives[j], -negatives[i]
	}
	return append(negatives, unsignedRadixSort(nonNegatives)...)
}
