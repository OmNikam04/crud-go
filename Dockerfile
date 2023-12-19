FROM golang:latest

# Set the current working directory inside the container
WORKDIR /crud-go

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main /crud-go/cmd/app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air"]