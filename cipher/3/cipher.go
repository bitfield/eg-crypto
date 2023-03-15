package cipher

func Shift(input rune) rune {
	result := input + 1
	if result > 'Z' {
		result -= 26
	}
	return result
}
