#  Exporting Data to CSV & Understanding Pointers

---



### CSV Introduction
CSV (Comma-Separated Values) files are a simple format used to store tabular data such as student grades, financial records, or logs.  
In Go, we can easily create and write data to a CSV file using the built-in **`encoding/csv`** and **`os`** packages.

---

###  Why Export Data to CSV?
- Easy to share and import into tools like Excel, Google Sheets, or databases.  
- Helps in data reporting and analysis.  
- Simple and lightweight file format.  

---

###  Example: Writing Student Data to CSV

```go
file, err := os.Create("student_data.csv")
if err != nil {
	fmt.Println("error while creating file", err)
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

fmt.Println("Univercity Name:", univercity.Name)
writer.Write([]string{"Univercity Name", "College Name", "Course Name", "Student Name", "Grade", "Result"})

for _, college := range univercity.Colleges {
	fmt.Println("College Name:", college.Name)
	for _, course := range college.Course {
		fmt.Println("Course Name :", course.Name)
		for _, student := range course.Students {
			var marks []int
			for _, sub := range student.Subjects {
				marks = append(marks, sub.Marks)
			}
			grade, result := CalculateGrade(marks)
			record := []string{univercity.Name, college.Name, course.Name, student.Name, grade, result}
			writer.Write(record)
		}
	}
}
```

---

###  Step-by-Step Explanation

1. **Create a CSV File**
   ```go
   file, err := os.Create("student_data.csv")
   ```
   - Creates a new file named `student_data.csv`.
   - If the file exists, it overwrites it.

2. **Initialize CSV Writer**
   ```go
   writer := csv.NewWriter(file)
   defer writer.Flush()
   ```
   - Creates a CSV writer object that buffers and writes records.
   - `Flush()` ensures any remaining data is written to disk.

3. **Write CSV Headers**
   ```go
   writer.Write([]string{"Univercity Name", "College Name", "Course Name", "Student Name", "Grade", "Result"})
   ```
   - Writes the column titles (header row) into the CSV file.

4. **Iterate Over Nested Structures**
   - Loops through university → college → course → student.
   - Builds records dynamically using data.

5. **Write Records**
   ```go
   writer.Write(record)
   ```
   - Each record represents one student’s full data.

6. **Close File**
   - `defer file.Close()` ensures file closes safely even if errors occur.

---


##  2. Understanding Pointers in Go

### What Are Pointers?
A **pointer** is a variable that stores the **memory address** of another variable.

Think of it as a reference that allows indirect access to a value stored somewhere in memory.

---

### Why Use Pointers?
-  **Efficiency:** Avoid copying large structures or arrays.
-  **Modify values directly:** Functions can change variables outside their scope.
-  **Memory Management:** Helps manage references efficiently.
-  **Go Routines & Concurrency:** Used to share data safely.

---

### Example: Pointer in Go

```go
func Pointes() {
	x := 10 // normal variable
	var p *int = &x
	fmt.Println(x)   // prints value of x
	fmt.Println(p)   // prints memory address of x
	fmt.Println(*p)  // dereference: get value of x using pointer

	// using new() function
	p1 := new(int)
	*p1 = 50
	fmt.Println(*p1)

	updateXValue(&x)
	noupdatex(x)
	fmt.Println("updated x value :", x)
}

func updateXValue(num *int) {
	*num = *num + 100
}

func noupdatex(n int) {
	n = n + 100
}
```

---

###  Explanation

| Concept | Description | Example |
|----------|--------------|---------|
| **Address-of Operator (`&`)** | Returns the memory address of a variable. | `p := &x` |
| **Dereference Operator (`*`)** | Accesses or modifies the value at a memory address. | `*p = 20` |
| **Pointer Declaration** | Declares a variable as a pointer type. | `var p *int` |
| **Using `new()`** | Allocates memory for a variable and returns a pointer to it. | `p := new(int)` |

---

###  Flow

1. `x := 10` → Variable `x` stores `10`.
2. `p := &x` → Pointer `p` stores address of `x`.
3. `*p` → Dereferencing `p` gives `10`.
4. `updateXValue(&x)` → Sends address of `x` to function → modifies actual variable.
5. `noupdatex(x)` → Passes copy of value → does not modify original.

---

### Real-Time Use Cases of Pointers in Go

| Use Case | Description |
|-----------|--------------|
| **Struct Updates** | Update struct data in functions without copying large data. |
| **Linked Lists / Trees** | Nodes use pointers to reference next elements. |
| **Database Connections** | Pass pointers to maintain shared state. |
| **JSON Unmarshalling** | `json.Unmarshal` uses pointers to assign parsed data. |
| **Performance Optimization** | Reduces memory usage by avoiding deep copies. |

---



###  Exercises

1. Modify your CSV export to include **total marks** and **average** columns.
2. Write a function using pointers that **swaps two integers**.
3. Create a struct `Employee` and pass its pointer to a function to update the salary.
4. Use `new()` to allocate memory for a float and assign a value through the pointer.
5. Explain why `noupdatex()` doesn’t change `x` in your own words.

---

