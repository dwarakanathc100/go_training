package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main1() {
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

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	rune := "ellow"
	fmt.Println(rune)
	// fmt.Println("Command-line arguments:", os.Args)
	file, err := os.Open("golang-day1-intro.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter your city: ")

	// if scanner.Scan() {
	// 	city := scanner.Text() // gets the input line
	// 	fmt.Println("City:", city)
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// name, err := reader.ReadString('\n') // reads until newline7
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }

	// name = strings.TrimSpace(name) // remove \n
	// fmt.Println("Hello,", name)
	// // var name string
	// var age int

	// fmt.Print("Enter your name and age: ")
	// // fmt.Scan(&name, &age) // & because Go passes addresses to store input
	// // fmt.Scanln(&name, &age) // stops at newline
	// fmt.Scanf("%s %d", &name, &age) // formatted input

	// fmt.Printf("Hello %s, you are %d years old\n", name, age)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(99))
	// }
	// fmt.Println(n)
	// r, err := divide(10, 0)
	// fmt.Println(r)
	// fmt.Println(err)

}
