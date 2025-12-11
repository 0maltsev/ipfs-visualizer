# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.25 AS build-stage

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /myapp ./cmd/app


# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /build

COPY --from=build-stage /myapp/ /myapp

ENV DOCKER_API_VERSION=1.41

EXPOSE 3001
RUN apk add libc6-compat

# USER nonroot:nonroot

ENTRYPOINT /myapp