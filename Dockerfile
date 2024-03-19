# Use a base image that supports ARM64 architecture for Apple Silicon (local machines)
FROM --platform=linux/arm64 ubuntu:20.04

# Install dependencies
RUN apt-get update && apt-get install -y \
    bpfcc-tools \
    linux-headers-generic \
    clang \
    llvm \
    libelf-dev \
    gcc \
    iproute2 \
    git \
    curl \
    ca-certificates \
    && apt-get clean

# Install Go for ARM64
RUN curl -OL https://golang.org/dl/go1.16.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go1.16.linux-arm64.tar.gz \
    && rm go1.16.linux-arm64.tar.gz

# Set environment variables for Go
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Copy your Go application source code and module files into the container
COPY src /go/src/go_user_agent

# Set the working directory to the Go user agent script's location
WORKDIR /go/src/go_user_agent/

# Initialize Go modules and download dependencies
RUN go get github.com/joho/godotenv@latest

# Build the Go user agent script
RUN go build -o go_user_agent

# Command to run the user agent script
CMD ["./go_user_agent"]