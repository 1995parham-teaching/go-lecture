# Go Lecture

## Introduction

Learn to develop cloud-native programs with Go!

## Review Me

One of the main steps in learning new language and its best practices is reviewing
writeen projects:

- <https://github.com/1995parham/koochooloo>:

  - In the first step, you need to review the project structure and find out how modules are related
  - Then we continue with running the docker-compose to have the requirements
  - And in the final step, we lunch the application and trying it with curl based on its swagger
  - This project use zap as a logger and pass it into its modules also each module
    has its metrics based on [otel](https://github.com/open-telemetry/).

- <https://github.com/1995parham/fandogh>:

  - This example containing the migration and how we store things on database

- <https://github.com/1995parham/saf>:

  - This example shows tracing in action with nats as a message queue
  - also, it contains a pyroscope to see how replacing echo with gofiber increase the performance
  - this project has helm and after knowing Kubernetes basis they can lunch it on the cloud with its helm

- <https://github.com/1995parham/k1s>:
  - in the first step, they review the server structure. the server is stateless and only returns simple responses.
  - they lunch it on the cloud with its manifests
  - using server and ingress to send requests and see how they distributed between instances
  - they also see how we can mount configuration on Kubernetes with configmap.
