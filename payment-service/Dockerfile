FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .

RUN go build -o payment-service .

EXPOSE 8082

CMD ["./payment-service"]

