# Day 6 : Go Basics and Microservices Architecture

## 1. Go Playground
The **Go Playground** is an online tool provided by the Go community where you can:
- Write and execute Go code in your browser.
- Compile and run the code directly online.
- Share the code with anyone through a **share option** (generates a unique link).


---

## 2. `package main`
In Go, every program starts with a **package declaration**.

- `package main` is special because it defines a standalone executable program.
- Without `package main`, the Go compiler will not create an executable.

```go
package main
```

---

## 3. `func main()`
- The `main` function is the **entry point** of a Go program.
- A file should have **only one `main()` function**, otherwise the compiler throws an error.
- Execution begins from `func main()`.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

---

## 4. Comments in Go
- **Single-line comments** start with `//`
- **Multi-line comments** are enclosed within `/* ... */`

```go
// This is a single-line comment

/*
This is a 
multi-line comment
*/
```

---

## 5. `fmt` Package and Formatting
The `fmt` package is used for formatted I/O.

- `%s` → String
- `%d` → Integer
- `%v` → Default format
- `%T` → Type of variable

```go
name := "Kim"
age := 25
fmt.Printf("Name: %s, Age: %d
", name, age)
fmt.Printf("Value: %v, Type: %T
", name, name)
```

---

## 6. Microservices Architecture

Amazon can be explained with **microservices**. Each service has a single responsibility.

### Architecture
- **User Service** → Manages user authentication and data.
- **Search Service** → Handles product search requests.
- **Inventory Service** → Tracks stock availability.
- **Order Service** → Manages orders placed by users.
- **Payment Service** → Handles transactions securely.
- **Shipping Service** → Manages product delivery.
- **Tracking Service** → Provides shipment tracking updates.
- **Notification Service** → Sends updates via email/SMS/app.

graph TD
    A[Amazon] --> B["User service (DB)"]
    A --> C[Search service]
    C --> D[Inventory service]
    D --> E[Order service]
    E --> F[Payment service]
    E --> G[Shipping service]
    G --> H[Tracking service]
    E --> I[Notification service]
---

## 7. Real-Time Tools for Development
To build and manage microservices effectively, developers use tools like:
- **Docker** → Containerization of services.
- **Kubernetes** → Orchestration and scaling of containers.
- **Prometheus + Grafana** → Monitoring and visualization.
- **Jenkins/GitHub Actions** → CI/CD pipelines.
- **Postman** → API testing.

---

## 8. Communication Between Services
Different technologies help services talk to each other:

- **REST API** → Standard HTTP methods (`GET`, `POST`, `PUT`, `DELETE`).
- **gRPC** → High-performance RPC framework.
- **Kafka** → Event-driven communication.
- **RabbitMQ** → Message queue system.
- **GraphQL** → Flexible querying system....etc

---

## 9. Go Packages for Development
Some important packages in Go include:
- `fmt` → Formatted I/O.
- `log` → Logging errors, warnings, and info.
- `bufio` → Buffered I/O for reading/writing efficiently.
- `net/http` → For building web servers and APIs.
- `os` → For file and environment handling.

---
