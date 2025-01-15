# Use Golang base image
FROM golang:1.23-alpine

# Set working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY main.go .

# Build the service
RUN go build -o auth-service .

# Expose the port the service will run on
EXPOSE 8081

# Command to run the executable
CMD ["./auth-service"]

