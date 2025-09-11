package main

import "fmt"

func main() {
	var s string = "Hello, World!"
	var literal string = `This is a raw string literal
	that can span multiple lines
	and includes "quotes" and \backslashes\ without needing escapes.`

	fmt.Println(s)
	fmt.Println(literal)
	f := '✅'
	fmt.Println(f)
	fmt.Printf("format %v specifier: %U \n ", f, f)
	fmt.Printf("format %v specifier: %c \n", f, f)

	fmt.Print(len(s))

}
