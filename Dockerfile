# Start with the builder stage
# Use a full-fledged Go environment to build the application
FROM golang:latest as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.* ./

# Download all dependencies
RUN go mod download

# Install swag using go install
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate Swagger documentation
RUN swag init

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start from a smaller base image for the runtime
FROM alpine:latest

# Set the working directory in the new image
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
