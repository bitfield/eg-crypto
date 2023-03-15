package cipher_test

import (
	"cipher"
	"testing"
)

func TestShiftTransformsAToB(t *testing.T) {
	t.Parallel()
	input := 'A'
	want := 'B'
	got := cipher.Shift(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestShiftTransformsZToA(t *testing.T) {
	t.Parallel()
	input := 'Z'
	want := 'A'
	got := cipher.Shift(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEncipherCorrectlyEnciphersPlaintext(t *testing.T) {
	t.Parallel()
	input := "THESE PRETZELS ARE MAKING ME THIRSTY"
	want := "UIFTF!QSFUAFMT!BSF!NBLJOH!NF!UIJSTUZ"
	got := cipher.Encipher(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
