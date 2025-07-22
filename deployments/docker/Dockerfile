FROM golang:1.23-alpine

# Install CompileDaemon for live reload
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Add Go bin to PATH
ENV PATH=$PATH:/go/bin

# Expose port
EXPOSE 8080

# For development with live reload
CMD ["CompileDaemon", "-directory=.", "-build=go build -o /tmp/server ./cmd/server", "-command=/tmp/server"]