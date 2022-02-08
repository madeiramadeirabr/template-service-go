# Core APIs Go Service Template

MadeiraMadeira boilerplate project to build scalable, testable and high performance Go microservices.

## Stack
- Go
- Docker
- gRPC

## Features
- [soon] REST and gRPC layers;
- Automatic pre-commit linting and testing;
- `configuration`: easily manage environment variables;
- [soon] `health_check` module: a gRPC and a REST endpoint that returns relevant information about the application status;
- [soon] `logger`: easily manage application logs following the [MMRFC1](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2317942893/MMRFC+1+-+Log) standards;
- Docker infrastructure with Docker Compose.

## Local development with Docker

### Setup
```bash
$ ./setup.sh
```

### Running the app

```bash
$ docker-compose -f docker-compose.dev.yml up
```

## Local development

### Setup
```bash
$ ./setup.sh
$ go mod download
```

### Running the app

```bash
$ go run ./cmd/go_service_template/main.go
```

## Test

```bash
# unit tests
$ go test ./...

# e2e tests
[soon]

# test coverage
[soon]
``` 

### Improve tests output

To see a more friendly test output, install `gotestsum` globally by running
```bash
$ go install gotest.tools/gotestsum@latest
```
And execute as
```bash
$ gotestsum --format testname 
```
at the project root directory.

---
Squad Core APIs