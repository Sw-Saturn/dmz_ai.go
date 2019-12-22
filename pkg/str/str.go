package str

import "strings"

//Contains is Substring detection
func Contains(a []string, x string) bool {
	for _, n := range a {
		if strings.Contains(x, n) {
			return true
		}
	}
	return false
}
