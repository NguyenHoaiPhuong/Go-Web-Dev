package utils

// SliceContainsString return true if slice contains string
func SliceContainsString(slice []string, str string) bool {
	for _, elem := range slice {
		if elem == str {
			return true
		}
	}
	return false
}
