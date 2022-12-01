FROM golang:1.18-alpine AS build_base

RUN apk add --update --no-cache ca-certificates git
RUN apk add --update --no-cache ca-certificates bash
RUN apk add --update --no-cache ca-certificates build-base 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .

COPY go.sum .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN chmod +x ./unitTest.sh
RUN ./unitTest.sh

# Build the Go app
RUN go build .

# Expose port 9000 to the outside world
EXPOSE 9090

# Run the executable
CMD ["./go-cake"]