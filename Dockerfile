FROM golang:1.23-alpine AS builder

# Move to working directory (/app).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# Build the application
RUN go build -ldflags="-s -w" -o /bin/swipe ./cmd

# Remove all source code files
RUN rm -r *
# Clean up
RUN go clean -modcache

# Stage run
FROM alpine:edge AS runner
WORKDIR /app

ENV HOST 0.0.0.0
ENV PORT 8000

COPY --from=builder /bin/swipe /bin/swipe

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata
