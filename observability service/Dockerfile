FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o observability-service ./cmd/server

FROM alpine

WORKDIR /app

COPY --from=builder /app/observability-service .

EXPOSE 8083

CMD ["./observability-service"]
