# Service folder name inside ./cmd
ARG SERVICE_PATH

##
## Stage: Base
##
FROM golang:1.17-alpine AS base
WORKDIR /app
ADD . /app
RUN go mod vendor \
    && go mod download


##
## Stage: Development
##
FROM base AS development
ARG SERVICE_PATH
RUN apk add curl \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
    && sed "s/{SERVICE_PATH}/${SERVICE_PATH}/g" ./docker/air.toml > /root/.air.toml
ENTRYPOINT ["air", "-c", "/root/.air.toml"]


##
## Stage: Build
##
FROM base AS builder
ARG SERVICE_PATH
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-service "./cmd/${SERVICE_PATH}"


##
## Stage: Release
##
FROM alpine:latest as release
WORKDIR /release
COPY --from=builder /app/go-service .
ENTRYPOINT ["/release/go-service"]
