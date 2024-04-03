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
- Errors
- Concurrency and channels
- `select`
- `go mod` and using packages
- An overview of advanced features
- Introduction to HTTP protocol
- HTTP server implementation
- Settings management paragraph
- Metric, Log and Tracing
- connection with the database using PostgreSQL and GORM
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

### Echo

0. Say hello to Echo
1. HTTP Handlers
2. Request Binding
3. Path Parameters
4. Query Strings

## Continue your journey

You can visit [go101](https://github.com/1995parham-learning/go101) which contains
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

- <https://github.com/1995parham/k1s>:
  - In the first step, we review the server structure. The server is stateless and only returns simple responses.
  - We it on the cloud with its manifests
  - using server and ingress to send requests and see how they distributed between instances
  - We also see how we can mount configuration on Kubernetes with [configmap](https://kubernetes.io/docs/concepts/configuration/configmap/).
