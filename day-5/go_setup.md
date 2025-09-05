

# Introduction to Go Programming (Day 5)


## 1. What is Go?
- Go is an open-source programming language developed by Google.
- Designed for **simplicity, concurrency, and performance**.
- Great for building **scalable backend systems, APIs, cloud-native applications, and DevOps tools**.

Official website: [https://go.dev/](https://go.dev/)

---

## 2. Install Go

### Download & Install
- Official download page: [https://go.dev/dl/](https://go.dev/dl/)
- Select the installer based on your OS:
  - **Windows:** MSI installer
  - **MacOS:** PKG installer or Homebrew (`brew install go`)
  - **Linux:** tarball installation or package manager (`sudo apt install golang-go`)

### Verify Installation
After installation, run:
```bash
go version
```
This should print the installed Go version.

---

## 3. Go Environment Variables

Go uses two important environment variables:

### `GOROOT`
- Points to where Go is installed.
- Example (Windows): `C:\Program Files\Go`
- Example (Linux/Mac): `/usr/local/go`

### `GOPATH`
- Workspace where your Go projects, source code, and dependencies reside.
- By default:
  - **Linux/Mac:** `$HOME/go`
  - **Windows:** `%USERPROFILE%\go`

### Add to PATH
Ensure `go` binary is available in your terminal by adding to PATH:

#### Linux/Mac (add to `~/.bashrc` or `~/.zshrc`):
```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

#### Windows:
- Add `C:\Program Files\Go\bin` and `%USERPROFILE%\go\bin` to PATH in **Environment Variables**.

Verify with:
```bash
go env
```

---

## 4. Go Playground
- Online tool to run and share Go code snippets.
- [https://go.dev/play/](https://go.dev/play/)

This is great for testing small programs without local setup.

---

## 5. Setting Up Visual Studio Code (VSCode)

### Install VSCode
- Download: [https://code.visualstudio.com/download](https://code.visualstudio.com/download)

### Install Go Plugin
- Open VSCode → Extensions (`Ctrl+Shift+X` / `Cmd+Shift+X`)
- Search for **Go (by Go Team at Google)**
- Install it → Provides linting, IntelliSense, debugging, and formatting.


## 6. Writing Your First Go Program

Create a file `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Gophers!")
}
```

Run it:
```bash
go run hello.go
```

Build it:
```bash
go build hello.go
./hello   # (Linux/Mac)
hello.exe # (Windows)
```

---

## 7. Useful Links
- Official Docs: [https://go.dev/doc/](https://go.dev/doc/)
- Effective Go (best practices): [https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)
- Go by Example: [https://gobyexample.com/](https://gobyexample.com/)
- Tour of Go (interactive): [https://go.dev/tour/](https://go.dev/tour/)
- Standard Library Reference: [https://pkg.go.dev/std](https://pkg.go.dev/std)

---

## 8. Summary
Today you:
1. Installed Go
2. Set up environment variables (`GOROOT`, `GOPATH`)
3. Installed VSCode & Go extension
4. Learned about Go Playground
5. Wrote your first Go program

Next step: We will dive into **Go basics – variables, types, functions, and control flow.** 🚀

