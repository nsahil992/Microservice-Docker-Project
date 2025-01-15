FROM golang:1.23-alpine

WORKDIR /app

# Copy only the Go module files first
COPY go.mod .
RUN go mod download

# Copy the application code
COPY main.go .

# Build the Go binary
RUN go build -o streaming-service .

# Expose the service port
EXPOSE 8083

# Run the application
CMD ["./streaming-service"]

