package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissing(t *testing.T) {
	shape := []byte(`
		README.md
		LICENSE
		.github/workflows/
		go.mod
	`)

	testCases := []struct {
		name     string
		shape    []byte
		root     string
		depth    int
		expected []string
	}{
		{
			name:     "NoneMissingDefaultDepth",
			root:     "./testdata",
			shape:    shape,
			depth:    0,
			expected: nil,
		},
		{
			name:     "NoneMissingHighDepth",
			root:     "./testdata",
			shape:    shape,
			depth:    10,
			expected: nil,
		},
		{name: "MissingNestedDir",
			root:  "./testdata",
			shape: shape,
			depth: 1,
			expected: []string{
				".github/workflows",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Missing(tc.root, tc.shape, tc.depth)
			assert.Nil(t, err)

			assert.Len(t, actual, len(tc.expected))
		})
	}
}
