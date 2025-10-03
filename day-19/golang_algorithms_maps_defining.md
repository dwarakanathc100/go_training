# Algorithms and Maps in Go

## 1. Algorithm for Searching an Element in a Slice

Before writing any program, we usually define an **algorithm**. 
An algorithm is a step-by-step procedure to solve a problem.

### Example Problem

Search for an element (6) in a slice `[1, 2, 4, 5, 6]`.

### Steps (Pseudo Algorithm)

1.  Define a variable with the value we are looking for (e.g., 6). 
2.  Loop through the slice (iterate each element). 
3.  Compare the iterated value with the declared value. 
4.  If a match is found, set a boolean flag and print the index. 
5.  Break the loop once the value is found. 
6.  If not found, print a message saying "value not found".

### Implementation

``` go
func algo() {
    s := []int{1, 2, 4, 5, 6}
    val := 6
    found := false

    for i, v := range s {
        if v == val {
            found = true
            fmt.Println("value found at index", i)
            break
        }
    }

    if !found {
        fmt.Println("value not found")
    }
}
```

### Time Complexity

-   In the **worst case**, we may need to check every element of the
    slice.
    -   Example: searching for 7 (not present) → must check all `n`
        elements. 
-   Hence, **Time Complexity = O(n)** (linear time). 
-   **Space Complexity = O(1)** (only using a few variables).

------------------------------------------------------------------------

## 2. Maps in Go

### Why Maps?

-   Arrays and slices store elements sequentially and are good for
    **indexed access**. 
-   But searching by value in an array/slice takes **O(n)** time. 
-   **Maps** allow us to store data as **key-value pairs** and retrieve
    values in **O(1) average time**.

------------------------------------------------------------------------

## 3. Map Declaration in Go

### 1. Nil Map

``` go
var m map[string]int
fmt.Println(m == nil) // true
```

-   A map declared but not initialized is called a **nil map**. 
-   It has no memory allocated, so you **cannot insert values** into
    it. 
-   You can only compare it with `nil` or read from it (which always
    returns zero values).

### 2. Using `make()`

``` go
makeMap := make(map[string]int)
makeMap["A"] = 20
makeMap["B"] = 19
makeMap["V"] = 23
makeMap["D"] = 26

```

-   The most common way to create a map. 
-   `make()` allocates memory and returns an initialized, empty map. 
-   You can directly start adding key-value pairs.

### 3. With Capacity 

``` go
mCap := make(map[string]int, 100)
```

-   Similar to using `make()`, but you also provide a **capacity
    hint**. 
-   Tells the runtime to optimize for storing around 100 entries. 
-   The map still grows dynamically if more entries are added.

### 4. Map Literal

``` go
literalMap := map[string]int{"A": 10, "B": 20, "C": 30}
```

-   A quick way to create a map with predefined key-value pairs. 
-   Very useful for small maps or configuration-like data.

### 5. Empty Map (Literal without values)

``` go
emptyMap := map[string]int{}
emptyMap["City"] = 10000
```

-   Similar to `make()`, but written as an empty literal. 
-   Allocates memory and gives an initialized empty map. 
-   You can directly insert values.

------------------------------------------------------------------------

## 4. Retrieving Values from a Map

``` go
value, ok := makeMap["V"] // 40 true
value, ok = makeMap["E"]
fmt.Println(value, ok) // 0 false (since "E" not present)
```

-   `ok` idiom → returns `true` if the key exists, else `false`.

------------------------------------------------------------------------

## 5. Time Complexity of Maps

-   **Average case:** O(1) → because of hashing. 
-   **Worst case:** O(n) (very rare, when many keys hash to the same
    bucket).

------------------------------------------------------------------------

## 6. How Maps Work Internally

-   Go maps use **hashing** to store key-value pairs. 
-   A hash function converts a key into an integer (hash code). 
-   This integer determines the **bucket index**. 
-   Keys inside the same bucket are managed internally (like linked list
    / array). 
-   Because keys are assigned to buckets based on hash values, **maps
    are unordered**.

### Example Visualization of Buckets

    Index   Bucket (Key -> Value)
    0       A -> 20
    1       B -> 19
    2       V -> 23
    3       D -> 26
    ...

-   When searching for `"A"`, Go:
    1.  Hashes `"A"` → gets bucket index. 
    2.  Looks inside bucket for key `"A"`. 
    3.  Returns value if found.

### Why Maps are Fast?

-   Instead of looping through all elements like a slice, we just
    compute hash → bucket index. 
-   This makes **lookup almost constant time**.

------------------------------------------------------------------------

## 7. Hashing Explanation

-   **Hash Function:** Converts input key (string, int, etc.) into a
    fixed integer. 
-   Example: `"A"` → hash code `0x1234...` → mapped to a bucket. 
-   Keys that collide (same bucket) are resolved by Go internally.


