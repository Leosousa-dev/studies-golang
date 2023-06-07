package main

import (
	"fmt"
)

func main() {
	var palavra string

	fmt.Print("✍️ Digite uma palavra.\n")
	fmt.Scan(&palavra)
	fmt.Printf("palavra: %s\n", ReverseStrings(palavra))
}

func ReverseStrings(str string) string {
	bytes := []byte(str)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	ReversedStr := string(bytes)
	return ReversedStr
}
