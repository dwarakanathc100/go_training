


# Global log configuration and Arrays 
# 1. Global log configuration
-   **Dynamic Log Level Control**: Log level is read from the `.env`
    file. This allows changes to be applied **without rebuilding the
    code**.
-   **Centralized Management**: Instead of hardcoding log levels, we can
    use a **separate configuration repository** where `.env` files or
    configs are stored. This allows updating log levels for all services
    in one place.
-   **Environment Specific**: Developers can run in `debug` mode
    locally, while production can restrict logs to `error` or `error` or `warn` to
    avoid noise.
-   **Advantages of this approach**:
    1.  **Flexibility** -- Switch log levels on the fly.
    2.  **Consistency** -- All services can follow the same
        configuration by pulling from a central config repo.
    3.  **No Downtime** -- No need to redeploy applications when only
        log levels need to be updated.
    4.  **Improved Observability** -- Helps filter out unnecessary logs
        in production while keeping detailed logs in
        development/testing.

### Log Level Precedence

    Fatal > Error > Warn > Info > Debug > Trace
	
### Example Code

``` go
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/dwarakanathc100/uuid-generator/generator"
    "github.com/joho/godotenv"
    log "github.com/sirupsen/logrus"
)

func initLogger() string {
    err := godotenv.Load()
    if err != nil {
        log.Info("No .env file found")
    }

    level, _ := os.LookupEnv("LOG_LEVEL")
    fmt.Println("Log level: ", level)
    return level
}

func main() {
    level := initLogger()
    parsedLevel, _ := log.ParseLevel(strings.ToLower(level))
    fmt.Println("Consuming custom-uuid!")
    uuid := generator.GenerateUUID()
    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(parsedLevel)

    log.WithField("uuid", uuid).Debug("new UUID generated - debug")
    log.WithField("uuid", uuid).Info("new UUID generated - Info")
    log.WithField("uuid", uuid).Error("new UUID generated - Error")
    // log.WithField("uuid", uuid).Fatal("new UUID generated - Fatal")
    log.WithField("uuid", uuid).Panic("new UUID generated - Panic")
    log.WithField("uuid", uuid).Warn("new UUID generated - Warn")
    log.WithField("uuid", uuid).Trace("new UUID generated - Trace")
}


```


# 2. Arrays in Go

An array is a collection of elements of the same type, stored in
contiguous memory locations.

-   **Fixed size**: Once declared, the length of an array cannot change.
-   **Elements are indexed (0-based).**
-   **Go arrays are value types** (copying an array copies all its
    elements).

## Why Arrays are Needed

-   To store multiple items of the same type without declaring multiple
    variables.
-   Useful for fixed-size collections (e.g., days in a week, months in a
    year).
-   Provides index-based access to elements with **O(1)** time
    complexity.

## Advantages of Arrays

-   Fixed size → predictable memory usage.
-   Fast element access (O(1)).
-   Useful for representing fixed sets (like chessboard squares,
    months, days).
-   Can store data in contiguous memory, efficient for low-level
    operations.

## Disadvantages of Arrays

-   Size is fixed → cannot grow or shrink dynamically.
-   Copying arrays can be expensive for large data sets.
-   Not flexible for dynamic data structures.
-   Rarely used directly in Go (usually slices are preferred).

# 2. Declaring Arrays in Go 

### a) Specify Length and Values

``` go
var arr [3]int = [3]int{1, 2, 3}
```

### b) Compiler Infers Length

``` go
arr := [...]int{10, 20, 30}
```

### c) Declare Without Initialization (Default Zero Values)

``` go
var arr [5]int   // [0 0 0 0 0]
```

### d) Initialize Specific Indexes

``` go
arr := [5]int{0: 10, 3: 30} // [10 0 0 30 0]
```

### e) Multidimensional Arrays

``` go
var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
```

### f) Using new keyword (returns pointer)

``` go
arr := new([5]int)
arr[0] = 10
fmt.Println(*arr) // [10 0 0 0 0]
```

### Example Code

``` go
package main

import "fmt"

func main() {
    i, j, k, l := 10, 11, 12, 13
    fmt.Println(i, j, k, l)

    // 1. Defining arrays with explicit length
    var arr [3]int = [3]int{10, 20, 30}
    fmt.Println(arr)

    // 2. Using ellipsis (length inferred)
    arr1 := [...]int{10, 20, 30}
    fmt.Println(arr1)

    // 3. Empty array (default values: 0 for int)
    var arr2 [5]int
    fmt.Println(arr2)

    // 4. Using index-based initialization
    arr3 := [5]int{0: 10, 1: 30, 2: 20}
    fmt.Println(arr3)

    // 5. Multi-dimensional arrays (Matrix)
    var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println(matrix)

    // 6. Using 'new' keyword
    arr4 := new([5]int)
    arr4[0] = 100
    arr4[1] = 200
    fmt.Println(*arr4)
}
```
