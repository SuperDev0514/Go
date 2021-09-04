// millerrabinprimality_test.go
// description: Test for Miller-Rabin Primality Test
// author(s) [Taj](https://github.com/tjgurwara99)
// see millerrabinprimalitytest.go

package prime

import "testing"

func TestMillerRabinTest(t *testing.T) {
	var tests = []struct {
		name     string
		input    int64
		expected bool
		rounds   int64
		err      error
	}{
		{"smallest prime", 2, true, 5, nil},
		{"random prime", 3, true, 5, nil},
		{"neither prime nor composite", 1, false, 5, nil},
		{"random non-prime", 10, false, 5, nil},
		{"another random prime", 23, true, 5, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := MillerRabinTest(test.input, test.rounds)
			if err != test.err {
				t.Errorf("For input: %d, unexpected error: %v, expected error: %v", test.input, err, test.err)
			}
			if output != test.expected {
				t.Errorf("For input: %d, expected: %v", test.input, output)
			}
		})
	}
}

func BenchmarkMillerRabinPrimalityTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = MillerRabinTest(23, 5)
	}
}
