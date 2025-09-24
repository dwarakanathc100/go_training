# Go Modules, Package Management, and Logging


------------------------------------------------------------------------

## 1. `go get` Command

The `go get` command is used to download and update dependencies in a Go
project.

-   **`go get -u ./...`**
    -   Updates all dependencies used in the current module and
        submodules to the latest **minor or patch versions**.
    -   Useful when you want to bring all libraries up to date without
        changing major versions that might break compatibility.
-   **`go get github.com/google/uuid@latest`**
    -   Fetches the latest available version of the `uuid` package
        directly from the repository.

    -   You can specify exact versions too:

        ``` bash
        go get github.com/google/uuid@v1.5.0
        ```

------------------------------------------------------------------------

## 2. Import Aliases in Go

Go allows you to use **aliases** when importing packages to avoid
conflicts or improve readability.

Example:

``` go
import (
    log "github.com/sirupsen/logrus"
    uuid "github.com/google/uuid"
)
```

-   Here, `log` and `uuid` act as **short names (aliases)** for the
    imported packages.
-   This avoids long references in the code and helps when multiple
    packages have the same name.

------------------------------------------------------------------------

## 3. `go.mod` File Explained

The `go.mod` file is used to track and manage dependencies for a Go
project.

Example:

``` go
module custom-uuid

go 1.21.6

require (
    github.com/google/uuid v1.5.0
)

replace github.com/google/uuid v1.5.0 => github.com/google/uuid v1.6.0

exclude github.com/google/uuid v1.1.0

retract [v1.0.0, v1.0.1]

toolchain go1.21.6
```

### Sections in `go.mod`:

-   **`go version`**
    -   Specifies the minimum Go version required to build the project.
-   **`require`**
    -   Declares dependencies and their versions that your module
        depends on.
-   **`replace`**
    -   Redirects a module to a different version or location.
    -   Example: Replace version `1.5.0` of `uuid` with version `1.6.0`.
-   **`exclude`**
    -   Explicitly prevents the use of certain versions of a dependency.
-   **`retract`**
    -   Marks versions as invalid (e.g., buggy releases). Developers are
        discouraged from using these.
-   **`toolchain`**
    -   Defines the Go toolchain version (e.g., `go1.21.6`) to ensure
        compatibility.

------------------------------------------------------------------------

## 4. `go build -mod=vendor`

-   The **vendor directory** contains a copy of all dependencies.

-   Running:

    ``` bash
    go build -mod=vendor
    ```

    forces Go to use the dependencies from the `vendor` folder instead
    of downloading them again.

-   This is useful for:

    -   Offline builds.
    -   Stable builds in CI/CD pipelines.

------------------------------------------------------------------------

## 5. Logging in Go

Logging is essential for debugging and monitoring applications. Using
`logrus`, you can log at different **levels**.

### Common Log Levels:

-   **Debug**
    -   For development and debugging details.

    -   Example:

        ``` go
        log.WithField("uuid", "12121").Debug("Debugging info")
        ```
-   **Info**
    -   General operational messages that show the flow of the
        application.
    -   Example: `log.Info("Consuming custom-uuid")`
-   **Warn**
    -   Warnings about potential issues but not errors.
    -   Example: `log.Warn("1 pod = 1 million customers")`
-   **Error**
    -   Errors that allow the application to continue running.
    -   Example: `log.Error("Database connection failed")`
-   **Fatal**
    -   Logs the message and then **terminates the program** with
        `os.Exit(1)`.
    -   Example: `log.Fatal("Fatal error occurred")`
-   **Panic**
    -   Logs the message and then panics (stack trace).
    -   Example: `log.Panic("Unexpected issue triggered panic")`

------------------------------------------------------------------------

### Logging Example Code:

``` go
package main

import (
    "fmt"
    gen "github.com/dwarakanathc100/uuid-generator/generator"
    log "github.com/sirupsen/logrus"
)

func main() {
    fmt.Println("Consuming custom-uuid!")
    uuid := gen.GenerateUUID()
    fmt.Println("Generated UUID:", uuid)

    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(log.DebugLevel)

    // Debug level logging
    log.WithField("uuid", uuid).Debug("New UUID generated!")

    // Info level logging
    log.Info("Consuming custom-uuid")

    // Warn level logging
    log.Warn("1 pod = 1 million customers")

    // Error level logging
    log.Error("Example error occurred")

    // Uncomment for fatal or panic demonstration
    // log.Fatal("Fatal error occurred")
    // log.Panic("Panic triggered")

    fmt.Println("custom-uuid generator:", gen.GenerateCusotmUUID())
}
```

------------------------------------------------------------------------

## 6. Environment-Specific Logging Example

-   **E0** -- Development environment logs.
-   **E1** -- Testing environment logs.
-   **E2 (QA)** -- Logs in QA testing.
-   **E3 (Staging)** -- Logs before production deployment.
-   **E3 (Production)** -- Logs in production.

This environment mapping helps in setting different log levels per
environment.

------------------------------------------------------------------------

# Exercises 

1.  Run `go get -u ./...` in your project and observe what changes in
    `go.mod` and `go.sum`.
2.  Add an alias for the `fmt` package and print text using the alias.
3.  Try using `replace` in your `go.mod` file to point a dependency to a
    different version.
4.  Implement a logger with `Info`, `Warn`, and `Error` messages.
5.  Explore the effect of `log.Fatal()` and `log.Panic()` in a small Go
    program.
