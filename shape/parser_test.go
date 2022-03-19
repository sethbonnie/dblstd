package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name   string
		shape  []byte
		result map[string]bool
	}{
		{
			name: "NoBlankLines",
			shape: []byte(`
				README.md
				LICENSE
				.github/workflows/
				go.mod
			`),
			result: map[string]bool{
				"README.md":         false,
				"LICENSE":           false,
				".github/workflows": true,
				"go.mod":            false,
			},
		},
		{
			name: "WithBlankLines",
			shape: []byte(`
					README.md

					LICENSE


					.github/workflows/
					
					go.mod
				`),
			result: map[string]bool{
				"README.md":         false,
				"LICENSE":           false,
				".github/workflows": true,
				"go.mod":            false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Parse(tc.shape)
			assert.Nil(t, err)
			assert.Equal(t, len(result), len(tc.result))
			for path, isDir := range result {
				_, ok := tc.result[path]
				assert.True(t, ok)
				assert.Equal(t, tc.result[path], isDir)
			}
		})
	}
}
