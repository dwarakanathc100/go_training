# Some Basic Concepts:


---

## 1. fmt Package Printing Formats

 
The `fmt` package in Go is part of the standard library and provides I/O formatting functions. It allows developers to print text, format variables, and build strings with custom formatting. Functions like `fmt.Println`, `fmt.Printf`, and `fmt.` are commonly used.  

 
Format verbs (like `%d`, `%s`, `%q`) act as placeholders in strings, instructing Go how to display a value. This provides flexibility in debugging and presenting values in different bases, character forms, or Unicode.  

### Commonly Used Format Verbs

| Verb | Meaning |
|------|----------|
| `%b` | Base 2 (binary) |
| `%c` | The character represented by the corresponding Unicode code point |
| `%d` | Base 10 (decimal) |
| `%o` | Base 8 (octal) |
| `%O` | Base 8 with `0o` prefix |
| `%q` | A single-quoted character literal safely escaped with Go syntax |
| `%x` | Base 16, with lower-case letters for a-f |
| `%X` | Base 16, with upper-case letters for A-F |
| `%U` | Unicode format: `U+1234`; same as `"U+%04X"` |

### Example

```go
a := 'A'
fmt.Println(a)
fmt.Printf("the value in %v is %c \n", a, a)
fmt.Printf("the value %v and in binary %b \n", a, a)
fmt.Printf("the value %v and in octa %o \n", a, a)
fmt.Printf("the value %v and in hexa %X \n", a, a)
fmt.Printf("the value %v and in utf %U \n", a, a)
fmt.Printf("the value %v and in quotation %q \n", a, a)
```

---

## 2. String Literals


A **string literal** in Go is a sequence of characters enclosed within quotes. It represents constant text values in a program.  

  
Go supports two forms of string literals:  
1. **Interpreted string literals**: Enclosed in double quotes `""`, supporting escape sequences (`\n`, `\t`, etc.).  
2. **Raw string literals**: Enclosed in backticks `` ` ```, which preserve the exact formatting of the text, including newlines and tabs.  

This makes raw strings especially useful for embedding multi-line text, SQL queries, JSON, or regular expressions.  

### Example

```go
// Interpreted string literal
s := "Hello\nWorld"

// Raw string literal
s := `Ironically, worrying about the effects of stress can make it worse.

Stress is a normal part of life and is unavoidable.`

fmt.Println(s)
```

---

## 3. Type Conversions

 
**Type conversion** in Go is the process of converting a value from one data type to another, explicitly specified by the programmer.  

 
Go is a **statically typed language**, which means variables have a fixed type at compile time. Unlike some languages, Go does not perform implicit type casting. This strict rule ensures type safety but requires the developer to explicitly convert values when necessary.  

Conversions follow the format:  
```go
T(v)
```
Where `T` is the target type and `v` is the value being converted.  

### Example

```go
var i int = 42
var f float64 = float64(i) // int → float64
var h float32 = 13.8
var j float64 = float64(h) // float32 → float64

fmt.Printf("the value of j is %v and its type is %T \n", j, j)
fmt.Printf("the value of i is %v and its type is %T \n", i, i)
fmt.Printf("the value of f is %v and its type is %T \n", f, f)
```

### Key Points
- No automatic type conversions.  
- Explicit casting prevents unexpected data loss.  
- Type safety is enforced at compile time.  

---

## 4. Control Statements


**Control statements** in Go are used to determine the flow of program execution based on conditions or values.  

 
Go provides:  
1. **`if-else`**: Conditional branching.  
2. **`switch-case`**: Cleaner alternative for multiple conditions.  
3. **Loops (`for`)**: To repeat actions (not covered here but equally important).  

These statements allow decision-making and execution of code blocks selectively.  

### Examples

#### If-Else Statement
```go
z := 100
if z > 100 {
    fmt.Println("z is greater than 100")
} else if z >= 100 {
    fmt.Println("z is greater than or equal to 100")
} else {
    fmt.Println("z is less than 100")
}
```

#### Switch Statement
```go
day := 3
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid day")
}
```

---

## Practice Exercises

1. Print a character in all available formats (`%b`, `%c`, `%d`, `%o`, `%x`, `%X`, `%U`, `%q`).  
2. Create an interpreted string with escape sequences and a raw string with multiple lines. Print both.  
3. Convert an `int` to `float64` and back. Print the type and value each time.  
4. Write a program using `if-else` to check if a number is positive, negative, or zero.  
5. Write a `switch` statement that prints the name of the month given its number (1 = January, 12 = December).  

---




