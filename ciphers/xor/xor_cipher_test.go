package xor_test

import (
	"TheAlgorithms/Go/ciphers/xor"
	"fmt"
	"reflect"
	"testing"
)

func ExampleNewXor() {
	const (
		seed = "Hello World"
		key  = 97
	)

	x := xor.NewXor()
	encrypted := x.Encrypt(byte(key), []byte(seed))
	fmt.Printf("Encrypt=> key: %d, seed: %s, encryptedText: %v\n", key, seed, encrypted)

	decrypted := x.Decrypt(byte(key), encrypted)
	fmt.Printf("Decrypt=> key: %d, encryptedText: %v, DecryptedText: %s\n", key, encrypted, string(decrypted))

	// Output:
	// Encrypt=> key: 97, seed: Hello World, encryptedText: [41 4 13 13 14 65 54 14 19 13 5]
	// Decrypt=> key: 97, encryptedText: [41 4 13 13 14 65 54 14 19 13 5], DecryptedText: Hello World
}

var xorTestData = []struct {
	description string
	input       string
	key         int
	encrypted   string
}{
	{
		"Encrypt letter 'a' with key 0 makes no changes",
		"a",
		0,
		"a",
	},
	{
		"Encrypt letter 'a' with key 1",
		"a",
		1,
		"`",
	},
	{
		"Encrypt letter 'a' with key 10",
		"a",
		10,
		"k",
	},
	{
		"Encrypt 'hello world' with key 0 makes no changes",
		"hello world",
		0,
		"hello world",
	},
	{
		"Encrypt 'hello world' with key 1",
		"hello world",
		1,
		"idmmn!vnsme",
	},
	{
		"Encrypt 'hello world' with key 10",
		"hello world",
		10,
		"boffe*}exfn",
	},
	{
		"Encrypt full sentence with key 64",
		"the quick brown fox jumps over the lazy dog.",
		64,
		"4(%`15)#+`\"2/7.`&/8`*5-03`/6%2`4(%`,!:9`$/'n",
	},
	{
		"Encrypt a word with key 32 make the case swap",
		"abcdefghijklmNOPQRSTUVWXYZ",
		32,
		"ABCDEFGHIJKLMnopqrstuvwxyz",
	},
}

var x *xor.Xor = xor.NewXor()

func TestXorCipherEncrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {
			encrypted := x.Encrypt(byte(test.key), []byte(test.input))
			if !reflect.DeepEqual(string(encrypted), test.encrypted) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.encrypted, string(encrypted))
			}
		})
	}
}

func TestXorCipherDecrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {
			decrypted := x.Decrypt(byte(test.key), []byte(test.encrypted))

			if !reflect.DeepEqual(string(decrypted), test.input) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.input, string(decrypted))
			}
		})
	}
}
