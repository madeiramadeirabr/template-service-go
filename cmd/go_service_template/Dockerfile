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
RUN apk add curl \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
ENTRYPOINT ["air", "-c", "./cmd/go_service_template/air.toml"]


##
## Stage: Build
##
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-service "./cmd/go_service_template"


##
## Stage: Production
##
FROM alpine:latest as production
WORKDIR /production
COPY --from=builder /app/go-service .
ENTRYPOINT ["/production/go-service"]
