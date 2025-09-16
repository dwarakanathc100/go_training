# Go Language Concepts: Loops, Floating Formats, and Functions

------------------------------------------------------------------------

## 1. Nested Loops

Nested loops are useful when we need to perform **repeated iterations
inside another loop**.\
Common use cases: - Working with 2D
arrays or matrices - Comparing elements from two collections

### Example: Multiplication Table

``` go
for i := 1; i <= 10; i++ { // Outer loop
    for j := 1; j <= 15; j++ { // Inner loop
        fmt.Printf("%5d ", i*j)
    }
    fmt.Println() // New line after each row
}
```

This prints a **10x15 multiplication table**.

#### Exercises

1.  Write a program to print a **5x5 square** of `*` using nested
    loops.
2.  Generate a **right-angled triangle** pattern with numbers.


------------------------------------------------------------------------

## 2. Range Loops

Range loops are useful for **iterating over collections** (arrays,
slices, maps, strings).\
They simplify accessing both index and value.

### Example 1: Iterating over an Array

``` go
arr := []int{10, 20, 30, 40, 50}
for index, value := range arr {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

### Example 2: Iterating over a String 

``` go
words := []string{"go", "is", "awesome"}
for _, word := range words {
    for j, char := range word {
        fmt.Printf("\tIndex: %d, Char: %c\n", j, char)
    }
}
```

#### Exercises

1.  Create a slice of integers and print only the **even numbers** using
    a range loop.(if you want you can skip this until we discuss slice)
2.  Write a program that counts how many times each **character
    appears** in a string.

------------------------------------------------------------------------

## 3. Floating-Point Formatting

Formatting helps in **controlling decimal places, width, and
alignment**.\
This is especially useful in **financial calculations, reports, and
tabular data**.

### Example:

``` go
n := 100 / 3.0
fmt.Printf("Value: %f\n", n)         // Default floating-point format
fmt.Printf("Value: %9f\n", n)        // Width of 9 characters
fmt.Printf("Value: %.2f\n", n)       // 2 decimal places
fmt.Printf("Value: %09.2f\n", n)     // Width 9, padded with zeros
fmt.Printf("Value: |%013.2f|\n", n)  // Width 13, padded, 2 decimals
```

Output (example):

    Value: 33.333333
    Value:  33.333333
    Value: 33.33
    Value: 000033.33
    Value: |000000033.33|

#### Exercises

1.  Print the result of `22.0 / 7` with **2, 4, and 6 decimal places**.\
2.  Display a floating-point number with **zero-padding and fixed
    width**.

------------------------------------------------------------------------

## 4. Functions in Go


A **function** is a block of code that performs a specific task.\
It improves **code reusability, readability, and modularity**.

### Syntax:

``` go
func functionName(parameter1 type, parameter2 type) (returnType1, returnType2) {
    // function body
    return returnValue1, returnValue2
}
```

### Examples

#### (a) Zero Arguments, No Return Value

``` go
func say() {
    fmt.Println("Hello, World!")
}
```

#### (b) Zero Arguments, One Return Value

``` go
func greet() string {
    return "Hello, Gopher!"
}
```

#### (c) Two Arguments, One Return Value

``` go
func add(a int, b int) int {
    return a + b
}
```

#### (d) Infinite Arguments (Variadic Function)

``` go
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### Example Usage in `main()`

``` go
say()
s := greet()
fmt.Println(s)
sum := add(3, 5)
fmt.Println("Sum:", sum)
total := sumAll(1, 2, 3, 4, 5)
fmt.Println("Total Sum:", total)
```

#### Exercises

1.  Write a function `multiply(a, b int) int` that returns the product
    of two numbers.
2.  Write a variadic function `average(nums ...float64) float64` that
    calculates the average of numbers.
3.  Implement a function that **reverses a string** and returns it.

------------------------------------------------------------------------

