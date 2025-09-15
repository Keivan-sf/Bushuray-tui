package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetBinPath(name string) (string, error) {
	paths := []string{
		"./" + name,
		filepath.Join(".", "bin", name),
		filepath.Join("/usr/bin", name),
		filepath.Join("/usr/local/bin", name),
	}

	for _, p := range paths {
		if info, err := os.Stat(p); err == nil && !info.IsDir() {
			return filepath.Abs(p)
		}
	}

	return "", fmt.Errorf("binary %q not found", name)
}
