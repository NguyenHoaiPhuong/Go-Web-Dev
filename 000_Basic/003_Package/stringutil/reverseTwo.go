package stringutil

// reverseTwo is unable to be exported because it starts with lowercase letter
// reverseTwo is able to be used by other functions inside package stringutil
func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
