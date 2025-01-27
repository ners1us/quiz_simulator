FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o quiz_simulator .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/quiz_simulator .
