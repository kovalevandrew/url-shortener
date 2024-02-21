package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test for learning purpose
func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size 6",
			size: 6,
		},
		{
			name: "size 10",
			size: 10,
		},
		{
			name: "size 15",
			size: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewRandomString(tt.size)
			strUniq := NewRandomString(tt.size)
			assert.Len(t, str, tt.size)
			assert.Len(t, strUniq, tt.size)
			assert.NotEqual(t, str, strUniq)
		})
	}
}
