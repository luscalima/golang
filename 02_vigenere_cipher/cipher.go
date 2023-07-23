package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "attack at dawn"
	key := "lemon"
	sanitizedText := strings.ToUpper(strings.ReplaceAll(plainText, " ", ""))
	sanitizedKey := strings.ToUpper(strings.ReplaceAll(key, " ", ""))
	
	// len Returns the length of the string in bytes
	textLen := len(sanitizedText)
	keyLen := len(sanitizedKey)


	var cipherText string

	for i := 0; i < textLen; i++ {
		c := 'A' + (sanitizedText[i] + sanitizedKey[i % keyLen]) % 26
		cipherText += string(c)
	}

	fmt.Println(cipherText)
}
