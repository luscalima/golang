package main

import (
	"fmt"
)

func main() {
	cipherText := "LXFOPVEFRNHR"
	key := "LEMON"
	
	// len Returns the length of the string in bytes
	textLen := len(cipherText)
	keyLen := len(key)

	var plainText string

	for i := 0; i < textLen; i++ {
		shift := int8(cipherText[i]) - int8(key[i % keyLen])
		c := 'A' + (26 + shift) % 26
		plainText += string(c)
	}

	fmt.Println(plainText)
}
