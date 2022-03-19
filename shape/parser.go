package shape

import (
	"bufio"
	"bytes"
	"strings"
)

// Parse takes a slice of bytes representing a shape and returns map of paths.
// If the path represents a directory the value is true, otherwise false.
func Parse(f []byte) (map[string]bool, error) {
	// Keeps track of the directory we're in
	paths := make(map[string]bool)
	reader := bytes.NewReader(f)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Ignore empty lines
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, "/") {
			line = strings.TrimSuffix(line, "/")
			paths[line] = true
		} else {
			paths[line] = false
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return paths, nil
}
