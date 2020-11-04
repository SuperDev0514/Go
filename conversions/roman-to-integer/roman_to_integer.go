/*
This algorithm will convert a standard roman number to an integer
https://en.wikipedia.org/wiki/Roman_numerals
Function receives a string as a roman number and outputs an integer
Maximum output will be 3999
Only standard form is supported
*/
package conversions

var romans = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func RomanToInteger(roman string) int {
	total := 0
	holder := 0
	for holder < len(roman) {
		if holder+1 < len(roman) && (romans[string(roman[holder])] < romans[string(roman[holder+1])]) {
			total += romans[string(roman[holder+1])] - romans[string(roman[holder])]
			holder += 2
		} else {
			total += romans[string(roman[holder])]
			holder++
		}
	}
	return total
}
