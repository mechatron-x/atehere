# Top level arguments
ARG GO_VERSION=1.23.2
ARG BUILD_OS=bookworm
ARG OS=alpine:3

# Build api stage
FROM golang:${GO_VERSION}-${BUILD_OS} AS build

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

RUN mkdir -p /usr/local/bin /var/lib/atehere/logs /var/lib/atehere/static

COPY --from=build /app/bin /usr/local/bin

CMD [ "atehere" ]