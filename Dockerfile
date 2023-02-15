# Start with the official Go image
FROM golang:1.19

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any necessary dependencies
RUN go mod download

# Build the app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Start the app
CMD ["./main"]