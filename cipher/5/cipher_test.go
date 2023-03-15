package cipher_test

import (
	"cipher"
	"testing"
)

type testCase struct {
	name        string
	input, want rune
}

func TestShift(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		input, want rune
	}{
		{
			name:  "turns 'A' into 'B'",
			input: 'A', want: 'B',
		},
		{
			name:  "wraps around the end of the alphabet",
			input: 'Z', want: 'A',
		},
		{
			name:  "leaves non-letters unchanged",
			input: ' ', want: ' ',
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := cipher.Shift(tc.input)
			if tc.want != got {
				t.Errorf("Shift(%q): want %q, but got %q", tc.input, tc.want, got)
			}
		})
	}
}

func TestEncipherCorrectlyEnciphersPlaintext(t *testing.T) {
	t.Parallel()
	input := "THESE PRETZELS ARE MAKING ME THIRSTY"
	want := "UIFTF QSFUAFMT BSF NBLJOH NF UIJSTUZ"
	got := cipher.Encipher(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
