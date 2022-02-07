##
## Stage: Base
##
FROM golang:1.17-alpine AS base
WORKDIR /app
ADD . /app


##
## Stage: Development
##
FROM base AS development
RUN apk add curl \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
    && go mod vendor \
    && go mod download
ENTRYPOINT ["air", "-c", "./cmd/go-service-template/air.toml"]


##
## Stage: Build
##
FROM base AS builder
COPY . .
RUN go mod vendor \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -a -o go-service "./cmd/go-service-template"


##
## Stage: Production
##
FROM builder as production
WORKDIR /production
COPY --from=builder /app/go-service .
ENTRYPOINT ["/go-service"]
