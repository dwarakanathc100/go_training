# Golang Structs, JSON Formatting, and Validation


In this session, we discussed about how to **expand structs**, represent hierarchical data, convert data to **JSON format**, and validate structs using the **Go validator** library.

We discussed the following key topics:
- How to expand and nest structs (e.g., `University`, `College`, `Course`, `Student`, `Subject`).
- Why JSON format is important and where it is commonly used.
- How to omit empty fields in JSON using struct tags (`omitempty`).
- How to enforce mandatory fields using struct validation rules (`validate:"required"`).
- How to use Go’s `validator` package for data validation.

---

## 1. Struct Expansion in Go



A **struct** in Go is a composite data type that groups together variables (fields) under a single name. Structs are useful for representing real world entities like students, courses, and universities.

You can expand a struct by **nesting** other structs within it, creating a clear data hierarchy.

### Example

```go
type Subject struct {
    Name  string `json:"name,omitempty"`
    Marks int    `json:"marks,omitempty"`
}

type Student struct {
    ID       int       `json:"id,omitempty"`
    Name     string    `json:"name,omitempty"`
    Subjects []Subject `json:"subjects,omitempty"`
    Grade    string    `json:"grade,omitempty"`
    Result   string    `json:"result,omitempty"`
}

type Course struct {
    Name     string    `json:"name,omitempty"`
    Students []Student `json:"students,omitempty"`
}

type College struct {
    Name   string   `json:"name,omitempty"`
    Course []Course `json:"course,omitempty"`
}

type Univercity struct {
    Name     string    `json:"name,omitempty" validate:"required"`
    Colleges []College `json:"colleges,omitempty"`
}
```

Here, the `University` struct is expanded to include `Colleges`, each of which contains `Courses`, which in turn include `Students` and `Subjects`.

---

## 2. JSON Formatting in Go

### Why JSON Format Is Useful

**JSON (JavaScript Object Notation)** is a lightweight data format widely used for data interchange between systems. It is human-readable, easy to parse, and language-independent.

### Use Cases
- Exchanging data between frontend and backend (e.g., APIs).
- Storing configurations or structured data.
- Serializing complex nested data structures.

### Example: Convert Struct to JSON

```go
uniJson, _ := json.MarshalIndent(univercity, "", "  ")
fmt.Println("University JSON:", string(uniJson))
```

### `omitempty` Tag
When `omitempty` is used, fields with empty or zero values are **excluded** from the JSON output.  
For example, if `Name` is an empty string, it won’t appear in the JSON.

---

## 3. Struct Validation Using Go Validator

### Why Validate Structs?

Validation ensures data integrity and avoids processing invalid data (e.g., missing names, negative marks).  
We use the `go-playground/validator` package to validate struct fields.

### Example

```go
validate := validator.New(validator.WithRequiredStructEnabled())

univercity := Univercity{
    Name: "", // Missing required field
}

err := validate.Struct(univercity)
if err != nil {
    fmt.Println("Validation Error:", err)
    return
}
```

### Explanation
- `validate:"required"` → ensures the field cannot be empty.
- If `Name` is empty, validation will fail.

---

## 4. Example Program

```go
func StructType() {
    validate := validator.New(validator.WithRequiredStructEnabled())

    // Struct definitions and data initialization as shown above...

    err := validate.Struct(univercity)
    if err != nil {
        fmt.Println("validation error", err)
        return
    }

    uniJson, _ := json.MarshalIndent(univercity, "", "  ")
    fmt.Println("University JSON:", string(uniJson))
}
```

---

## 5. Exercises

### Exercise 1: Validate Required Fields
- Add a validation rule to make **Student Name** required.
- Try running the program with an empty student name and observe the validation error.

### Exercise 2: Validate Numeric Ranges
- Add validation for `Marks` to ensure marks are between 0 and 100.
  ```go
  Marks int `json:"marks,omitempty" validate:"gte=0,lte=100"`
  ```
- Test the validation by setting `Marks: 120` and observing the error.

### Exercise 3: JSON Output Practice
- Remove some field values and print JSON output. Check if omitted fields are excluded.


## 6. Refer the below links for validator tags and json

- [Go Validator Package](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [Go JSON Encoding](https://pkg.go.dev/encoding/json)
