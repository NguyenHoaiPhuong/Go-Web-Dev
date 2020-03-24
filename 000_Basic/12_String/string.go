package main

// RandomString : generate a string from strChars string
func RandomString(length int, strChars string) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(strChars)
	fmt.Println(chars)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
