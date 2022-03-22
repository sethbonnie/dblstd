package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name      string
		shapeSpec []byte
		result    Paths
	}{
		{
			name: "NoBlankLines",
			shapeSpec: []byte(`
				README.md
				LICENSE
				.github/workflows/
				go.mod
			`),
			result: Paths{
				"README.md":         false,
				"LICENSE":           false,
				".github/workflows": true,
				"go.mod":            false,
			},
		},
		{
			name: "WithBlankLines",
			shapeSpec: []byte(`
					README.md

					LICENSE


					.github/workflows/
					
					go.mod
				`),
			result: Paths{
				"README.md":         false,
				"LICENSE":           false,
				".github/workflows": true,
				"go.mod":            false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := NewShape(tc.shapeSpec)
			assert.Nil(t, err)
			assert.Equal(t, len(s.paths), len(tc.result))
			for path, isDir := range s.paths {
				_, ok := tc.result[path]
				assert.True(t, ok)
				assert.Equal(t, tc.result[path], isDir)
			}
		})
	}
}
