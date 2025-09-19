# Debugging in VSCode and Taking User Input in Go

## What is Debugging?

Debugging is the process of identifying, analyzing, and fixing bugs
(errors) in a program.\
It helps developers understand why a program is not working as expected
and corrects the issues to ensure the program runs smoothly.

### Need for Debugging

-   To identify logical, runtime, or syntax errors.
-   To ensure the application runs as expected without crashing.
-   To improve the performance and reliability of applications.

### Where to Use Debugging?

-   During software development when unexpected behavior occurs.
-   While testing features before releasing them to production.
-   When fixing reported issues in existing applications.

### Why to Use Debugging?

-   To save time by quickly identifying the exact source of the problem.
-   To understand the program flow better.
-   To maintain clean and error-free code.

### Steps to Debug in VSCode

1.  Open your project in VSCode.
2.  Go to the file where you want to debug.
3.  Set a **breakpoint** by clicking next to the line numbers.
4.  Run the program in **Debug mode** (press `F5` or select *Run and
    Debug* from the left menu).
5.  VSCode will pause execution at the breakpoint.
6.  Use the **Debug toolbar** to step through:
    -   *Step Over* → execute line by line without going into functions.
    -   *Step Into* → enter inside function calls.
    -   *Step Out* → move out of the current function.
    -   *Continue* → run the program until the next breakpoint.
7.  Observe variable values in the **Debug Panel**.

------------------------------------------------------------------------

## Taking Inputs from User in Go

### Why Do We Need Inputs from Users?

User inputs are essential for making applications interactive and
dynamic. Examples include: - **ATM Machine**: User enters a PIN and the
amount to withdraw. - **Hotel Billing System**: User provides details
like food items and quantities to generate a bill. - **Command-Line
Applications**: Inputs are typed directly in the terminal and processed
by the backend language (Go, Java, Python, etc.). - **Web
Applications**: Inputs are provided via forms (React.js, Angular, etc.),
then sent to backend applications through APIs.

### Types of Applications

1.  **Command-Line Applications (CLI)**
    -   Input is entered via terminal/console.
    -   Example: `go run main.go` then entering values in console.
    -   Input is captured directly by backend programming languages like
        Go, Java, Python.
2.  **Web Applications**
    -   Input is provided through UI forms or fields in the browser.
    -   Frontend frameworks like **React.js** or **Angular** handle user
        interactions.
    -   Data is then sent to backend applications (Go, Java, Node.js,
        etc.) via APIs.

------------------------------------------------------------------------

## User Input in Go

Go provides several functions for reading input from users. The most
common ones are:

### 1. `fmt.Scan()`

-   Reads **space-separated values** from standard input.
-   Stops reading when it encounters a **space or newline**.
-   If fewer values are entered than expected, the program waits for
    more input.
-   If more values are entered, the extra ones remain in the input
    buffer.

**Example:**

``` go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age)
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}
```

------------------------------------------------------------------------

### 2. `fmt.Scanln()`

-   Similar to `fmt.Scan()`, but it **stops reading at a newline**.
-   It is useful when taking input in one line and then pressing enter.
-   If there are extra values beyond what variables expect, the program
    throws an error.

**Example:**

``` go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanln(&name, &age)
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}
```

------------------------------------------------------------------------

### 3. `fmt.Scanf()`

-   Provides **formatted input** using format specifiers (like `printf`
    but for reading).
-   Useful when you want inputs in a specific format.
-   `%s` → string, `%d` → integer, `%f` → float, etc.

**Example:**

``` go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age (format: Name Age): ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}
```

------------------------------------------------------------------------

## Scenarios for User Inputs

-   **Banking Systems (ATM, NetBanking)** → Enter PIN, amount, account
    number.
-   **E-commerce Websites** → Enter item details, address, and payment
    info.
-   **Billing Systems** → Input product name, quantity, and payment
    mode.
-   **Reservation Systems** → Input name, date, and seat selection.
-   **Survey/Feedback Forms** → Collect user opinions.
-   **Educational Apps** → Take inputs like name, answers, and roll
    numbers.

------------------------------------------------------------------------

## Exercises

### Debugging Exercise

1.  Write a nested loop program that prints values of `i` and `j`.
    -   Run the program and observe the output.
    -   Use VSCode Debugger to step through and check how `i` and `j`
        change.

**Example:**

``` go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 2; j++ {
            fmt.Printf("i = %d, j = %d\n", i, j)
        }
    }
}
```

### Function Call Exercise

2.  Write a function `add(a, b int) int` that returns the sum of two
    numbers.
    -   Call this function from `main` and print the result.
    -   Debug using VSCode to step into the function and see how
        parameters are passed.

**Example:**

``` go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(2, 3)
    fmt.Printf("Sum is: %d\n", result)
}
```

### Input Practice Exercise

3.  Write a program to take `name` and `age` from the user using:
    -   `fmt.Scan()`
    -   `fmt.Scanln()`
    -   `fmt.Scanf()`\
        Compare how each behaves with different inputs.

------------------------------------------------------------------------


