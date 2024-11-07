# # Dockerfile
# FROM golang:1.20-alpine

# # Set the working directory inside the container
# WORKDIR /app

# # Copy everything into the container
# COPY . .

# # Install dependencies and compile the Go program
# RUN go mod download

FROM golang:1.23.0

WORKDIR /app

COPY . .

RUN go mod download
# Run the application
CMD ["go", "run", "cmd/app/main.go"]
