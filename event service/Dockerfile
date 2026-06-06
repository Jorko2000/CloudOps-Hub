FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o event-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/event-service .

EXPOSE 50051

CMD ["./event-service"]
