package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func repoRoot() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to generate package path")
	}

	return filepath.Dir(filepath.Dir(filename)), nil
}

func ChallengePath() (string, error) {
	root, err := repoRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, "core", "cmd", "cmd.go"), nil
}

func PkgPath(day int) (string, error) {
	root, err := repoRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, "core", fmt.Sprintf("day%d", day)), nil
}
