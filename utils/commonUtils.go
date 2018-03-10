package utils

import (
	"strings"
)

func AllNotEmpty(s ...string) bool {
	for _, v := range s {
		if strings.TrimSpace(v) == "" {
			return false
		}
	}
	return true
}
