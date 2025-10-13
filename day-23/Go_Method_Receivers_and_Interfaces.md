# Method Receivers and Interfaces in Go


In Go, **method receivers** and **interfaces** are fundamental concepts that enable **object-oriented behaviour** such as **encapsulation**, **composition**, and **polymorphism**  without classical inheritance.

---

##  Method Receivers

### What is a Method Receiver?
A **method receiver** allows you to define methods that are associated with a specific **type (struct)**.  
In other words, methods define **behaviour** for that type.

### Syntax
```go
func (receiver ReceiverType) MethodName(args...) returnTypes {
    // method body
}
```

### Example
```go
type BankAccount struct {
    Balance int
}

func (a *BankAccount) Deposit(amount int) {
    a.Balance += amount
}

func BalanceTest() {
    b := BankAccount{Balance: 5000}
    b.Deposit(2000)
    fmt.Println(b.Balance)
}
```
**Output:**
```
7000
```

###  Explanation
- The method `Deposit` has a **pointer receiver** `*BankAccount`, meaning it modifies the actual struct value.
- When you call `b.Deposit(2000)`, Go automatically takes the address of `b` (i.e., `(&b).Deposit()`).
- You use a **value receiver** when you **don’t need to modify** the receiver.

---

### Why Do We Need Method Receivers?
- To associate behaviour (methods) with data (structs).
- To implement interfaces.
- To achieve polymorphism and clean code organization.

---

###  Real-time Use Case
A `BankAccount` struct can have methods like:
```go
Deposit(), Withdraw(), GetBalance()
```
These methods operate on the same account instance and maintain encapsulation.

---

### Value Receiver vs Pointer Receiver
| Type | Syntax | Behaviour | Use Case |
|------|---------|------------|----------|
| Value Receiver | `(a Account)` | Works on a copy | When no state modification needed |
| Pointer Receiver | `(a *Account)` | Works on the original value | When updating receiver’s data |

---

##  Interfaces in Go

### Definition
An **interface** in Go defines a **set of method signatures** (a contract).  
Any type that implements **all methods** of the interface **implicitly satisfies** that interface.


Think of it as a contract:

“If your type implements all the methods listed in this interface, then you satisfy it.”

Go interfaces are implicit — you don’t need to declare implements InterfaceName.

### Syntax
```go
type InterfaceName interface {
    Method1(param Type) ReturnType
    Method2(param Type) ReturnType
}
```

---

## Use Case  Payment Gateway
Imagine a **payment system** where multiple gateways (PayPal, GPay, PhonePe) exist.  
Each must provide the same functions: `Pay()` and `Refund()`.

### Interface Declaration
```go
type PaymentGateway interface {
    Pay(amount float64) string
    Refund(txnId string) string
}
```

### Implementation
```go
type Paypal struct{}
func (p Paypal) Pay(amount float64) string   { return "paypal implementation" }
func (p Paypal) Refund(txnId string) string  { return "paypal refund" }

type PhonePe struct{}
func (p PhonePe) Pay(amount float64) string  { return "phonepe implementation" }
func (p PhonePe) Refund(txnId string) string { return "phonepe refund" }

type Gpay struct{}
func (g Gpay) Pay(amount float64) string     { return "gpay implementation" }
func (g Gpay) Refund(txnId string) string    { return "gpay refund" }
```

### Usage
```go
func PaymentInitiated(paymentGateway PaymentGateway, amount float64) string {
    return paymentGateway.Pay(amount)
}

func PaymentRefund(paymentGateway PaymentGateway, txnId string) string {
    return paymentGateway.Refund(txnId)
}

func Payment() {
    paypal := Paypal{}
    phonePe := PhonePe{}
    gpay := Gpay{}

    s := PaymentInitiated(paypal, 1000.00)
    s1 := PaymentInitiated(phonePe, 1000.00)
    s2 := PaymentInitiated(gpay, 1000.00)

    fmt.Println(s)
    fmt.Println(s1)
    fmt.Println(s2)
}
```
**Output:**
```
paypal implementation
phonepe implementation
gpay implementation
```

---

## Polymorphism in Go
Polymorphism allows different types to **respond to the same interface** in their own way.

- `PaymentGateway` acts as an abstraction.
- All payment types (`Paypal`, `Gpay`, `PhonePe`) implement the same methods differently.
- The same function `PaymentInitiated()` works with any gateway.

---

## Composition in Go
Go promotes **composition over inheritance**.  
A struct can **embed** another struct or interface to reuse behaviour.
We’ll see how Go allows one struct to “reuse” functionality of another by embedding it — i.e., composition.
This is different from inheritance (which Go doesn’t have). Instead, Go promotes composition over inheritance

## Real-World Scenario
Imagine we are building a Payment Gateway System that processes payments via different methods — like Credit Card and PayPal.
We also want to log every transaction.

Rather than having every payment type implement its own Log() method separately,
we can create a reusable Logger struct and embed it into multiple payment types.

```package main

import "fmt"

// Logger struct — reusable component
type Logger struct{}

func (Logger) Log(message string) {
	fmt.Println("[LOG]:", message)
}

// PaymentProcessor interface — defines behavior for payment processing
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// CreditCardPayment struct embeds Logger — composition in action
type CreditCardPayment struct {
	Logger
	CardNumber string
}

func (c CreditCardPayment) ProcessPayment(amount float64) {
	c.Log(fmt.Sprintf("Processing Credit Card payment of $%.2f for card %s", amount, c.CardNumber))
	fmt.Println("Credit Card Payment Successful!")
}

// PayPalPayment struct also embeds Logger
type PayPalPayment struct {
	Logger
	Email string
}

func (p PayPalPayment) ProcessPayment(amount float64) {
	p.Log(fmt.Sprintf("Processing PayPal payment of $%.2f for account %s", amount, p.Email))
	fmt.Println("PayPal Payment Successful!")
}

// Function to simulate any type of payment
func MakePayment(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount)
}

func main() {
	cc := CreditCardPayment{CardNumber: "1234-5678-9999-0000"}
	pp := PayPalPayment{Email: "user@example.com"}

	// Both types automatically have Log() due to composition
	MakePayment(cc, 150.75)
	MakePayment(pp, 89.50)
}

```

---

## Empty Interface (`interface{}`)
The **empty interface** represents **zero methods**, so **every type** satisfies it.

### Use Case
Used when a function should accept values of **any type**.

### Example
```go
func Describe(i interface{}) {
    fmt.Printf("Type = %T, Value = %v\n", i, i)
}

func ExampleEmptyInterface() {
    Describe(42)
    Describe("hello")
    Describe(true)
}
```
**Output:**
```
Type = int, Value = 42
Type = string, Value = hello
Type = bool, Value = true
```

---

## Ways to Declare and Define Interfaces in Go

### 1. **Basic Interface**
A **basic interface** contains one or more method declarations without any implementation.  
Any type that defines all these methods **automatically implements** this interface — **no explicit keyword** like `implements` or `extends` is required.

### Example
```go
type Reader interface {
    Read() string
}

type File struct{}

func (f File) Read() string {
    return "Reading data from file..."
}

func main() {
    var r Reader
    r = File{}
    fmt.Println(r.Read())
}
```
### Explanation
- `Reader` is an interface that expects a `Read()` method returning a string.
- `File` struct implements this interface by defining the `Read()` method.
- Since Go uses **structural typing**, there’s no need for explicit implementation.

---

## 2.  Interface with Multiple Methods

### Definition
Interfaces can define **multiple methods**, and any type implementing all those methods automatically satisfies the interface.

### Example
```go
type Device interface {
    Start() error
    Stop() error
}

type Fan struct{}

func (f Fan) Start() error {
    fmt.Println("Fan started")
    return nil
}

func (f Fan) Stop() error {
    fmt.Println("Fan stopped")
    return nil
}
```
### Explanation
- `Device` defines two behaviours — `Start()` and `Stop()`.
- The `Fan` struct provides concrete implementations for both.
- You can now create functions that accept `Device` interface and work with any type that satisfies it — achieving **polymorphism**.

---

## 3.  Embedded Interface (Interface Composition)

### Definition
Go allows **interface composition**, meaning you can **combine multiple interfaces** into one larger interface.  
This is Go’s version of **interface inheritance**, promoting **composition over inheritance**.

### Example
```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

type ReadWriter interface {
    Reader
    Writer
}
```
### Explanation
- `ReadWriter` is a **composed interface** that embeds both `Reader` and `Writer`.
- Any type that implements both `Read()` and `Write()` methods satisfies `ReadWriter`.
- This design allows flexibility and modularity when building large systems.

### Real-world Analogy
Think of `Reader` as a “download” capability and `Writer` as an “upload” capability.  
A `ReadWriter` like a network connection can **both download and upload**.

---

## 4.  Empty Interface (`interface{}`) (`we will discuss in next class`)


The **empty interface** does not have any methods.  
Because it has **no requirements**, **every type** in Go satisfies it automatically.

### Example
```go
type Any interface{}

func Describe(i Any) {
    fmt.Printf("Type = %T, Value = %v\n", i, i)
}

func main() {
    Describe(42)
    Describe("hello")
    Describe(true)
}
```
### Explanation
- `Any` is equivalent to `interface{}` and can hold **any type of value** — similar to `Object` in Java or `any` in TypeScript.
- Useful in cases like logging, collections (`[]interface{}`), or functions accepting **dynamic types**.

### Real-world Use Case
Used in libraries like `fmt`, `json.Marshal`, and other generic utilities where **type flexibility** is needed.

---

## 5. Interface with Method Receivers


This approach shows how **structs implement interfaces** by providing **methods with receivers** (either value or pointer).

### Example
```go
type Worker interface {
    Work(hours int) string
}

type Employee struct{}

func (e Employee) Work(hours int) string {
    return fmt.Sprintf("Worked %d hours", hours)
}

func main() {
    var w Worker
    w = Employee{}
    fmt.Println(w.Work(8))
}
```
### Explanation
- The `Worker` interface expects a `Work(int) string` method.
- The `Employee` struct implements it by defining a method with a **value receiver**.
- When `w.Work(8)` is called, it runs the `Employee` implementation.
- If the receiver were a pointer (e.g., `func (e *Employee) Work()`), only a pointer to `Employee` would satisfy the interface.

### Real-world Example
Used when defining behaviours like `Run()`, `Execute()`, or `Perform()` in systems involving jobs, workers, or tasks.

---

## 6. Interface Variable Assignment


You can **assign** a concrete type to an **interface variable** as long as the type implements all the interface’s methods.  
Once assigned, you can call interface methods polymorphically.

### Example
```go
type PaymentGateway interface {
    Pay(amount float64) string
    Refund(txnId string) string
}

type Paypal struct{}

func (p Paypal) Pay(amount float64) string {
    return "Paypal payment done"
}

func (p Paypal) Refund(txnId string) string {
    return "Paypal refund successful"
}

func main() {
    var p PaymentGateway
    p = Paypal{} // assigning concrete type to interface
    fmt.Println(p.Pay(1000.0))
    fmt.Println(p.Refund("TXN123"))
}
```
### Explanation
- `Paypal` implements `PaymentGateway` by providing both `Pay` and `Refund` methods.
- The variable `p` of type `PaymentGateway` can now **store any type** that satisfies this interface.
- This enables **runtime polymorphism**, where different payment gateways (`GPay`, `PhonePe`, etc.) can be used interchangeably.

---


**By combining method receivers and interfaces**, Go enables developers to design flexible, modular, and testable code without traditional inheritance.

##  7. Exercises 

1. **Bank Application:**  
   - Create a struct `BankAccount` with methods `Deposit`, `Withdraw`, and `CheckBalance`.  
   - Use **pointer receivers** where necessary.

2. **Employee Management:**  
   - Define an interface `Employee` with methods `GetSalary()` and `GetDetails()`.  
   - Implement it for `FullTimeEmployee` and `ContractEmployee`.

3. **Shape Interface:**  
   - Create an interface `Shape` with methods `Area()` and `Perimeter()`.  
   - Implement it for `Rectangle` and `Circle`.  
   - Demonstrate **polymorphism**.

4. **Empty Interface Challenge:**  
   - Write a function that accepts `interface{}` and prints type and value dynamically.(we will discuss in next class)
