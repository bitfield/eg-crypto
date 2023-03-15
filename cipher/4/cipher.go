package cipher

import "strings"

func Shift(input rune) rune {
	result := input + 1
	if result > 'Z' {
		result -= 26
	}
	return result
}

func Encipher(plaintext string) string {
	ciphertext := new(strings.Builder)
	for _, r := range plaintext {
		ciphertext.WriteRune(Shift(r))
	}
	return ciphertext.String()
}
