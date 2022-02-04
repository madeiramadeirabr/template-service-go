# Core APIs Go Service Template

MadeiraMadeira boilerplate project to build scalable, testable and high performance Go microservices.

* Note env file are just required for development
## Stack
- Go
- Docker
- gRPC

## Featues
- REST and gRPC layers;
- Automatic pre-commit linting;
- `ConfigService`: easily manage environment variables;
- `Health` module: a gRPC and a REST endpoint that returns relevant information about the application status;
- `LogService`: easily manage application logs;
- Docker infrastructure with Docker Compose.


## Docker for local development

### Running the app

```bash
$ docker-compose up
```

## Local development

### Installation
```bash
$ go mod download
```

### Running the app
```bash
$ go run .\cmd\go-service-template\main.go
```

## Test

```bash
# unit tests
$ go test ./test/

# e2e tests
// TODO

# test coverage
// TODO
``` 
---
Squad Core APIs