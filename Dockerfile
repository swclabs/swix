FROM golang:1.21-alpine AS builder

ARG APP_MODULE=""

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the binary app.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /bin/swipe ./cmd/${APP_MODULE}

# Remove all source code files
RUN rm -r *
RUN go clean -modcache