#  Golang Error Handling, Defer, Panic & Recover 


##  1. `errors.Unwrap()` 
### **What It Does**
`errors.Unwrap()` returns the **next wrapped error** inside another error.
It helps in **unwrapping** layered or chained errors created using `fmt.Errorf("... %w", err)`.

### **How It Works**
When you use `fmt.Errorf("... %w", err)`, you create a wrapped error.
`errors.Unwrap()` allows you to access the original (inner) error inside it.

### **Example**
```go
wrappedErr := fmt.Errorf("withdraw failed: %w", ErrAtmOffline)
fmt.Println(errors.Unwrap(wrappedErr))
// Output: Currently ATM is offline
```


```go
e := errors.Unwrap(err)
fmt.Println("unwrapped error:", e)
```
If the returned error (`err`) was wrapped using `%w`, `errors.Unwrap()` would give you the original underlying error.

### **Use Case**
Useful when you have multiple layers of errors and want to identify the **root cause** of failure.

#### Example
In an ATM transaction system:
- The top-level error may say: `"Transaction Failed: ATM Offline"`
- Using `errors.Unwrap()`, you can find the **root cause** — `ErrAtmOffline` — which helps you decide whether to **retry** or **abort** the transaction.

---

##  2. `defer()` 

### **What It Does**
`defer` delays the execution of a function or statement until the **surrounding function returns**.
It is often used for **cleanup tasks** — like closing a file, unlocking a mutex, or printing logs or releasing/closing database connections .

### Example 
```go
defer fmt.Println("Defer statement :- Funds withdraw functionality completed:-")
defer deferEx()
```
These lines will execute **after** the function `TestCustomError()` finishes executing, no matter how the function exits (even if it panics).

### **Use Case**
To ensure resources are always released properly, even during errors.

#### Example
If you open a bank transaction file:
```go
f, _ := os.Open("transactions.txt")
defer f.Close()
```
The file will **always be closed**, even if a panic occurs later.

---

##  3. `panic()`

### **What It Does**
`panic()` is used to **abort the normal execution flow** when an unexpected condition occurs.
When a panic happens, Go stops executing the current function and begins **unwinding** the stack — calling all deferred functions.


```go
if amount > 10000 {
    panic("Transaction exceeds the limit")
}

if amount > balance {
    panic("Insufficient funds")
}
```
This simulates **runtime crashes** (like exceeding transaction limit or insufficient funds).  
But since you have a `recover()` inside the deferred function, your program **won’t crash completely** — it’ll gracefully recover.

### **Use Case**
Use `panic()` for **critical errors** that should not be ignored, like corrupted data or system configuration failures.

#### **Example**
ATM software panics when someone tries to withdraw more than the legal limit — preventing system abuse.

---

##  4. `recover()` 

### **What It Does**
`recover()` is used **inside a deferred function** to handle a panic.
If called during a panic, it **stops** the panic and **returns the error message** that caused it.

### **Example**
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic", r)
        fmt.Println("Transaction aborted:-")
    }
}()
```
Here, if any panic occurs inside `withdraw()`, your program will not crash, it will **recover**, print the panic reason, and continue running safely.

### **Use Case**
`recover()` is used in critical systems to **gracefully handle unexpected conditions** without bringing the entire program down.

#### **Example**
In ATM systems, if one withdrawal operation fails, the whole server should **not crash**.  
`recover()` ensures the failure is logged and the system continues serving other users.

---



```go 
var ErrAtmOffline = errors.New("Currenlty ATM is offline")
var TxnFailedError = errors.New("Transaction Failed, train again")

// custom errors
type InsufficientFundsError struct {
	Balance  float64
	Withdraw float64
	Loc      string
	ATMNo    string
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds %.2f but tried to withdraw %.2f in Location %s and ATMNo %s", e.Balance, e.Withdraw, e.Loc, e.ATMNo)
}

func withdrawFromCrypto() error {
	return errors.ErrUnsupported
}
func withdraw(balance, amount float64, loc, atmno string, atmOnline bool) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
			fmt.Println("Transaction aborted:-")
		}
	}()

	var errs []error
	if !atmOnline {
		errs = append(errs, ErrAtmOffline)
	} else if atmOnline {
		errs = append(errs, TxnFailedError)
	}
	// if amount > balance {
	// 	errs = append(errs, &InsufficientFundsError{Balance: balance, Withdraw: amount, Loc: loc, ATMNo: atmno})
	// }
	if amount > 10000 {
		panic("Transaction exceeds the limit")
	}

	if amount > balance {
		// return fmt.Errorf("withdrawl faild %w", &InsufficientFundsError{Balance: balance, Withdraw: amount, Loc: loc, ATMNo: atmno})
		panic("Insufficient funds")
	}

	if len(errs) == 0 {
		fmt.Println("Withdraw successful !")
		return nil
	}
	return errors.Join(errs...)
}

func TestErrUnsupported() {
	err := withdrawFromCrypto()
	if errors.Is(err, errors.ErrUnsupported) {
		fmt.Println("Operation is not supported")
	}
}

func TestCustomError() {

	fmt.Println("Executing TestCustomError()")

	err := withdraw(100000, 12000, "New city", "ATM10001", false)
	var fundsError *InsufficientFundsError

	e := errors.Unwrap(err)
	fmt.Println("unwrapped error:", e)
	if errors.As(err, &fundsError) {
		fmt.Printf("Custom error detected \n Balance  %.2f \n withdraw %.2f \n Location : %s \n", fundsError.Balance, fundsError.Withdraw, fundsError.Loc)
	}

	if errors.Is(err, ErrAtmOffline) {
		fmt.Println("Please try again later - the ATM is offline")
	} else if errors.Is(err, TxnFailedError) {
		fmt.Println("error occured while processing your request:", err)
		
	}
	defer fmt.Println("Defer statement :- Funds withdraw functionality completed:-")
	defer deferEx()
	fmt.Println("program execution completed:-")
}




##  Summary Table

| Function        | Purpose | When to Use | Real-World Analogy |
|-----------------|----------|--------------|---------------------|
| `errors.Unwrap()` | Extract the inner error | Debugging layered errors | Opening nested envelopes to find the root cause |
| `defer()` | Delay function execution until return | Cleanups | Always locking the door before leaving |
| `panic()` | Abort normal execution on severe errors | Irrecoverable issues | Emergency stop button |
| `recover()` | Handle panic and resume control | Graceful failure recovery | Safety net during a fall |

---

##  Exercises 

1. **Exercise 1 — Unwrap Practice**  
   - Create a custom error using `fmt.Errorf("wrapper: %w", ErrAtmOffline)`  
   - Use `errors.Unwrap()` to print the inner error.

2. **Exercise 2 — Panic & Recover**  
   - Write a function that divides two numbers.  
   - Use `panic()` if the denominator is zero, and `recover()` to handle it gracefully.

3. **Exercise 3 — Defer Order**  
   - Write three `defer` statements and print messages to see in which order they execute.

4. **Exercise 4 — Custom Error Struct**  
   - Create a struct `TransactionError` with fields `Code`, `Message`, and `ATMNo`.  
   - Return it as an error when invalid transaction occurs, and detect it using `errors.As()`.

5. **Exercise 5 — Combined Handling**  
   - Combine `defer()`, `panic()`, and `recover()` in one program to simulate an ATM withdrawal process.

---

## Key points

- Always use **`defer`** to release resources or print cleanup logs.  
- Use **`panic`** only for severe, unexpected errors.  
- Use **`recover`** to make programs robust and prevent crashes.  
- Use **`errors.Unwrap()`** to trace back the root cause of an error.  

---

