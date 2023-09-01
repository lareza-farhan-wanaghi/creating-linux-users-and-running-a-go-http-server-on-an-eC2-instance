# Use the latest official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download the dependencies specified in the go.mod file
RUN go mod download

# Build the Go application and generate an executable named 'main'
RUN go build -o main .

# Set the command to run when the container starts
CMD ["./main"]
