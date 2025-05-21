# Build stage: compile and build dependencies
FROM golang:1.22-alpine as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum separately for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o server ./cmd/main.go

# Final Image step
FROM alpine:3.19

# Create user first to run
RUN addgroup --system app && adduser --system --ingroup app app
WORKDIR /app
COPY --from=builder /app/server .
USER app

# Set the binary as the entrypoint
CMD ["./server"]