
# Golang Testing 

## 1. Introduction to Testing

Testing is a critical phase in software development that ensures code quality, correctness, and reliability. It helps verify that your code behaves as expected and reduces the risk of defects being introduced into production.

In large systems like banking software, multiple teams work on different modules such as:

```go
//Bank ---> 
// finacle (core banking software used by many banks)

// Core banking modules (developed and tested by different teams)
   user account maintenance     [handled by one team] [stored in one repo]
   money deposit module         [handled by another team] [different repo]
   money withdrawal module
   statement generation module
   email/SMS notification module
   loan processing department
   EMI (loan repayment) department

```

In such systems, testing ensures that each module and their integrations function correctly without breaking other parts of the system.

---

## 2. Why Do We Need Testing?

- Detect bugs early – before code reaches production.  
- Ensure stability – verify that new changes don’t break existing functionality.  
- Improve confidence – makes developers confident about code quality.  
- Enable refactoring – safely modify code with assurance that tests will catch regressions.  
- Reduce maintenance cost – fixing bugs early is much cheaper than fixing them in production.  

---

## 3. Types of Software Testing

| Type | Description | Example |
|------|--------------|----------|
| Unit Testing | Tests individual functions or small components in isolation. | `Add(a, b int)` function |
| Integration Testing | Tests how multiple modules interact with each other. | Testing `withdraw` + `statement generation` |
| System Testing (E2E) | Tests the entire application workflow from start to finish. | Testing bank application end-to-end scenario |
| Acceptance Testing | Ensures software meets business requirements. | Checking if banking transactions are correct per policy |
| Regression Testing | Ensures new updates don’t break existing features. | Running test suite after each release |
| Performance Testing | Checks how fast or scalable a system is. | Load testing a transaction API |

---

## 4. Go’s Built-in Testing Framework

Go has a built-in testing package called `testing`, which makes it easy to write and run unit tests.

### Key Features

- No external dependencies (comes with Go).
- Supports unit tests, benchmark tests, and example tests.
- Integrated with `go test` command.
- Provides test coverage, parallel testing, and subtests.

### How It Works Internally

- Every test file must end with `_test.go` (e.g., `add_test.go`).
- Each test function must start with `Test` and take a parameter `t *testing.T`.
- The `go test` command automatically discovers and runs all tests.
- Functions like `t.Error()`, `t.Errorf()`, `t.FailNow()` are used to report failures.

### Common Functions

| Function | Description |
|-----------|--------------|
| `t.Log()` | Logs information about the test. |
| `t.Error()` | Logs error and continues execution. |
| `t.Errorf()` | Logs formatted error message and continues. |
| `t.Fatal()` | Logs message and stops execution of that test. |
| `t.Run()` | Runs a subtest. |
| `t.Parallel()` | Runs test in parallel. |

---

## 5. Example: Unit Testing the Add Function

### add.go
```go
package main

func Add(a, b int) int {
	return a + b
}
```

### Single Test Case
```go
package main

import "testing"

func TestAddSingle(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("expected %d but got %d", expected, result)
	}
}
```

### Multiple Test Cases (Table-Driven Tests)
```go
package main

import "testing"

func TestAdd(t *testing.T) {

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -1, -3},
		{"zero numbers", 0, 0, 0},
		{"positive negative combination numbers", -2, -1, -3},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add scenario: %s (%d, %d): expected %d, got %d", tt.name, tt.a, tt.b, tt.expected, result)
		}
	}
}
```

---



---

## 6. Advantages of Testing

- Enhances software quality and reliability.  
- Builds developer confidence.  
- Helps automate regression checks.  
- Improves maintainability and documentation.  
- Reduces production failures.  

---

## 7. Exercises 

1. Write a function `Subtract(a, b int) int` and create test cases for it.  


