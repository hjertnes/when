// Package utils contains various utility functions
package utils

import (
"os"
"strings"
)

// Exist Checks if file or folder exists.
func Exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReplaceTilde Replaces tilde with value of HOME.
func ReplaceTilde(input string) string {
	home := os.Getenv("HOME")
	return strings.Replace(input, "~", home, 1)
}

