# Day 4 Notes

## 1. Require at Least One Review Before Merging a Pull Request

To enforce code quality and collaboration, you can configure GitHub to require at least one review approval before merging a PR.

### Steps:
1. Go to your repository in GitHub.
2. Navigate to **Settings** → **Branches**.
3. Under **Branch protection rules**, click **Add rule**.
4. In the **Branch name pattern**, type the branch you want to protect (e.g., `main`).
5. Check the option **Require a pull request before merging**.
6. Enable **Require approvals** and set the number of required approvals (at least 1).
7. Optionally, enable **Dismiss stale pull request approvals when new commits are pushed** to ensure new changes are reviewed.
8. Click **Create** or **Save changes**.

Now, GitHub will not allow merging into the protected branch until at least one reviewer approves the pull request.

---

## 2. GitHub Workflow to Trigger Tests on Commit

You can use GitHub Actions to automatically run tests whenever commits are pushed.

### Example: Simple Go Function and Test

**add.go**
```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	println(add(2, 3))
	fmt.Println(add(10, 20))

```

**add_test.go**
```go
package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

```

### GitHub Actions Workflow File

Create a file at `.github/workflows/go-tests.yml`:

```yaml
name: Go Tests

on:
  push:
    branches:
      - main
      - 'feature/*'
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'   

      - name: Install dependencies
        run: go mod tidy

      - name: Run tests
        run: go test -v ./...
```

### Explanation
- Workflow triggers on `push` and `pull_request` events to the `main` branch.
- Uses `actions/checkout` to pull repo code.
- Sets up Go environment.
- Runs `go test ./...` to execute all tests.

### Exercises:
- **Exercise 1:** Create a branch protection rule for your `main` branch requiring at least 1 review. Open a pull request and try to merge **without any review** → observe that GitHub blocks the merge.
- **Exercise 2:** To approve the pull request, you can:
- Use your secondary GitHub account, or
- Add me (`dwarakanathc44-debug`) as a collaborator so I can approve your PR.the pull request.
- **Exercise 3:**  Push your code (with `add.go` and `add_test.go`) to GitHub.
- Confirm that the workflow runs automatically when you commit.
- Modify the test to fail intentionally and observe GitHub Actions showing a failed check.
- Fix the test and push again → see the workflow succeed.


---

---
