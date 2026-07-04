package html

import (
	"path/filepath"
	"strings"
)

func isImage(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func add(a, b int) int {
	return a + b
}
