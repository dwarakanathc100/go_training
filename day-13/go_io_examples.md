# Go Input and Output Examples with bufio, os.Args, and Writers


## 1. Using `bufio.NewReader(os.Stdin)`

The `bufio.NewReader` is used to read input from the user (keyboard) line by line or until a specific delimiter.

### Example:
```go
func bufferedReader() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n') // Reads input until newline
    name = strings.TrimSpace(name)      // Removes extra spaces or newline
    fmt.Printf("Hello, %s!\n", name)
}
```

### Explanation:
- `bufio.NewReader(os.Stdin)` creates a buffered reader connected to standard input (keyboard).
- `ReadString('\n')` reads input until the user presses **Enter**.
- `strings.TrimSpace(name)` removes extra whitespace and newline characters.
- The input is then printed back with a greeting.

---

## 2. Using `bufio.NewScanner(file)`

The `bufio.NewScanner` is useful for reading input line by line, especially from files.

### Example:
```go
func newScanner() {
    file, err := os.Open("test.md")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text() // Reads each line
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
```

### Explanation:
- `os.Open("test.md")` opens a file.
- `bufio.NewScanner(file)` creates a scanner to read the file line by line.
- `scanner.Text()` retrieves the current line as a string.
- This is useful for processing text files line by line efficiently.

---

## 3. Using `os.Args` (Command-Line Arguments)

The `os.Args` slice allows you to capture arguments passed to your program when it is executed.

### Example:
```go
func commandLineArgs() {
    args := os.Args
    if len(args) < 3 {
        fmt.Println("Please provide your name as a command-line argument.")
        return
    }
    name := args[1]
    age := args[2]
    fmt.Printf("name, %s!\n", name)
    fmt.Printf("age, %s!\n", age)
}
```

### Explanation:
- `os.Args` is a slice where:
  - `os.Args[0]` → the program name.
  - `os.Args[1]` → the first argument.
  - `os.Args[2]` → the second argument, and so on.
- This allows programs to accept user input directly when running from the terminal.

Example run:
```bash
go run main.go Ram 25
```
Output:
```
name, Ram!
age, 25!
```

---

## 4. Using `bufio.NewWriter(file)`

The `bufio.NewWriter` is used to write data into files efficiently using a buffer.

### Example:
```go
func csvwriter() {
    file, error := os.Create("details.csv")
    if error != nil {
        fmt.Println("Error creating file:", error)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    writer.WriteString("Name,Age,City\n")
    writer.WriteString("Alice,30,New York\n")
    writer.WriteString("Bob,25,Los Angeles\n")
    writer.WriteString("Charlie,35,Chicago\n")
    fmt.Println("CSV file created successfully.")
}
```

### Explanation:
- `os.Create("details.csv")` creates (or overwrites) a file.
- `bufio.NewWriter(file)` creates a buffered writer for the file.
- `writer.WriteString(...)` writes strings into the file.
- `writer.Flush()` ensures that all data in the buffer is written to the file.

---

See the go examples : https://go.dev/tour/basics/1
