package main

import "fmt"

func main() {
	flotingPoint()
	// stringOperations()

}

func stringOperations() {
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

func flotingPoint() {
	var f float32 = 13.10

	// var g float64 = 3.14159265358979323846
	var h float64 = float64(f)
	fmt.Println(f)
	fmt.Printf("f: %f, g: %f\n", f, h)
	fmt.Printf("actual ff: %v,actual h: %v\n", f, h)

	fmt.Printf("f: %.2f, g: %.3f\n", f, h)
	fmt.Printf(" :%f \n", h)
	var j float32 = 13.8
	var k float64 = 13.8
	var l float64 = float64(j)

	fmt.Printf("j: %v \n", j)
	fmt.Printf("k: %v %.20f \n", k, k)

	fmt.Printf("l: %v and %.200f \n", l, l)

	m := 42.0
	n := 100 / 3.0
	fmt.Printf("first m: %f \n", m)
	fmt.Printf("second m: %.2f \n", m)
	fmt.Printf("3rd m: %9f \n", m)
	fmt.Printf("n: %f \n", n)
	fmt.Printf("n: %.2f \n", n)
	fmt.Printf("n: %.90f \n", n)

	fmt.Printf("|%9f|\n", 4.0)
	fmt.Printf("|%9f|\n", 123456.9)
	fmt.Printf("|%9f|\n", 1234567890999)
	fmt.Printf("|%9.2f|\n", 199.1)

	word := "GoLang"

	for _, ch := range word {
		// fmt.Printf("Index: %d, Character: %c\n", i, ch)
		// fmt.Printf("Index: %d, dec: %d\n", i, ch)
		// fmt.Printf("Index: %d, uni: %U\n", i, ch)
		fmt.Printf("%b ", ch)

	}

}
