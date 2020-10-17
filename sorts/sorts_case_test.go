//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"crypto/rand"
	"math/big"
	"sort"
)

type sortTest struct {
	input    []int
	expected []int
	name     string
}

var uarr []int = makeRandNonNegativeArray(500_000)
var arr []int = makeRandArray(500_000)

var sortTests = []sortTest{
	//Sorted slice
	{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted"},

	//Reversed slice
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed"},

	//Reversed slice, even length
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Even"},

	//Random order with repetitions
	{[]int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
		[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}, "Random order with repetitions"},

	//Empty slice
	{[]int{}, []int{}, "Empty"},
	//Single-entry slice
	{[]int{1}, []int{1}, "Singleton"},

	//500k non negative int values sort
	{uarr, getSortedVersion(uarr), "Large Random"},

	//500k int values sort
	{arr, getSortedVersion(arr), "Large Random2"},
}

func makeRandNonNegativeArray(size int) []int {
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		temp, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
		vals[i] = int(temp.Int64())
	}
	return vals
}

func makeRandArray(size int) []int {
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		temp, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
		vals[i] = int(temp.Int64()) - 250_000
	}
	return vals
}

func getSortedVersion(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return a
}
