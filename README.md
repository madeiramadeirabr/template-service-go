# Core APIs Go Service Template

MadeiraMadeira boilerplate project to build scalable, testable and high performance Go microservices.

## Table of Contents
- [Stack](#stack)
- [Features](#features)
- [Hands On](#hands-on)
    - [Development](#development)
    - [Test](#test)
    - [Release](#release)

## Stack
- Go
- Docker
- gRPC

## Features
- [soon] REST and gRPC layers;
- [soon] Automatic pre-commit linting and testing;
- `configuration`: easily manage environment variables;
- [soon] `health_check` module: a gRPC and a REST endpoint that returns relevant information about the application status;
- [soon] `logger`: easily manage application logs following the [MMRFC1](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2317942893/MMRFC+1+-+Log) standards;
- Docker infrastructure with Docker Compose.



## Hands On

### Development

#### Using Docker
```bash
$ ./setup.sh
$ docker-compose -f docker-compose.dev.yml up
```

#### Without Docker
```bash
$ ./setup.sh
$ go mod download
$ go run ./cmd/SERVICE_NAME/main.go

#eg:
$ go run ./cmd/go_service_template/main.go
```

### Test
```bash
# unit tests
$ go test ./...

# e2e tests
[soon]

# test coverage
[soon]
```
#### Improve tests output

To see a more friendly test output, install `gotestsum` globally by running
```bash
$ go install gotest.tools/gotestsum@latest
```
And execute as below at the project root directory.
```bash
$ gotestsum --format testname 
```

### Release
```bash
$ docker build \
    --target release \
    -t SERVICE-NAME:VERSION \
    -f cmd/SERVICE_NAME/Dockerfile .
```
---
Squad Core APIs