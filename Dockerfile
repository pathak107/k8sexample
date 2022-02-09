# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /k8Example

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /k8Example /k8Example

EXPOSE 3000
EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/k8Example", "--server=grpc"]