# Use the official Golang base image
FROM golang:1.20-alpine

# Set the working directory within the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Download dependencies
RUN go mod init twitter-handle-checker
RUN go mod tidy

# Build the Go application
RUN go build -o main

# Run the compiled Go binary
CMD ["./main"]
