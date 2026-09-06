FROM golang:alpine AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /usr/local/bin/app ./cmd

FROM alpine:latest
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

RUN addgroup -S -g 10001 app
RUN adduser -S -D -H -u 10001 -G app app

USER app:app
ENTRYPOINT ["app"]
