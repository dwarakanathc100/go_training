# Arrays example and Slices

## Arrays with Loops

Example of arrays using `for` and `range` loops:

``` go
var utf [10000]string
for i, j := 0, 1; j <= 10000; i, j = i+1, j+1 {
    utf[i] = string(j)
}
fmt.Println(utf)

// iterate using for loop
for k := 0; k < len(utf); k++ {
    fmt.Printf(" %d %v %b \n", k, utf[k], k)
}

var a = [10]int{1, 2, 3, 4, 6, 7, 8, 10, 100}
for i, v := range a {
    fmt.Printf("index %d and Value %d \n", i, v)
}

arr := new([5]int)
arr[0] = 10
fmt.Println(*arr)
fmt.Println(&arr)

size := []int{10, 20, 30, 40, 50, 10, 20, 30, 40, 50, 100, 10000, 1002000}
fmt.Println(len(size))
fmt.Println(unsafe.Sizeof(size))
```

------------------------------------------------------------------------

## Slices 

### What is a Slice?

A slice is a flexible, powerful, and more commonly used data structure
than arrays in Go. Slices build on top of arrays and provide dynamic
sizing, making them very useful when the size of data is not fixed.

### Uses of Slices

-   Dynamic size: Unlike arrays, slices grow and shrink as needed.
-   More convenient: Provide built-in functions like `append`, `copy`,
    etc.
-   Efficient: Reference semantics avoid unnecessary copying.

### Disadvantages of Slices

-   Slices still depend on the underlying array; operations may
    reallocate memory if capacity is exceeded.
-   Can lead to memory leaks if large backing arrays are unintentionally
    referenced.
-   Performance overhead when resizing frequently.

------------------------------------------------------------------------

## Ways of Creating Slices

### 1. Empty Slice

An empty slice is declared without any elements.
It has length **0** and capacity **0**. It's useful when you want to
build a slice dynamically using `append`.

``` go
var s []int
fmt.Println(s)         // []
fmt.Println(len(s))    // 0
fmt.Println(cap(s))    // 0
```

------------------------------------------------------------------------

### 2. Slice Literal

A slice literal creates a slice and initializes it with values
directly.
This is the most common way to create slices when the values are already
known.

``` go
s1 := []int{1, 2, 4}
fmt.Println(s1)        // [1 2 4]
```

------------------------------------------------------------------------

### 3. Append to a Slice

Appending allows you to add elements to an existing slice.
If the slice exceeds its current capacity, Go automatically allocates a
new underlying array with a larger capacity.

``` go
s1 = append(s1, 4, 5, 6)
fmt.Println(s1)        // [1 2 4 4 5 6]
```

------------------------------------------------------------------------

### 4. Slice from an Array

You can create a slice by referencing a portion of an array.\
The slice shares the underlying array, so changes to the slice will
reflect in the array.

``` go
oldArray := [9]int{10, 20, 30, 40, 50, 50, 60, 70, 80}
sa := oldArray[1:4]
fmt.Println(sa)        // [20 30 40]
```

------------------------------------------------------------------------

### 5. Using `make()` Function

The `make()` function allows you to create a slice by specifying its
**length** and **capacity**.
This is useful when you know in advance how many elements you might
need, which avoids repeated memory reallocations.

``` go
makeSlice := make([]int, 3, 5)
fmt.Println(makeSlice) // [0 0 0]
makeSlice[0] = 100
makeSlice = append(makeSlice, 400)
fmt.Println(makeSlice) // [100 0 0 400]
```

------------------------------------------------------------------------

## Exercises 
1.  Create an array of 20 integers and print only even numbers using a
    `for` loop.
2.  Use a `range` loop to calculate the sum of elements in an array of
    size 10.
3.  Create a slice from an existing array and modify the slice. Observe
    changes in the original array.
4.  Use `append` to dynamically increase the size of a slice. Print its
    length and capacity after each append.
5.  Create a slice using `make()` with capacity 5. Append 10 elements
    and print how the capacity changes.
