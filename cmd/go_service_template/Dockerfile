
##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /go-service "./cmd/go_service_template" 

##
## Deploy
##
FROM build as run

WORKDIR /

COPY --from=build /go-service /go-service

EXPOSE 8080

ENTRYPOINT ["/go-service"]