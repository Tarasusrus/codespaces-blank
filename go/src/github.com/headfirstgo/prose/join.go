package prose

import "strings"

func JoinWithCommas(pharases []string) string {
	result := strings.Join(pharases[:len(pharases) -1], ", ")
	result += " and "
	result += pharases[len(pharases)-1]
	return result
}