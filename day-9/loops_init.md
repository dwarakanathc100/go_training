# Loops and init()


## init() Function

The `init()` function in Go is a special function that executes **before
the `main()` function**.\
It is often used for tasks such as:

-   Initializing variables
-   Setting up configuration
-   Establishing database connections
-   Logging setup

``` go
func init() {
    println("init printing before main!")
    // database connection
    // configuration loading
    // logging setup
}
```

**Output:**

    init printing before main!

------------------------------------------------------------------------

## Sequence of Execution

Functions in Go run in the order they are called inside `main()`.

``` go
func m1() {
    println("m1")
}

func m2() {
    println("m2")
}

func main() {
    println("Hello, World!")
    m1()
    println("End")
    m2()
}
```

**Output:**

    Hello, World!
    m1
    End
    m2

------------------------------------------------------------------------

## If Statement with Short Variable Declaration

Go allows you to declare variables inside an `if` statement. The scope
of the variable is limited to the `if` block.

``` go
if marks := 90; marks >= 90 {
    println("A+")
} else {
    println("B")
}
```

**Output:**

    A+

------------------------------------------------------------------------

## Operators in Go

### Comparison Operators:

``` go
a, b := 10, 20
println(a > b)   // false
println(a < b)   // true
println(a == b)  // false
println(a != b)  // true
println(a >= b)  // false
println(a <= b)  // true
```

**Output:**

    false
    true
    false
    true
    false
    true

### Logical Operators:

``` go
a, b := 10, 20
println(a > 5 && b > 15)  // true (both conditions true)
println(a > 5 || b < 15)  // true (one condition true)
println(a != b)           // true
```

**Output:**

    true
    true
    true

------------------------------------------------------------------------

## Loops in Go

### 1. Standard For Loop

Used when the number of iterations is known.

``` go
for i := 1; i <= 5; i++ {
    println("i:", i)
}
```

**Output:**

    i: 1
    i: 2
    i: 3
    i: 4
    i: 5

------------------------------------------------------------------------

### 2. Conditional Loop

Similar to a **while loop** in other languages.

``` go
j := 1
for j <= 5 {
    println("j:", j)
    j++
}
```

**Output:**

    j: 1
    j: 2
    j: 3
    j: 4
    j: 5

------------------------------------------------------------------------

### 3. Infinite Loop with Break

Runs endlessly unless stopped using a `break` condition.

``` go
count := 0
for {
    println("infinite loop")
    count++
    if count == 5 {
        break
    }
}
```

**Output:**

    infinite loop
    infinite loop
    infinite loop
    infinite loop
    infinite loop

------------------------------------------------------------------------

### 4. Continue in Loop

Skips the current iteration and moves to the next.

``` go
for k := 1; k <= 10; k++ {
    if k%2 != 0 {
        continue
    }
    println("even number:", k)
}
```

**Output:**

    even number: 2
    even number: 4
    even number: 6
    even number: 8
    even number: 10

------------------------------------------------------------------------

## Break and Continue

-   **break**  immediately exits the loop.
-   **continue**  skips the current iteration and jumps to the next.

------------------------------------------------------------------------

# Exercises for Students

1.  Write a program using `init()` to print `"Program starting..."`
    before `main()` runs.
2.  Use an `if` statement with variable declaration to check whether a
    number is **even or odd**.
3.  Practice using comparison operators (`>`, `<`, `==`, `!=`) with
    different numbers.
4.  Write a program that prints numbers from 1 to 20 but **skips
    multiples of 3** using `continue`.
5.  Write an infinite loop that keeps running until the user enters the
    number `0` (hint: use `break`).
6.  Write a loop that prints the first 10 even numbers.

------------------------------------------------------------------------

