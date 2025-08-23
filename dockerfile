FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o docker-api-service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/docker-api-service .

EXPOSE 8005

CMD ["./docker-api-service"]
