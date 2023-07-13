# Use the official Go image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the files into the container
COPY . .


RUN go mod init main.go
# Build the Go application
RUN go build -o concurrent .

# Set the command to run the application with the flags
ENTRYPOINT ["./concurrent", "-file", "testcur.csv", "-concurrent", "false"]


