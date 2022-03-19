package shape

import (
	"os"
	"path/filepath"
	"strings"
)

// Missing walks the given repo path and compares its contents with the given
// shape file returning any items missing required paths.
// The depth argument controls how deep to walk through the repo. By default
// we walk through the whole repo.
func Missing(repoPath string, shape []byte, depth int) ([]string, error) {
	paths, err := walkDir(repoPath, depth)
	if err != nil {
		return nil, err
	}

	requiredPaths, err := Parse(shape)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)

	for _, path := range paths {
		if _, ok := requiredPaths[path]; ok {
			seen[path] = true
		}
	}

	if len(seen) < len(requiredPaths) {
		missing := []string{}
		for path := range requiredPaths {
			if _, ok := seen[path]; !ok {
				missing = append(missing, path)
			}
		}
		return missing, nil
	}
	return nil, nil
}

func walkDir(root string, depth int) ([]string, error) {
	paths := []string{}

	root = strings.TrimPrefix(root, "./")

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if depth > 0 && strings.Count(path, "/") > depth {
				return nil
			}
			path = strings.TrimPrefix(path, root+"/")
			paths = append(paths, path)
			return nil
		})
	if err != nil {
		return nil, err
	}
	return paths, nil
}
