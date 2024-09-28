package reverse

func Reverse(input string) string {
	runes := []rune(input)
    for i := 0; i < len(runes)/2; i++ {
        runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
    }
    return string(runes)
}
