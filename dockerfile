# Use Go builder
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Only copy go.mod (no go.sum yet)
COPY go.mod ./

# Download modules (will skip if no external modules)
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o docker-api-service .

# Final lightweight image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/docker-api-service .

EXPOSE 8005

CMD ["./docker-api-service"]
