# Step 1: Build the Go application
FROM golang:1.22.3-alpine AS base

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Step 3: Build the main application
FROM base AS build

ENV DB_NAME=sea_salon
ENV DB_USER=root
ENV DB_PASS=password
ENV DB_HOST=127.0.0.1
ENV MAX_RETRIES=5

# Copy the rest of the application code
COPY . .


# Build the application
RUN go build -o main .

# Step 4: Run the Go application
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=build /app/main .

# Copy status.json
COPY status.json ./

# Expose port 8080
EXPOSE 3000

# Command to run the application
CMD ["./main"]

