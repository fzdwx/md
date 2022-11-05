package utils

import (
	"runtime"
	"strings"
)

const (
	lf   = "\n"
	crlf = "\r\n"
)

// CleanCr if os is windows, newline symbol is `\r\n`, replace to '\n'
// because textarea.Model use '\n'
func CleanCr(body string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(body, crlf, lf)
	}
	return body
}
