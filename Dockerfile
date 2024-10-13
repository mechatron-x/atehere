# Top level arguments
ARG GO_VERSION=1.23.2
ARG BUILD_OS=bookworm
ARG OS=alpine:3

# Build api stage
FROM golang:${GO_VERSION}-${BUILD_OS} AS build

ARG GOOS=linux
ARG GOARCH=arm64
ARG CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
    go mod tidy && \
    go mod verify

COPY . ./

RUN go build -o bin/ -v -x ./cmd/atehere

# Publish stage
FROM ${OS}

RUN apk update && \
    apk upgrade

RUN mkdir -p /etc/atehere

WORKDIR /usr/local/bin

COPY --from=build /app/bin ./

CMD [ "atehere" ]