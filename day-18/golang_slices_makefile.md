# Golang Slices Advanced Concepts and Makefile Tutorial

------------------------------------------------------------------------

## 1. Copying Slices using `copy()`


The `copy()` function is used to copy elements from one slice to
another. 
- The number of elements copied will be the **minimum length** of the
source and destination slices. 
- It's useful when you need to duplicate data or move it into another
slice.

### Example:

``` go
func copyEx() {
    src := []int{2, 3, 4, 5, 6, 7}
    dest := make([]int, 3)
    copy(dest, src)

    fmt.Println(dest) // Output: [2 3 4 ]
}
```

------------------------------------------------------------------------

## 2. Sub-slices


A sub-slice is created using the syntax `s[start:end]`. 
- It includes elements from `start` index to `end-1`. 
- Useful when you want to work on a part of a slice without creating a
new one.

### Example:

``` go
func subslice() {
    s := []int{1, 2, 3, 4, 5, 6, 7, 8}
    sub := s[1:4]

    fmt.Println(sub) // Output: [2 3 4]
}
```

------------------------------------------------------------------------

## 3. Delete Slice Elements


Go doesn't have a built-in delete for slices. 
We use slicing + `append()` to remove elements.

### Example:

``` go
func deleteEx() {
    s := []int{1, 2, 3, 4, 5, 6, 7}
    i := 2
    s = append(s[:i], s[i+1:]...) // Removes element at index 2

    fmt.Println(s) // Output: [1 2 4 5 6 7]
}
```

------------------------------------------------------------------------

## 4. Nil vs Empty Slice


-   A **nil slice** has no underlying array. 
-   An **empty slice** has length 0 but is not nil. 
    Useful to differentiate when checking initialization.

### Example:

``` go
func nilCheck() {
    var s []int
    if s == nil {
        fmt.Println("slice is nil")
    }

    s = []int{}
    if len(s) == 0 {
        fmt.Println("slice is empty but not nil")
    }

    fmt.Println("check if s is nil ?", s == nil)
}
```

------------------------------------------------------------------------

## 5. String Slice


Slices can hold strings. `append()` can be used to join two slices or
add elements.

### Example:

``` go
func stringSlice() {
    s := []string{"Monday", "Tue", "Wed"}
    sub := []string{"Thu", "Fri", "Sat", "Sun"}

    s = append(s, sub...)
    fmt.Println(s) // [Monday Tue Wed Thu Fri Sat Sun]

    s = append(s, "Jan", "Feb")
    fmt.Println(s) // [Monday Tue Wed Thu Fri Sat Sun Jan Feb]
}
```

------------------------------------------------------------------------

## 6. Multi-Dimensional Slice


Slices can be multi-dimensional, allowing us to store complex data like
tables or records.

### Example:

``` go
func multiDimSlice() {
    var recordsOfRecords [][][]string

    var records1 [][]string
    student1 := []string{"Maths", "Science", "Social", "Eng"}
    student2 := []string{"Science", "Social", "Eng"}

    records1 = append(records1, student1)
    records1 = append(records1, student2)

    var records2 [][]string
    student3 := []string{"Social", "Science", "Maths", "Eng"}
    student4 := []string{"Science", "Social", "Eng"}

    records2 = append(records2, student3)
    records2 = append(records2, student4)

    recordsOfRecords = append(recordsOfRecords, records1)
    recordsOfRecords = append(recordsOfRecords, records2)

    fmt.Println(recordsOfRecords)
}
```

------------------------------------------------------------------------

## 7. Increment Slice Elements


You can update slice elements directly using indexing.

### Example:

``` go
func incrementSlice() {
    s := make([]int, 1)
    s[0] = 10
    s[0]++

    fmt.Println(s) // Output: [11]
}
```

------------------------------------------------------------------------

# Makefile

## Why Makefile?

-   Automates repetitive commands like build, test, run, clean. 
-   Makes projects easier to maintain and collaborate on. 
-   Common in large projects for consistency.

## Example Makefile:

``` makefile
# Simple Makefile for Golang Project

APP_NAME=myapp

build:
    go build -o $(APP_NAME) main.go

run: build
    ./$(APP_NAME)

test:
    go test ./...

clean:
    rm -f $(APP_NAME)
```

## Usage:

``` bash
make build   # Builds the project
make run     # Runs the project
make test    # Runs all tests
make clean   # Removes built files
```

## Tutorial Link:

[Makefile in Go Projects](https://makefiletutorial.com/)

------------------------------------------------------------------------

# Exercises 

1.  Create a slice of integers and copy it into another slice using
    `copy()`.
2.  Extract sub-slices of different ranges and print them.
3.  Write a function to delete the 1st, last, and middle element from a
    slice.
4.  Differentiate between nil slice and empty slice by printing them.
5.  Create a string slice for days of the week and append months to it.
6.  Create a 2D slice representing marks of 3 students in 3 subjects.
7.  Create a Makefile that compiles and runs a simple Go program.

------------------------------------------------------------------------
