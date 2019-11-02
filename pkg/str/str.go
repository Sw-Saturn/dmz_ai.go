package str

import "strings"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if strings.Contains(x, n) {
			return true
		}
	}
	return false
}