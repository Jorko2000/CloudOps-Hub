FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o operator ./cmd

FROM alpine

WORKDIR /app

COPY --from=builder /app/operator .

CMD ["./operator"]
