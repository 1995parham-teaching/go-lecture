# Go Lecture

## Introduction

> Learn to develop cloud-native programs with Go!

The Golang programming language is very similar to the C programming language, and its purpose is to reduce complexity in program development.
This language is widely used to implement web servers, applications, and container management tools.
Among the tools that have been developed with this language are:

- [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Docker](https://github.com/moby/moby)
- [NATS](https://github.com/nats-io/nats-server)
- [Prometheus](https://github.com/prometheus/prometheus)

In the last few years, this language has opened its place in Iranian companies and it is used for the development of backend services.

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
