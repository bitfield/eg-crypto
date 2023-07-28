package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bitfield/shift"
)

func TestEncipherTransforms(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		key         byte
		input, want []byte
	}{
		{
			key:   1,
			input: []byte("HAL"),
			want:  []byte("IBM"),
		},
		{
			key:   2,
			input: []byte("SPEC"),
			want:  []byte("URGE"),
		},
		{
			key:   3,
			input: []byte("PERK"),
			want:  []byte("SHUN"),
		},
		{
			key:   4,
			input: []byte("GEL"),
			want:  []byte("KIP"),
		},
		{
			key:   7,
			input: []byte("CHEER"),
			want:  []byte("JOLLY"),
		},
		{
			key:   10,
			input: []byte("BEEF"),
			want:  []byte("LOOP"),
		},
	}
	for _, tc := range tcs {
		name := fmt.Sprintf("%s + %d = %s", tc.input, tc.key, tc.want)
		t.Run(name, func(t *testing.T) {
			got := shift.Encipher(tc.input, tc.key)
			if !bytes.Equal(tc.want, got) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
