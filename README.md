# Core APIs Go Service Template

MadeiraMadeira boilerplate project to build scalable, testable and high performance Go microservices.

## Table of Contents
- [Stack](#stack)
- [Features](#features)
- [Getting Started](#getting-started-with-boilerplate)
- [Hands On](#hands-on)
    - [Development](#development)
    - [Test](#test)
    - [Release](#release)

## Stack
- [Go](https://go.dev)
- [Docker](https://www.docker.com)

## Features
- Automatic pre-commit code analysis and lint with `golint`, `govet` and `goimports`;
- Automatic pre-push testing;
- `configuration`: easily manage environment variables and app configuration;
  - it reads a `.env` file when its available, and load system variables when it's not;
  - it also validates the available variables based on the `AppConfig` struct;
- `healthcheck`: REST endpoint that returns relevant information about the application status;
- `logger`: easily manage application logs following the [MMRFC1](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2317942893/MMRFC+1+-+Log) standards;
- Docker infrastructure with Docker Compose.


## Getting started with boilerplate

Download and extract the project, then:
```bash
$ mv go-service-template-production {your_project_name}
$ git init
$ git remote add origin https://github.com/{user}/{repo}.git
```

> Hint: use `$ git remote -v` to verify new remote


## Hands On

### Development

#### Docker setup
```bash
$ ./tools/setup.sh
$ docker-compose -f docker-compose.yml up
```

#### Local setup
```bash
$ ./tools/setup.sh
$ go mod download
$ go run ./cmd/SERVICE_NAME/main.go

# eg:
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

To see a more friendly test output, install [gotestsum](https://github.com/gotestyourself/gotestsum) globally by running
```bash
$ go install gotest.tools/gotestsum@latest
```
And execute at the project root directory:
```bash
$ gotestsum --format testname 
```

### Release
```bash
$ docker build \
    --target release \
    --build-arg SERVICE_PATH=_SERVICE_NAME_ \
    -t _SERVICE_NAME_:VERSION \
    -f docker/Dockerfile .

# eg:
$ docker build \
    --target release \
    --build-arg SERVICE_PATH=go_service_template \
    -t go-service-template-production:latest \
    -f docker/Dockerfile .
```
---
Squad Core APIs • [MadeiraMadeira](https://www.madeiramadeira.com.br)
