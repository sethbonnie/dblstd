package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissing(t *testing.T) {
	testCases := []struct {
		name     string
		shape    []byte
		root     string
		expected map[string]bool
	}{
		{
			name: "NoneMissing",
			root: "./testdata",
			shape: []byte(`
				README.md
				LICENSE
				.github/workflows/
				go.mod
			`),
			expected: map[string]bool{},
		},
		{name: "MissingDir",
			root: "./testdata",
			shape: []byte(`
				README.md
				LICENSE
				go.mod
				.vscode/
			`),
			expected: map[string]bool{
				".vscode": true,
			},
		},
		{name: "MissingFile",
			root: "./testdata",
			shape: []byte(`
				README.md
				LICENSE
				go.mod
				.env
			`),
			expected: map[string]bool{
				".env": false,
			},
		},
		{name: "MissingNestedFile",
			root: "./testdata",
			shape: []byte(`
				README.md
				LICENSE
				go.mod
				.vscode/settings.json
			`),
			expected: map[string]bool{
				".vscode/settings.json": false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewShape(tc.shape)
			assert.Nil(t, err)
			actual, err := s.Missing(tc.root)
			assert.Nil(t, err)

			assert.Len(t, actual, len(tc.expected))
			for path := range actual {
				assert.Equal(t, actual[path], tc.expected[path])
			}
		})
	}
}
