# Golang Error Handling 

## 1. What is an Error?
In Go, an **error** represents a problem that occurs during the execution of a program for example, a failed login, division by zero, or missing file.  
Go treats errors as **values** rather than exceptions, which means you can **return and handle them explicitly** instead of letting the program crash.

---

## 2. Why Do We Need Error Handling?
Error handling in Go is essential because it:
- Prevents program crashes.
- Helps locate issues with meaningful messages.
- Ensures smooth recovery from unexpected situations.
- Improves reliability, maintainability, and debugging.

---

## 3. Where to Use Error Handling
You should handle errors anywhere the program might fail:
- Opening or reading a file  
- Connecting to a database  
- Making an HTTP request  
- Parsing user input  
- Performing mathematical operations (like division)  

---

## 4. How Go Handles Errors
Go doesn’t have exceptions. Instead, functions **return error values**.  
A function that can fail usually returns two values  the expected result and an error object:

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

This style forces developers to **handle errors immediately**, leading to cleaner and more reliable programs.

---

## 5. Ways to Declare or Create Errors in Go

### (1) `errors.New()`
The simplest and most common way to create a basic error message.

- Defined in the `errors` package.
- Returns an `error` object with a static message.
- Best used for simple, constant error messages (no formatting required).

#### Example
```go
package main

import (
    "errors"
    "fmt"
)

func checkUser(username string) error {
    if username == "" {
        // Creates a simple static error
        return errors.New("username cannot be empty")
    }
    return nil
}

func main() {
    err := checkUser("")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

####  When to Use
Use `errors.New()` when:
- You just need a fixed error message.
- No dynamic or contextual information is required.
  
Example: `"file not found"`, `"connection failed"`, `"invalid input"`

---

### (2) `fmt.Errorf()`
A more advanced and flexible way to create errors using formatted strings.


- Defined in the `fmt` package.
- Allows string formatting using verbs like `%s`, `%d`, etc.
- Can include **contextual or dynamic data** in the message.
- Supports error wrapping using `%w`.

####  Example (Simple Formatting)
```go
func login(username, password string) error {
    if username != "admin" || password != "1234" {
        // Create formatted error message
        return fmt.Errorf("invalid credentials for user %s", username)
    }
    return nil
}

func main() {
    err := login("guest", "wrong")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```


####  When to Use
Use `fmt.Errorf()` when:
- You want to include variables or context in your error message.
- You need to wrap another error.
- You want meaningful, traceable logs.

---

### (3) **Custom Error Types**
Go allows you to create **structured custom errors** by defining a type that implements the `Error()` method of error interface.


- Used for domain-specific, detailed error messages.
- You can include **fields** like error code, timestamp, or any contextual info.
- Makes your code more readable and testable.

####  Example
```go
package main

import "fmt"

// Define custom error struct
type InsufficientFundsError struct {
    Balance  float64
    Withdraw float64
    Location string
    ATMNo    string
}

// Implement Error() method for the struct
func (e *InsufficientFundsError) Error() string {
    return fmt.Sprintf(
        "Insufficient funds: balance %.2f, attempted %.2f at %s (ATM: %s)",
        e.Balance, e.Withdraw, e.Location, e.ATMNo)
}

// Function using custom error
func withdraw(balance, amount float64, loc, atm string) error {
    if amount > balance {
        return &InsufficientFundsError{
            Balance:  balance,
            Withdraw: amount,
            Location: loc,
            ATMNo:    atm,
        }
    }
    fmt.Println("Withdrawal successful!")
    return nil
}

func main() {
    err := withdraw(1000, 2000, "London", "ATM007")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

####  When to Use
Use custom errors when:
- You need to represent complex business logic.
- You want to include structured data in your error.
- You need type assertion or custom error handling logic.

####  Example Use Case
In a **banking system**, you might return different types of errors like:
- `InsufficientFundsError`
- `AccountLockedError`
- `InvalidPINError`

You can then use **type assertion** to handle each differently:
```go
if err, ok := err.(*InsufficientFundsError); ok {
    fmt.Println("Please check your balance:", err)
}
```

---

## 6. Summary Comparison

| Type | Function | Description | Best For |
|------|-----------|--------------|-----------|
| `errors.New()` | Creates static errors | Simple, fixed errors | Simple validation |
| `fmt.Errorf()` | Creates formatted errors, supports wrapping | Dynamic messages, context-rich errors | Most common use |
| Custom Error | Struct implementing `Error()` | Rich, structured, domain-specific errors | Enterprise applications |

## 7 : Exercises 
### 1. Basic Validation Error

Create a function validateAge(age int) error that:

Returns an error using `errors.New()` if age is less than 18.
Otherwise, prints "Access granted!".

### 2. Dynamic Error Message

Create a function `getEmployee(id int)` error that:

Returns an error using `fmt.Errorf("employee with ID %d not found", id) if id is not in the range 1–5.`

Prints "Employee found!" otherwise.


### 3. Custom Error Assertion
Create multiple custom errors (e.g. `InvalidPINError`, `CardExpiredError`) and handle each using **type assertion**.

**Expected Output:**
```
Error: CardExpiredError - Please renew your card before use
