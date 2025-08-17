# Build Golang code
FROM golang:1.24.5-bookworm AS buildergo
#RUN useradd -m builder
WORKDIR /app/go

ADD ./go.mod ./
RUN go mod download

COPY . .

RUN go build -a -o main .

# Use eclipse-temurin as the base image for Java
FROM ubuntu:24.04

# Install Go, Supervisor, and wget
RUN apt-get update && apt-get install -y \
    supervisor \
    sockperf \
    wget \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=buildergo /app/go/main ./main

ENTRYPOINT ["/app/main"]