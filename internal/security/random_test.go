package security

import "testing"

func TestRandomBytes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		n    int
	}{
		{
			name: "8 bytes",
			n:    8,
		},
		{
			name: "32 bytes",
			n:    32,
		},
		{
			name: "64 bytes",
			n:    64,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b := RandomBytes(tt.n)
			if len(b) != tt.n {
				t.Fatalf("expected len(b) == %d, but got %d", tt.n, len(b))
			}
		})
	}
}
