package main

import (
	"fmt"
	"strings"
)

// EJEMPLO
//
//	palindrome
func esPalindrome(text string) {
	var textReverso string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverso += string(text[i])
	}
	if text == textReverso {
		fmt.Println("Es Palindrome")
	} else {
		fmt.Println("No es palindrome")
	}
}
func main() {
	frase := "raDar"
	fmt.Println(frase)
	esPalindrome(frase)
}
