# Use Golang base image
FROM golang:1.23-alpine

# Set working directory inside the container
WORKDIR /app

# Copy only the Go module files first
COPY go.mod .
RUN go mod download

# Copy the application code
COPY main.go .

# Build the service
RUN go build -o auth-service .

# Expose the port the service will run on
EXPOSE 8081

# Command to run the executable
CMD ["./auth-service"]

