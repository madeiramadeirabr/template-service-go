##
## Build
##
FROM golang:1.17-alpine AS build

RUN apk add curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app
ADD . /app

RUN go mod download

EXPOSE 8080

ENTRYPOINT ["air", "-c", "./cmd/go-service-template/air.toml"]