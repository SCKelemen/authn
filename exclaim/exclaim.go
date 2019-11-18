package exclaim

import "strings"

// Exclaim yells your message louder
func Exclaim(message string) string {
	return strings.ToUpper(message) + "!!!"
}
