package rsacipher

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var rsaTestData = []struct {
	description string
	input       string
}{
	{
		"Encrypt letter 'a'",
		"a",
	},
	{
		"Encrypt 'hello world'",
		"hello world",
	},
	{
		"Encrypt full sentence",
		"the quick brown fox jumps over the lazy dog.",
	},
	{
		"Encrypt full sentence from rsacipher.go main function",
		"I think RSA is really great",
	},
}

func TestModularExponentiation(t *testing.T) {
	t.Run("Test for modular exponentiation", func(t *testing.T) {
		expected := 2
		actual := modularExponentiation(2, 5, 3)
		if actual != expected {
			t.Errorf("Actual: %d and Expected: %d", actual, expected)
		}
	})
	t.Run("Test 2 for modular exponentiation", func(t *testing.T) {
		expected := 1099
		actual := modularExponentiation(200405, 5, 1342)
		if actual != expected {
			t.Errorf("Actual: %d and Expected: %d", actual, expected)
		}
	})
}

// This test is not consistent - consider commenting out until we find the problem
func TestRSACipher(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	bits := 17

	p := generatePrimes(1 << bits)
	q := generatePrimes(1 << bits)
	for p == q {
		q = generatePrimes(1 << bits)
	}
	n := p * q

	delta := lcm(p-1, q-1)

	e := generatePrimes(delta)
	d := modularMultiplicativeInverse(e, delta)

	for _, test := range rsaTestData {
		t.Run(test.description, func(t *testing.T) {

			message := []rune(test.input)
			expected := toASCII(message)

			encrypted := EncryptRSA(expected, e, n)
			decrypted := DecryptRSA(encrypted, d, n)

			if actual := decrypted; !reflect.DeepEqual(actual, expected) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %v, actual %v", expected, ToRune(actual))
			}
		})
	}
}
