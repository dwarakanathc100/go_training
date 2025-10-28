
# Go test case/test scenarios example

```go
func withdraw(balance, amount float64, loc, atmno string, atmOnline bool) any {

	if !atmOnline {
		return ErrAtmOffline.Error()
	}

	if amount <= 0 {
		return errors.New("Invalid amount").Error()
	}
	if amount > balance {
		return errors.New("Insufficient funds").Error()
	}
	return "Withdraw successful"
}
```



---

##  Test scenarios
```go
package errors

import (
    "fmt"
    "testing"
)

func TestWithdraw(t *testing.T) {
    tests := []struct {
        name       string
        balance    float64
        amount     float64
        loc, atmno string
        atmOnline  bool
        expected   string
    }{
        {"Successful withdrawal", 1000, 500, "Loc1", "ATM100", true, "Withdraw successful"},
        {"withdrawal with exact amount", 1000, 1000, "Loc2", "ATM12", true, "Withdraw successful"},
        {"ATM is offline", 1000, 100, "Loc2", "ATM12", false, "Currently ATM is offline"},
        {"Insufficient funds case", 500, 1000, "Loc2", "ATM12", true, "Insufficient funds"},
        {"Withdraw zero amount", 1000, 0, "Loc2", "ATM12", true, "Invalid amount"},
        {"Withdraw negative amount", 1000, -50, "Loc2", "ATM12", true, "Invalid amount"},
    }

    for _, tt := range tests {
        tt := tt

        t.Run(tt.name, func(t *testing.T) {
            result := withdraw(tt.balance, tt.amount, tt.loc, tt.atmno, tt.atmOnline)
            got := fmt.Sprint(result)
            if got != tt.expected {
                t.Fatalf("expected %q, got %q", tt.expected, got)
            }
        })
    }
}
```

- This defines a _table_ (slice) of test cases. Each row contains inputs and the `expected` output.
- Using a table lets you add many cases cleanly and iterate over them in the same test loop.




## `t.Run()` 

- `t.Run(name string, f func(t *testing.T))` runs the given function `f` as a *subtest* named `name`.
- Subtests are useful to see which specific row failed in the table output, the test runner prints each failing subtest name.
- Subtests can be nested and reported individually, and they support `go test -run` filtering by name.
- When using `t.Parallel()` inside a subtest, tests run concurrently 


---



## Exercises
1. Clone repository :https://github.com/dwarakanathc100/banking-system
2. Understand all packages
3. Write test cases for loan, emi packages (add few more functions)
4. Go through the reression and feature packages