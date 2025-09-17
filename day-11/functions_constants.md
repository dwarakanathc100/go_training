**Continuation from yesterday’s clas...**
### e. No Return Type (Void-like)

Functions that perform an action but do not return any
value.

Use case: Logging, printing, modifying variables.

Example:

``` go
func print(a, b int) {
    fmt.Println("sum is ", a+b)
}
```

Output:

    sum is 3

------------------------------------------------------------------------

### f. Multiple Return Values

Functions in Go can return multiple values. This is very
common in Go, especially for returning both result and error.

Use case: Returning quotient and remainder in division.

Example:

``` go
func divide(a, b int) (int, int, int) {
    quotient := a / b
    remainder := a % b
    c := a + b
    return quotient, remainder, c
}
```

Output:

    Quotient: 2 Remainder: 1 c: 10

------------------------------------------------------------------------

### g. Named Return Values

You can name return values in the function signature. The
function can then return without explicitly listing values.

Use case: Improves readability.

Example:

``` go
func namedReturn(a, b int) (q, r int) {
    q = a / b
    r = a % b
    return
}
```

Output:

    Quotient: 2 Remainder: 1

------------------------------------------------------------------------

### h. Anonymous Functions (Function Literals)

Functions without a name. They can be assigned to variables.

Use case: Short, temporary functionality.

Example:

``` go
var anon = func(x, y int) int {
    return x + y
}
fmt.Println("Anonymous function result:", anon(3, 4))
```

Output:

    Anonymous function result: 7

------------------------------------------------------------------------

### i. Immediate Invocation (IIFE)

Anonymous functions can be invoked immediately.

Use case: Running logic instantly without storing it.

Example:

``` go
result := func(x, y int) int {
    return x * y
}(30, 40)

fmt.Println("IIFE result:", result)
```

Output:

    IIFE result: 1200

------------------------------------------------------------------------

## 2. Default Values in Go

In Go, variables have default "zero values" if not initialized.

Example:

``` go
func defaultValues() {
    var i int
    var f float64
    var b bool
    var s string
    var p *int
    var arr []int
    var m map[string]int
    var ch chan int
    var inter interface{}

    fmt.Println("int:", i)
    fmt.Println("float64:", f)
    fmt.Println("bool:", b)
    fmt.Println("string:", s)
    fmt.Println("pointer:", p)
    fmt.Println("array:", arr)
    fmt.Println("map:", m)
    fmt.Println("channel:", ch)
    fmt.Println("interface:", inter)
}
```

Output:

    int: 0
    float64: 0
    bool: false
    string:
    pointer: <nil>
    array: []
    map: map[]
    channel: <nil>
    interface: <nil>

------------------------------------------------------------------------

## 3. Random Integers

Go provides `math/rand` package to generate random numbers.

Example:

``` go
func generateRandomNumber() {
    for i := 0; i < 5; i++ {
        fmt.Println(rand.Intn(100))
    }
}
```

Output (example, changes every run):

    81
    12
    45
    3
    98

------------------------------------------------------------------------

## 4. Constants in Go

Constants are immutable values. They can be declared with or without
type.

### a. With Type and Without Type

``` go
const x = 10     // untyped constant
const y int = 20 // typed constant
```

------------------------------------------------------------------------

### b. Multiple Constants

``` go
const width, height int = 100, 50
const area = width * height
```

------------------------------------------------------------------------

### c. Using `iota`

`iota` is used to create enumerated constants.

Example:

``` go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday = 20
    Thursday  = iota
    Friday
    Saturday
)
```

Output:

    Days: 0 1 2 20 4 5 6

Another example:

``` go
const (
    A = iota
    B
    C = 10
)
```

Output:

    0 1 10

------------------------------------------------------------------------

## Exercises

1.  Write a function that calculates the factorial of a number using
    recursion.
2.  Modify the `divide` function to also return a boolean indicating
    whether division was exact (no remainder).
3.  Write an anonymous function that squares a number and immediately
    invoke it.
4.  Experiment with `iota` to Read, Write, Execute constants for permissions
    

