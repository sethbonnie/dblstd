package shape

import (
	"bufio"
	"bytes"
	"strings"
)

type Paths map[string]bool

type Shape struct {
	paths Paths
}

func NewShape(spec []byte) (*Shape, error) {
	paths, err := parse(spec)
	if err != nil {
		return nil, err
	}
	return &Shape{paths}, nil
}

// Parse takes a slice of bytes representing a shape and returns map of paths.
// If the path represents a directory the value is true, otherwise false.
func parse(spec []byte) (Paths, error) {
	// Keeps track of the directory we're in
	paths := Paths{}
	reader := bytes.NewReader(spec)
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
