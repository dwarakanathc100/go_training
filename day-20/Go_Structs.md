
#  Structs in Go 

##  Introduction
A **struct** (short for *structure*) in Go is a **composite data type** that groups together variables (called *fields*) under a single name.  
It allows you to represent different types of data as one logical unit — making your code cleaner, more organized, and easier to manage.

Structs are used when:
- You want to represent **real-world entities** (like Student, Employee, Product, etc.)
- You need to **combine multiple data types** into one entity.
- You want to **avoid using deeply nested maps**, which become hard to read and maintain.

---

## Why Structs Over Maps?

| Feature | Map | Struct |
|----------|------|--------|
| Key Type | Keys must be of one type (e.g., string, int). | Fields can have different types. |
| Flexibility | Dynamic and can grow easily. | Static — field names and types are fixed. |
| Readability | Complex nested maps are difficult to read. | Structs are organized and self-descriptive. |
| Validation | No built-in field validation. | Can define typed and validated fields. |
| Use Case | For dynamic key-value storage. | For fixed structure representation. |

---

##  Basic Struct Syntax

```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    Field3 Type3
}
```

### Example:
```go
type Student struct {
    Name  string
    Age   int
    Grade string
}
```

---

##  Creating and Initializing Structs

There are **multiple ways** to define and initialize structs in Go.

### 1 **Using a Struct Literal**
```go
s := Student{Name: "Ram", Age: 21, Grade: "A"}
```

### 2 **Using the `var` Keyword**
```go
var s Student
s.Name = "Syam"
s.Age = 22
s.Grade = "B"
```

### 3 **Using the `new()` Function**
```go
s := new(Student)
s.Name = "Ram"
s.Age = 23
s.Grade = "A+"
```

### 4 **Anonymous Structs (No Name Type)**
Used for temporary or one-time data structure.
```go
person := struct {
    Name string
    Age  int
}{
    Name: "Ram",
    Age:  10,
}
```

### 5 **Pointer to Struct**
Structs can be used with pointers to avoid copying data.
```go
s := &Student{Name: "Kiran", Age: 20}
fmt.Println(s.Name)
```

---

##  Example: Structs for Student Data

```go
package main

import "fmt"

func StructType() {
	type Subject struct {
		Name  string
		Marks int
	}

	type Student struct {
		ID       int
		Name     string
		Subjects []Subject
		Grade    string
		Result   string
	}

	type Course struct {
		Name     string
		Students []Student
	}

	type College struct {
		Name   string
		Course []Course
	}

	college := College{
		Name: "ABC College",
		Course: []Course{
			{
				Name: "MBA",
				Students: []Student{
					{
						ID:   1001,
						Name: "Ram",
						Subjects: []Subject{
							{Name: "Computers", Marks: 90},
							{Name: "Maths", Marks: 80},
							{Name: "Stats", Marks: 70},
						},
						Grade:  "A+",
						Result: "PASS",
					},
				},
			},
			{
				Name: "MCA",
				Students: []Student{
					{
						ID:   1002,
						Name: "Syam",
						Subjects: []Subject{
							{Name: "Computers", Marks: 88},
							{Name: "Maths", Marks: 76},
							{Name: "Stats", Marks: 68},
						},
						Grade:  "A",
						Result: "PASS",
					},
				},
			},
		},
	}

	fmt.Println(college)
}
```

---

##  Accessing Struct Fields
```go
fmt.Println(college.Name)
fmt.Println(college.Course[0].Students[0].Subjects[0].Name)
```

---

##  Exercises


Modify the `StructType()` program to calculate `Grade` and `Result` **dynamically** instead of hardcoding them.

###  Instructions:
1. Create a **new function** named `CalculateGrade(marks []int) (string, string)` that:
   - Takes a list of marks as input.
   - Calculates the **average score**.
   - Returns **Grade** and **Result** dynamically.

2. Use the following grading logic:
   - Average ≥ 90 → Grade: A+, Result: PASS  
   - Average ≥ 75 and < 90 → Grade: A, Result: PASS  
   - Average ≥ 50 and < 75 → Grade: B, Result: PASS  
   - Otherwise → Grade: F, Result: FAIL  

3. Modify your main struct to call `CalculateGrade()` when assigning values to each student.

### Example Function:
```go
func CalculateGrade(marks []int) (string, string) {
	total := 0
	for _, m := range marks {
		total += m
	}
	avg := total / len(marks)

	if avg >= 90 {
		return "A+", "PASS"
	} else if avg >= 75 {
		return "A", "PASS"
	} else if avg >= 50 {
		return "B", "PASS"
	}
	return "F", "FAIL"
}
```

###  Expected Output (Sample)
```
College Name: ABC College
Course: MCA
Student: Ram | Grade: A+ | Result: PASS
Student: Syam | Grade: A | Result: PASS
```

---


