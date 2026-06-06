FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o provisioning-service ./cmd/server

EXPOSE 8082

CMD ["./provisioning-service"]
