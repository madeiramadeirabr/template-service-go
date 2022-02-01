
##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-service

##
## Deploy
##
FROM build as run

WORKDIR /

COPY --from=build /go-service /go-service

EXPOSE 8080

ENTRYPOINT ["/go-service"]