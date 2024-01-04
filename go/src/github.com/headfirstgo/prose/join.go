package prose

import "strings"

func JoinWithCommas(pharases []string) string {
	if len(pharases) == 1 {
		return pharases[0]
	} else 
	if len(pharases) == 2 {
		return pharases[0] + " and " + pharases[1]
	} else {
		result := strings.Join(pharases[:len(pharases)-1], ", ")
		result += ", and "
		result += pharases[len(pharases)-1]
		return result
	}
}
