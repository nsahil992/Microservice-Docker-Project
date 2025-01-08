# Start with the Go base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o streaming-service .

# Expose the application port
EXPOSE 8082

# Run the application
CMD ["./streaming-service"]

