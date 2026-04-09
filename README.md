<h1 align="center">🐼 Learning Go 🤓</h1>
<h6 align="center">Learn to develop cloud-native programs with Go!</h6>

<p align="center">
  <img alt="banner" src="./.github/assets/banner.jpg" height="200px" />
  <br />
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/1995parham-teaching/go-lecture/build.yaml?style=for-the-badge">
</p>

## Introduction

The Golang programming language is very similar to the C programming language,
and its purpose is to reduce complexity in program development.
This language is widely used to implement web servers, applications, and container management tools.
Among the tools that have been developed with this language are:

- [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Docker](https://github.com/moby/moby)
- [NATS](https://github.com/nats-io/nats-server)
- [Prometheus](https://github.com/prometheus/prometheus)

In the last few years, this language has opened its place in Iranian companies,
and it is used for the development of backend services.

> [!NOTE]
> C background is **required** for learning Go.

> [!IMPORTANT]
> The examples in this repository target **Go 1.26** and make use of features
> introduced in that release. See the
> [Go 1.26 features used in this repo](#go-126-features-used-in-this-repo)
> section below for the highlights.

## Outline (As a separate course)

- History
- Variables and constants
- Calculation
- Conditions
- Loops
- Functions
- Strings
- Arrays and slices
- `map`
- `struct`
- `interface`
- Pointers
- Errors (including the generic `errors.AsType[E]` from Go 1.26)
- Concurrency and channels (`sync.WaitGroup.Go`, `sync.OnceValue`)
- `select`
- Generics (including self-referential type parameters from Go 1.26)
- `go mod` and using packages
- An overview of advanced features
- Introduction to HTTP protocol
- HTTP server implementation with `net/http`'s pattern-based routing
  and structured logging via `log/slog` (`slog.NewMultiHandler`)
- Graceful shutdown with `signal.NotifyContext` cancel causes
- Settings management
- Metric, Log and Tracing
- Connection with the database using PostgreSQL and GORM
- Introduction to Docker and containerization

At the beginning of the course, an introduction to the Go language is made and students implement simple programs with it.
Since the implementation of web servers is one of the important uses of the Go programming language,
we will review the structure of the HTTP protocol and then implement a simple web server in Go.
In this implementation, we try to familiarize ourselves with the structure of large programs created in Go and review details such as
Configuration or Metrics, which are of great value in real systems.
Finally, a MongoDB and PostgreSQL databases are added to this web server,
the purpose of which is to familiarize students with database interfaces in the Go.

## Outline (As a part of internet engineering course)

Using Go for design and implementing servers contains two major steps.
First is about learning Go itself and another step is learning an HTTP framework (here we go with Echo).
Reviewing these source codes are useful for learning Go but there aren't enough.

### Go

0. Hello World
1. Constants and Variables
2. Calculation
3. Conditions
4. Loops
5. Strings
6. Arrays
7. Slices
8. Arrays and Slices
9. Maps
10. Structs
11. Interfaces
12. Pointers
13. Structs with Pointers
14. `strconv`
15. Function with multiple-return
16. Errors
17. Concurrency
18. Function Type
19. Channels
20. Pipelines
21. Select
22. JSON
23. `go.mod`
24. Packages
25. `defer`
26. `variadic`
27. `regex`
28. `sync.Once` / `sync.OnceValue`
29. `panic`
30. UTF-8
31. The `nil` interface trap
32. Type aliases vs. type definitions
33. Self-referential generics (Go 1.26)

### Go 1.26 features used in this repo

This repository has been refreshed to use **Go 1.26** language and library
features. The most notable ones — and where to find a runnable example for
each — are:

| Feature                                                        | Where                                                                            |
| -------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `new(expr)` builtin: pointer + initial value in one step       | [`22-json/main.go`](./22-json/main.go)                                           |
| `errors.AsType[E]` — generic, type-safe `errors.As`            | [`16-errors/main.go`](./16-errors/main.go)                                       |
| Self-referential generic type parameters (`Adder[A Adder[A]]`) | [`33-generics-self-referential/main.go`](./33-generics-self-referential/main.go) |
| `slog.NewMultiHandler` for log fan-out                         | [`httpserver/main.go`](./httpserver/main.go)                                     |
| `signal.NotifyContext` + `context.Cause` for graceful shutdown | [`httpserver/main.go`](./httpserver/main.go)                                     |

The repository also benefits from earlier-but-still-modern features that are
worth highlighting in class:

- `sync.WaitGroup.Go(func)` instead of `Add` / `go` / `defer Done`
  (see [`17-concurrency/main.go`](./17-concurrency/main.go) and
  [`20-pipeline/main.go`](./20-pipeline/main.go))
- `sync.OnceValue` instead of `sync.Once` + sentinel field
  (see [`28-once/main.go`](./28-once/main.go))
- `for range N` instead of the C-style `for i := 0; i < N; i++`
  (see [`09-map`](./09-map/main.go), [`19-channels-1`](./19-channels-1/main.go),
  [`20-pipeline`](./20-pipeline/main.go), [`23-sharding`](./23-sharding/main.go))
- `math/rand/v2` (see [`21-select/main.go`](./21-select/main.go))
- `net/http` pattern-based routing — `mux.HandleFunc("GET /hello/{username}", …)`
  (see [`httpserver/main.go`](./httpserver/main.go))
- `mime.ParseMediaType` for tolerant `Content-Type` checks
  (see [`httpserver/handler/hello.go`](./httpserver/handler/hello.go))

### Keeping the codebase modern

Go 1.26 ships a much expanded `go fix` with a suite of _modernizer_ analyzers
that automatically rewrite older idioms (`interface{}` → `any`,
`for i := 0; i < n; i++` → `for range n`, `WaitGroup.Add`/`Done` → `WaitGroup.Go`,
and many more). To run all of them across this module:

```sh
just modernize
# or, equivalently:
go fix ./...
```

## Continue your journey 🧳

You can visit [Go101](https://github.com/1995parham-learning/go101) which contains
some more advance concepts of Golang.

## Review Me

One of the main steps in learning new language and its best practices is reviewing
written projects:

- <https://github.com/1995parham/koochooloo>:
  - In the first step, you need to review the project structure and find out how modules are related
  - Then we continue with running the docker-compose to have the requirements
  - And in the final step, we lunch the application and trying it with curl based on its swagger
  - This project use zap as a logger and pass it into its modules also each module
    has its metrics based on [otel](https://github.com/open-telemetry/).

- <https://github.com/1995parham/fandogh>:
  - This example containing the migration and how we store things on the [MongoDB](https://www.mongodb.com/) database.

- <https://github.com/1995parham/saf>:
  - This example shows tracing in action with [NATS](https://nats.io/) as a message queue.
  - Also, We can use profiler to see how replacing [Echo](https://echo.labstack.com/) with [GoFiber](https://gofiber.io/) increase the performance.
  - This project has Helm chart and after knowing Kubernetes basis we can lunch it on the cloud with its Helm.

- <https://github.com/1995parham-teaching/k1s>:
  - In the first step, we review the server structure. The server is stateless and only returns simple responses.
  - We it on the cloud with its manifests
  - using server and ingress to send requests and see how they distributed between instances
  - We also see how we can mount configuration on Kubernetes with [configmap](https://kubernetes.io/docs/concepts/configuration/configmap/).
