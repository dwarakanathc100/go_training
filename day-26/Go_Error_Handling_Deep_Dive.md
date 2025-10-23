#  Go Error Handling Deep Dive (`errors.Is`, `errors.As`, `errors.Join`, `errors.ErrUnsupported`)

We'll explore the following built-in functions with **real-life ATM
withdrawal and login examples**:

-   `errors.Is()`
-   `errors.As()`
-   `errors.Join()`
-   `errors.ErrUnsupported`

------------------------------------------------------------------------

##  Introduction

Go provides a simple yet powerful way to handle errors.
Traditionally, you'd return an error from a function and check if it's
`nil`.
However, with Go 1.13 and later, the `errors` package added **richer
error inspection utilities** to simplify identifying, combining, and
unwrapping errors.

------------------------------------------------------------------------

## 1. `errors.Is()`

### syntax:

`errors.Is(err, target error) bool`

Checks whether an error is **equal to** or **wraps** another specific
error.
It's like saying, "Does this error or anything inside it match this
known error?"

------------------------------------------------------------------------

###  Example 

``` go
var ErrAtmOffline = errors.New("Currenlty ATM is offline")
var TxnFailedError = errors.New("Transaction Failed, train again")
```

Example

``` go
if errors.Is(err, ErrAtmOffline) {
		fmt.Println("Please try again later - the ATM is offline")
	}
```

### Explanation

-   `errors.Is()` compares the returned error 
    `errors.ErrAtmOffline`.
-   If they match, it prints **"OPlease try again later - the ATM is offline"**.
-   Useful for detecting **specific, known conditions** like network
    timeouts, unsupported operations, etc.

###  Output

    Please try again later - the ATM is offline

###  Real-world Use Case

-   When a payment gateway doesn't support a certain transaction type.
-   When a database driver doesn't support a particular SQL command.

------------------------------------------------------------------------

## 2. `errors.As()`

### syntax:

`errors.As(err, &target) bool`

This function checks whether an error or any error in its chain
**matches a specific type** and, if it does, **assigns it** to the
target variable.
In other words, `errors.As()` helps you **extract custom error data**.

------------------------------------------------------------------------

###  Example 

``` go
func TestCustomError() {
    err := withdraw(1000, 2000, "New city", "ATM10001", false)

    var fundsError *InsufficientFundsError
    if errors.As(err, &fundsError) {
        fmt.Printf("Custom error detected \n Balance %.2f \n withdraw %.2f \n Location : %s \n",
            fundsError.Balance, fundsError.Withdraw, fundsError.Loc)
    }
}
```



-   The function `withdraw()` may return a **custom struct-based error**
    like this:

    ``` go
    &InsufficientFundsError{
        Balance: 1000,
        Withdraw: 2000,
        Loc: "New city",
        ATMNo: "ATM10001",
    }
    ```

-   Using `errors.As()`, we can **extract** this error into a variable
    of type `*InsufficientFundsError`.

-   Once extracted, we can access its internal fields such as `Balance`,
    `Withdraw`, and `Loc`.

###  Output

    Custom error detected
    Balance 1000.00
    Withdraw 2000.00
    Location : New city

### Real-world Use Case

-   Extracting structured error data like API error codes or validation
    failures.
-   For example, `json.Unmarshal()` returns `*json.SyntaxError`, which
    you can detect using `errors.As()`.

------------------------------------------------------------------------

## 3. `errors.Join()`

### syntax:

`errors.Join(errs ...error) error`

Combines multiple errors into a **single error object**.
The resulting error still lets you use `errors.Is()` or `errors.As()` to
check for specific causes.

------------------------------------------------------------------------

###  Example 

``` go
func withdraw(balance, amount float64, loc, atmno string, atmOnline bool) error {
    var errs []error

    if !atmOnline {
        errs = append(errs, ErrAtmOffline)
    } else {
        errs = append(errs, TxnFailedError)
    }

    if amount > balance {
        errs = append(errs, &InsufficientFundsError{
            Balance:  balance,
            Withdraw: amount,
            Loc:      loc,
            ATMNo:    atmno,
        })
    }

    if len(errs) == 0 {
        fmt.Println("Withdraw successful !")
        return nil
    }

    return errors.Join(errs...)
}
```



-   We may face **multiple issues** in one withdrawal attempt:
    -   ATM is offline
    -   Transaction failed
    -   Insufficient funds
-   Instead of returning only one, `errors.Join()` **combines all** into
    one error chain.

Even though it returns a single error, you can still check specific ones
later:

``` go
if errors.Is(err, ErrAtmOffline) { ... }
if errors.As(err, &fundsError) { ... }
```

###  Output Example

    Custom error detected
    Balance 1000.00
    Withdraw 2000.00
    Location : New city
    Please try again later - the ATM is offline

###  Real-world Use Case

-   Handling multiple API failures together (e.g., logging, database,
    and cache errors).
-   Aggregating multiple validation errors before sending feedback to
    the user.

------------------------------------------------------------------------

## 4. Constant: `errors.ErrUnsupported`



`errors.ErrUnsupported` is a **predefined sentinel error** in Go's
`errors` package.

###  Purpose

Used to indicate that an operation or method is **not supported** by the
current implementation.

------------------------------------------------------------------------

###  Example 

``` go
func withdrawFromCrypto() error {
    return errors.ErrUnsupported
}
```



-   This simulates a scenario where cryptocurrency withdrawal is **not
    yet implemented**.
-   When this function is called, it returns the built-in
    `ErrUnsupported` error.
-   The test function (`TestErrUnsupported`) detects it using
    `errors.Is()`.

###  Output

    Operation is not supported

###  Real-world Use Case

-   File systems that don't support certain operations (e.g., deleting
    read-only files).
-   Network clients where certain HTTP verbs like PATCH or OPTIONS are
    unsupported.

------------------------------------------------------------------------

##  Summary Table

  -----------------------------------------------------------------------------
  Function / Constant       Description    Use Case          Example Output
  ------------------------- -------------- ----------------- ------------------
  `errors.Is()`             Checks if an   Detecting known   "Operation is not
                            error matches  errors            supported"
                            or wraps                         
                            another                          

  `errors.As()`             Extracts a     Accessing         Prints custom
                            custom error   structured error  struct values
                            type for       data              
                            detailed info                    

  `errors.Join()`           Combines       Handling multiple Prints multiple
                            multiple       simultaneous      error messages
                            errors into    failures          
                            one                              

  `errors.ErrUnsupported`   Standard Go    Indicating        "Operation is not
                            error for      non-implemented   supported"
                            unsupported    features          
                            actions                          
  -----------------------------------------------------------------------------

------------------------------------------------------------------------

##  Practice Exercises 

1.  Create a new custom error type `ExceededDailyLimitError` with fields
    `Limit` and `Attempted`.
2.  Modify the `withdraw()` function to include this error when the
    withdrawal exceeds a daily limit.
3.  Combine it with other errors using `errors.Join()`.
4.  Write a new function `TestExceededLimit()` to:
    -   Use `errors.As()` to detect the custom error.
    -   Print its fields individually.
    -   Use `errors.Is()` to detect if `ErrAtmOffline` occurred as well.

------------------------------------------------------------------------


