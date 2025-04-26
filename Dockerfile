# Use official Go image
FROM golang:1.19-alpine

# Set working directory
WORKDIR /app

# Copy go.mod (no go.sum needed)
COPY go.mod ./

# Download dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the app
CMD ["./main"]
