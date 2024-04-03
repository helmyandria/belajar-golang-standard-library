package main

import "fmt"

func main() {
	text1 := "Go"
	text2 := "Lang"

	fmt.Println("Hello '", text1, text2, "'")

	fmt.Printf("Hello '%s %s'\n", text1, text2)

}
