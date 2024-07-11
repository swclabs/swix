FROM golang:1.22-alpine AS builder

ARG APP_MODULE=""

# Move to working directory (/app).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the binary app.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN cp -r pkg/migration /tmp
RUN go build -ldflags="-s -w" -o /bin/swipe ./cmd/${APP_MODULE}

# Remove all source code files
RUN rm -r *
RUN mkdir -p pkg && cp -r /tmp/migration/ pkg
RUN go clean -modcache