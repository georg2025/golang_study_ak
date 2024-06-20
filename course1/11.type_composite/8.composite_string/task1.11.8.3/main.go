package main

func main() {
}

func getBytes(s string) []byte {
	bytes := []byte{}
	for i := 0; i < len(s); i++ {
		bytes = append(bytes, s[i])
	}
	return bytes
}

func getRunes(s string) []rune {
	runes := []rune{}
	for _, i := range s {
		runes = append(runes, i)
	}
	return runes
}
