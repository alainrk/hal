package utils

// Unique removes duplicate strings from a slice.
func Unique(s []string) []string {
	inResult := make(map[string]struct{})
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = struct{}{}
			result = append(result, str)
		}
	}
	return result
}
