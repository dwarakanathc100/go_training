
# Golang Interview questions


## 1. Basics

**Q1. What is Go and what are its advantages?**  
Go is a statically typed, compiled language developed at Google. It emphasizes simplicity, readability, and efficient concurrency. Advantages include fast compilation, built-in garbage collection, native support for concurrency via goroutines and channels, and strong standard libraries.

**Q2. How do you declare variables in Go?**  
Variables can be declared using `var` or shorthand `:=`. `var x int = 10` declares with type, while `y := 20` infers type automatically. Go variables have zero values if uninitialized.

**Q3. What is the difference between constants and variables?**  
Constants are immutable values defined with `const` and must be compile-time expressions. Variables can change at runtime and are declared with `var` or `:=`.

**Q4. Explain the zero value in Go.**  
Zero values are default values assigned to uninitialized variables: 0 for numbers, false for booleans, "" for strings, and nil for pointers, slices, maps, channels, functions, and interfaces.

**Q5. What are packages in Go?**  
A package is a collection of Go source files in the same directory. The `main` package is the entry point. Packages promote modularity and code reuse.

**Q6. Explain defer, panic, and recover.**  
`defer` postpones function execution until surrounding function returns. `panic` stops normal execution and unwinds the stack. `recover` can regain control within a deferred function to handle panics gracefully.

## 2. Data Types and Control Flow

**Q7. What are basic data types in Go?**  
Go supports integers, floats, complex numbers, booleans, strings, arrays, slices, maps, structs, pointers, and interfaces.

**Q8. Explain arrays and slices.**  
Arrays have fixed size and are value types. Slices are dynamic, built on arrays, and reference underlying array data. Slices have length and capacity and can be expanded using `append`.

**Q9. What are maps in Go?**  
Maps are unordered collections of key-value pairs. Declared as `map[keyType]valueType`. Use `make` to initialize and `delete` to remove keys.

**Q10. How do loops work in Go?**  
Go has only `for` loops, which can be traditional (`for i:=0;i<5;i++`), while-style (`for i<5`), or range-style (`for index,value := range slice`).

**Q11. Explain switch in Go.**  
Switch statements can test values or types, support multiple cases per branch, and automatically break after each case unless `fallthrough` is used.

## 3. Structs and Interfaces

**Q12. What is a struct in Go?**  
Structs are composite types grouping fields. Example: `type Person struct { Name string; Age int }`. Structs can have methods attached.

**Q13. How do you define and implement interfaces?**  
Interfaces define a set of method signatures. Any type implementing those methods satisfies the interface implicitly. Example: `type Speaker interface { Speak() string }` and `type Dog struct{}; func (d Dog) Speak() string { return "woof" }`.

**Q14. What is an empty interface?**  
`interface{}` can hold any type, similar to `Object` in other languages. Type assertions or switches are used to extract the underlying value.

**Q15. What is interface embedding?**  
Interfaces can embed other interfaces to create composite interfaces containing multiple methods.

## 4. Pointers and Methods

**Q16. What is a pointer in Go?**  
A pointer stores the memory address of a variable. Use `&` to get the address and `*` to dereference.

**Q17. Difference between pointer and value receiver?**  
Pointer receivers modify the original struct, while value receivers operate on a copy. Use pointer receivers to avoid copying large structs or to modify state.

**Q18. How do you define methods in Go?**  
Methods are functions with a receiver. Example: `func (c *Counter) Inc() { c.n++ }` allows `Counter` to have its own methods.

**Q19. When to use pointer receivers?**  
Use pointer receivers when the method modifies the struct or when copying the struct is expensive.

**Q20. Can a value type implement an interface with pointer receiver methods?**  
No, only the pointer type implements the interface if the method has a pointer receiver.

## 5. Error Handling

**Q21. How are errors handled in Go?**  
Errors are returned as `error` type values. Functions return `error` alongside results, which must be checked by the caller.

**Q22. What is errors.Is?**  
`errors.Is(err, target)` checks if `err` matches `target` or wraps it. Useful for comparing wrapped errors.

**Q23. What is errors.As?**  
`errors.As(err, &target)` checks if any error in the chain is assignable to `target` type. Allows accessing custom error fields.

**Q24. What is errors.Join?**  
`errors.Join(err1, err2)` combines multiple errors into a single error chain.

**Q25. How to create custom errors?**  
Define a struct implementing `Error() string`. Example: `type MyErr struct { Msg string; Code int } func (e MyErr) Error() string { return e.Msg }`.

**Q26. Explain fmt.Errorf("%w", err).**  
It wraps an existing error to preserve its type for `errors.Is` and `errors.As` operations.

## 6. Concurrency

**Q27. What are goroutines?**  
Goroutines are lightweight threads managed by Go runtime, started with `go` keyword. They allow concurrent execution of functions.

**Q28. What are channels?**  
Channels are typed conduits for communication between goroutines. Unbuffered channels block until both send and receive occur. Buffered channels allow asynchronous sends up to capacity.

**Q29. Explain select statement.**  
`select` waits on multiple channel operations and executes one that is ready. Useful for timeouts and multiplexing channels.

**Q30. What are WaitGroups and Mutex?**  
`sync.WaitGroup` waits for multiple goroutines to finish. `sync.Mutex` protects shared data from concurrent access.

**Q31. How do you detect race conditions?**  
Use `go test -race` to detect unsynchronized concurrent access to shared variables.

## . Popular  Programs

##  `1. Word Count & Repeated Characters

###  Program: Word Count
```go
func wordCount(s string) map[string]int {
    words := strings.Fields(s)
    count := make(map[string]int)
    for _, word := range words {
        count[word]++
    }
    return count
}
````

### 2. Program: Repeated Characters
```go
func repeatedChars(s string) map[rune]int {
    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }
    return freq
}
```




### 3.Reverse a String
```go
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### 4. Check Palindrome
```go
func isPalindrome(s string) bool {
    return s == reverse(s)
}
```

### 5. Find Maximum Number in Array
```go
func max(arr []int) int {
    maxVal := arr[0]
    for _, v := range arr {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```

### 6.Fibonacci Series
```go
func fibonacci(n int) []int {
    seq := make([]int, n)
    seq[0], seq[1] = 0, 1
    for i := 2; i < n; i++ {
        seq[i] = seq[i-1] + seq[i-2]
    }
    return seq
}
```

### 7.Factorial
```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

### 8. Swap Two Numbers
```go
func swap(a, b int) (int, int) {
    return b, a
}
