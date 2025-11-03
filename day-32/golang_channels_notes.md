# Golang Channels - Notes with Examples

## 1. Introduction to Channels in Go

Channels are Go’s built-in mechanism for **safe communication** between goroutines. They help goroutines exchange data and synchronize execution without explicit locking mechanisms like mutexes.

Think of channels as **pipes** through which data flows — one goroutine sends data, another receives it.

### Basic Syntax
```go
ch := make(chan int)       // unbuffered channel
ch := make(chan int, 3)    // buffered channel with capacity 3
```

### Channel Operations
| Operation | Syntax | Description |
|------------|--------|--------------|
| Send | `ch <- value` | Sends value to channel |
| Receive | `val := <-ch` | Receives value from channel |
| Close | `close(ch)` | Closes channel (no more sends allowed) |

---

## 2. Why Do We Need Channels?

- To **communicate** between goroutines safely.
- To **avoid race conditions** without explicit synchronization tools.
- To **coordinate** the execution of concurrent tasks.
- To **simplify data sharing** — no need for complex mutex logic.

---

## 3. Types of Channels in Go

### 3.1 Unbuffered Channels
Unbuffered channels don’t store any data. The send and receive must happen **simultaneously**.

**Example:**
```go
func main1() {
    ch := make(chan string)
    go func() { ch <- "hi" }() // send
    msg := <-ch // receive
    fmt.Println(msg)
}
```

**Key Points:**
- The sender blocks until the receiver is ready.
- The receiver blocks until a value is available.

**Use Case:** Synchronization between goroutines.

**Another Example:**
```go
func main2() {
    c := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("sending", i)
            c <- i
            time.Sleep(1 * time.Second)
        }
    }()

    go func() {
        for {
            time.Sleep(1 * time.Second)
            fmt.Println("receiving", <-c)
        }
    }()

    time.Sleep(10 * time.Second)
}
```
---

### 3.2 Buffered Channels
Buffered channels have a **capacity** that allows sending a fixed number of elements without an immediate receiver.

**Example:**
```go
func main() {
    ch := make(chan int, 3) // buffer capacity = 3
    go func() {
        for i := 0; i < 14; i++ {
            ch <- i
        }
    }()

    fmt.Println("Buffered channel length:", len(ch))

    for i := 0; i < 14; i++ {
        fmt.Println("received", <-ch)
        time.Sleep(3 * time.Second)
    }
}
```

**Key Points:**
- Sending blocks when the buffer is full.
- Receiving blocks when the buffer is empty.
- Useful for decoupling producers and consumers.

---

### 3.3 Directional Channels
Go allows creating **send-only** and **receive-only** channels.

**Example:**
```go
func sendData(ch chan<- int) {
    for i := 1; i <= 3; i++ {
        fmt.Println("Sending data:", i)
        ch <- i
    }
    close(ch)
}

func receiveData(ch <-chan int) {
    for val := range ch {
        fmt.Println("Received:", val)
    }
}

func main() {
    ch := make(chan int)
    go sendData(ch)
    receiveData(ch)
}
```
**Use Case:** Prevent accidental misuse e.g., ensuring a function only sends data.

---

### 3.4 Ranging Over Channels

When iterating over a channel using `range`, Go doesn’t know how many elements will come, so you **must close the channel** once all sends are done. Otherwise, the loop will block forever.

**Example:**
```go
func main3() {
    c := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()

    for n := range c { // Must close channel, else infinite wait
        fmt.Println(n)
    }
}
```

If you don’t close a channel that’s being ranged over, it leads to **deadlock** (panic).

---

### 3.5 Checking Channel Status Using `ok`
When receiving from a closed channel, Go returns two values — the received value and a boolean `ok`.

**Example:**
```go
msg, ok := <-ch
fmt.Println(msg, ok)
```
If `ok == false`, the channel is closed.

**Full Example:**
```go
func main5() {
    ch := make(chan string)

    go func() {
        for i := 1; i < 4; i++ {
            ch <- fmt.Sprintf("message #%d", i)
        }
        close(ch)
    }()

    for v := range ch {
        fmt.Println("got:", v)
    }

    msg, ok := <-ch
    fmt.Printf("closed: %#v \tok=%v\n", msg, ok)

    // Sending to a closed channel will panic
    ch <- "Hi" // Panic
}
```

---

### 3.6 Sending on a Closed or Nil Channel

| Scenario | Behavior |
|-----------|-----------|
| Send on closed channel | **Panic** |
| Receive from closed channel | Returns zero value and `ok=false` |
| Send/Receive on nil channel | **Blocks forever** |

**Example:**
```go
var sendCh chan int // nil channel
sendCh <- 1         // blocks forever
```

---

### 3.7 The `select` Statement
`select` in Go is like a `switch` for channels. It lets you wait on multiple channel operations.

**Example:**
```go
func main() {
    sendCh := make(chan int)
    recvCh := make(chan int)

    go func() { sendCh <- 1 }()
    go func() { recvCh <- 2 }()

    select {
    case val := <-sendCh:
        fmt.Println("Received from sendCh:", val)
    case val := <-recvCh:
        fmt.Println("Received from recvCh:", val)
    }
}
```
**Key Points:**
- `select` picks the first available case.
- If multiple are ready, one is chosen **randomly**.
- A `default` case can be used to avoid blocking.

**Use Case:** Handling multiple goroutines or timeouts concurrently.

---

## 4. Avoiding Race Conditions with Channels
When multiple goroutines write to a channel, we use a `sync.WaitGroup` to synchronize them.

**Example:**
```go
func main() {
    c := make(chan int)
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        wg.Done()
    }()

    go func() {
        wg.Wait()
        close(c)
    }()

    for n := range c {
        fmt.Println(n)
    }
}
```

---

## 5. Context in Go

`context` in Go helps manage cancellation, timeouts, and deadlines across goroutines.

### Why Use Context?
- To **cancel** long-running goroutines.
- To set **timeouts** and **deadlines** for operations.
- To pass request-specific data through call chains.

### Example 1: Basic Context
```go
ctx := context.Background()
```
`context.Background()` is the root context — typically used at the top level.

### Example 2: Context with Timeout
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

select {
case <-time.After(5 * time.Second):
    fmt.Println("operation completed")
case <-ctx.Done():
    fmt.Println("timeout:", ctx.Err())
}
```
**Output:**
```
timeout: context deadline exceeded
```

### Example 3: Context with Cancel
```go
ctx, cancel := context.WithCancel(context.Background())

ch := make(chan int)

// worker goroutine
go func() {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("worker stopped")
            return
        default:
            fmt.Println("working...")
            time.Sleep(1 * time.Second)
        }
    }
}()

time.Sleep(3 * time.Second)
cancel() // stop worker
```

### Example 4: Context Values
```go
ctx := context.WithValue(context.Background(), "userID", 101)
fmt.Println("UserID:", ctx.Value("userID"))
```
**Note:** Context values should only be used for request-scoped data.

---

## 6. Real-World Example: Chat Application Using Channels

A channel-based chat server allows multiple clients to communicate in real time using TCP connections.

### Core Logic
```go
var (
    clients    = make(map[net.Conn]string)
    mutex      sync.Mutex
    messagesCh = make(chan string)
)
```

### Broadcaster Goroutine
```go
func broadcaster() {
    for msg := range messagesCh {
        mutex.Lock()
        for conn := range clients {
            fmt.Fprintln(conn, msg)
        }
        mutex.Unlock()
    }
}
```

### Client Handler
```go
func handleClient(conn net.Conn) {
    defer conn.Close()
    fmt.Fprint(conn, "Enter your name: ")
    nameInput := bufio.NewReader(conn)
    name, _ := nameInput.ReadString('\n')
    name = strings.TrimSpace(name)

    mutex.Lock()
    clients[conn] = name
    mutex.Unlock()

    messagesCh <- fmt.Sprintf("%s joined the chat!", name)

    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            mutex.Lock()
            delete(clients, conn)
            mutex.Unlock()
            messagesCh <- fmt.Sprintf("%s left the chat.", name)
            return
        }
        messagesCh <- fmt.Sprintf("[%s]: %s", name, strings.TrimSpace(message))
    }
}
```

**How It Works:**
- `messagesCh` broadcasts messages to all connected clients.
- The broadcaster listens continuously on `messagesCh`.
- When a user sends a message, it’s written into the channel and distributed.

---



## 8. Summary 
| Type | Description | Use Case |
|------|--------------|-----------|
| Unbuffered | No capacity, sender blocks until receiver ready | Synchronization |
| Buffered | Fixed capacity, allows async send/receive | Decoupling producer-consumer |
| Directional | Restricted to send or receive | Safe API design |
| Closed | Indicates no more data | Loop termination |

---

## 9. Exercises 

1. Write a program that sends numbers 1–10 through an unbuffered channel.
2. Use a buffered channel of capacity 5 to print values concurrently.
3. Demonstrate sending to a closed channel and handle the panic.
4. Use the `ok` variable to check if a channel is closed.
5. Write a `select` example that reads from two channels.
---
# Setup Guide for PostgreSQL, Docker Desktop, and PgAdmin4, Postman



## 1. Download PostgreSQL
- Visit the official PostgreSQL download page:  
  [PostgreSQL Downloads](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads)

## 2. Download Docker Desktop
- Get Docker Desktop from the official website:  
  [Docker Desktop Download](https://www.docker.com/products/docker-desktop/)

## 3. Download PgAdmin4
- Download PgAdmin4 for Windows:  
  [PgAdmin4 Download](https://www.pgadmin.org/download/pgadmin-4-windows/)

## 4. Sign Up for Docker Hub
- If you need a Docker Hub account, sign up here:  
  [Docker Hub Signup](https://app.docker.com/signup)

## 5. download Postman
 [Postman](https://www.postman.com/downloads/)

---

Make sure to follow the installation instructions on each respective website for a smooth setup.


