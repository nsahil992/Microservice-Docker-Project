# Microservice Docker Project: Authentication, Payment, and Streaming

This project demonstrates how to build a scalable web application using Microservices Architecture and Docker Compose. It includes three core services—Authentication, Payment, and Streaming—each running in its own Docker container. Docker Compose is used to orchestrate and scale these containers for easy management and seamless scaling.

### 📝 Project Overview

In this project, the application is divided into three services:

Authentication Service: Manages user login.
Payment Service: Handles payment processing.
Streaming Service: Manages video streaming functionality.
Each service is containerized using Docker. Docker Compose is then used to orchestrate the services, making it easier to manage multiple containers and scale them as needed.

### 🛠️ Technologies Used

Go (Golang) for building the microservices.
Docker for containerization.
Docker Compose for orchestrating the containers.
REST APIs for communication between services.
```
📁 Project Structure

├── auth-service
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── payment-service
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── streaming-service
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   └── go.sum
└── docker-compose.yml
```
### 🚀 Getting Started

To get the project up and running locally, follow these steps:

1. Clone the Repository

```
git clone https://github.com/nsahil992/Microservice-Docker-Project.git
cd Microservice-Docker-Project
```

2. Build and Run the Containers with Docker Compose
   Ensure that Docker and Docker Compose are installed on your machine.
```
Run the following command to build and start all services:
docker compose up --build -p
```
This will build the Docker images for each service and start the containers.

3. Access the Services
```
Authentication Service: http://localhost:8081/login
Payment Service: http://localhost:8082/pay
Streaming Service: http://localhost:8083/stream
```

