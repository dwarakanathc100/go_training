# Go Concurrency, Parallelism, and Synchronization


------------------------------------------------------------------------

## **1. Understanding Concurrency and Parallelism**

**Concurrency** means dealing with multiple tasks at once, not
necessarily executing them simultaneously but managing them efficiently
so that tasks make progress together.

**Parallelism**, on the other hand, means executing multiple tasks at
the same time, actually running them on different CPU cores.

Go is designed to handle concurrency efficiently. It can run thousands
of lightweight goroutines within the same OS thread, switching between
them using the Go runtime scheduler.

------------------------------------------------------------------------

## **2. CPU, Cores, Threads, and Goroutines**

A **CPU** (Central Processing Unit) contains one or more **cores**. Each
core can handle one or more **threads** at a time. When you run a Go
program, the **Go scheduler** manages how goroutines (Go's lightweight
threads) are executed across available threads and cores.

Each goroutine runs on an OS thread, and Go's runtime automatically
handles mapping multiple goroutines to available threads. This means you
don't need to manually assign goroutines to cores, Go takes care of
that.

You can view system-level information using the `runtime` package, such
as:  `runtime.NumCPU()` --- number of logical CPUs. -
`runtime.NumGoroutine()` --- number of currently running goroutines. -
`runtime.GOOS` --- operating system. - `runtime.GOARCH` --- system
architecture.

------------------------------------------------------------------------
## 3. Example: Goroutines

``` go
go
package main

import (
    "fmt"
    "runtime"

    "time"
)



func test1() {
    for i := 0; i <= 10; i++ {
        time.Sleep(time.Millisecond * 2)
        fmt.Println("First loop :", i)
    }

}

func test2() {
    for i := 0; i <= 10; i++ {
        fmt.Println("Second loop :", i)
        time.Sleep(1 * time.Second)
    }
    
}

func main() {
    wg.Add(2)
    fmt.Println("Num of CPU:", runtime.NumCPU())
    fmt.Println("Num of Goroutines:", runtime.NumGoroutine())
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Arch:", runtime.GOARCH)

    go test1()
    go test2()

    fmt.Println("After creating goroutines")
    fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

}
```
Two goroutines are created to execute `test1()` and
`test2()` concurrently. The main goroutine exits wihout waiting for other go routines.So goroutines wont be able to finish the task.

## **4. Example: Goroutines and WaitGroups**

``` go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

var wg sync.WaitGroup

func test1() {
    for i := 0; i <= 10; i++ {
        time.Sleep(time.Millisecond * 2)
        fmt.Println("First loop :", i)
    }
    wg.Done()
}

func test2() {
    for i := 0; i <= 10; i++ {
        fmt.Println("Second loop :", i)
        time.Sleep(1 * time.Second)
    }
    wg.Done()
}

func main() {
    wg.Add(2)
    fmt.Println("Num of CPU:", runtime.NumCPU())
    fmt.Println("Num of Goroutines:", runtime.NumGoroutine())
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Arch:", runtime.GOARCH)

    go test1()
    go test2()

    fmt.Println("After creating goroutines")
    fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

    wg.Wait()
}
```

Two goroutines are created to execute `test1()` and
`test2()` concurrently.  The main goroutine waits for both to finish
using `wg.Wait()`.  `WaitGroup` ensures the program doesn't exit before
goroutines finish.

------------------------------------------------------------------------

## **5. WaitGroup Methods**

-   `Add(n)` --- adds n goroutines to wait for.
-   `Done()` --- signals that one goroutine is finished.
-   `Wait()` --- blocks until all goroutines complete.

Example:

``` go
package main

import (
    "fmt"
    "sync"
)

func work(wg *sync.WaitGroup, name string, n int) {
    defer wg.Done()
    for i := 0; i < n; i++ {
        fmt.Printf("name %s, i:%d 
", name, i)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go work(&wg, "Go", 10)
    go work(&wg, "GoLang", 20)
    wg.Wait()
}
```

------------------------------------------------------------------------

## **6. Race Conditions**

A **race condition** occurs when two or more goroutines access the same
shared resource (like a variable or map) simultaneously, and at least
one of them modifies it. This causes unpredictable results.

To detect race conditions:

    go run -race main.go

------------------------------------------------------------------------

## **6. Preventing Race Conditions with Mutex**

A **mutex** (short for mutual exclusion) allows only one goroutine to
access a section of code at a time.

``` go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var count int
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            mu.Lock()
            count++
            mu.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final count:", count)
}
```

 `mu.Lock()` ensures that only one goroutine modifies
`count` at a time.  Without the mutex, multiple goroutines could modify
it simultaneously, causing incorrect results.

------------------------------------------------------------------------

## **7. Atomic Package**

The `sync/atomic` package provides low-level atomic memory primitives
that allow goroutines to safely modify shared variables without using a
mutex.

### **Example: Using Atomic Operations**

``` go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var count int64 = 0
    var wg sync.WaitGroup
    workers := 1000
    wg.Add(workers)
    for i := 0; i < workers; i++ {
        go func() {
            defer wg.Done()
            atomic.AddInt64(&count, 1) // Atomic increment
        }()
    }
    wg.Wait()
    fmt.Println(count)
}
```

 `atomic.AddInt64()` ensures the increment operation
happens atomically. - It prevents data races and eliminates the need for
mutexes for simple operations.

------------------------------------------------------------------------

## **8. Summary**

  -----------------------------------------------------------------------
  Concept                       Description
  ----------------------------- -----------------------------------------
  **Concurrency**               Managing multiple tasks at once.

  **Parallelism**               Running multiple tasks simultaneously.

  **Goroutine**                 Lightweight thread managed by the Go
                                runtime.

  **WaitGroup**                 Synchronizes goroutines and waits for
                                completion.

  **Mutex**                     Prevents simultaneous access to shared
                                data.

  **Atomic**                    Provides low-level, thread-safe
                                operations without locks.
  -----------------------------------------------------------------------

------------------------------------------------------------------------

