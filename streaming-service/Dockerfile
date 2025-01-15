FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .

RUN go build -o streaming-service .

EXPOSE 8083

CMD ["./streaming-service"]

