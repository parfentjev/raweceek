FROM golang:1.26.2-alpine AS builder
WORKDIR /app/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o /app/out/raweceek .

FROM alpine:3.23.4 AS release
COPY --from=builder /app/out/raweceek /usr/local/bin/raweceek
CMD ["raweceek"]
