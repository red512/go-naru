# Start from the official Go image, version 1.16
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY . .

# Build the Go binary
RUN go build -o naru ./cmd

# Set the environment variables for the container
ENV MY_APP_PORT=8080

# Expose port 8080
EXPOSE 8080

# Start the container with the built binary
CMD ["./naru", "server"]
