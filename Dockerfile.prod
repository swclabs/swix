# Stage build
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
# Copy migration files to /tmp
RUN cp -r pkg/migration /tmp
# Build the application
RUN go build -ldflags="-s -w" -o /bin/swipe ./cmd/${APP_MODULE}

# Remove all source code files
RUN rm -r *
# Copy migration files back to /app
RUN mkdir -p pkg && cp -r /tmp/migration/ pkg
# Clean up
RUN go clean -modcache

# Stage run
FROM alpine:edge as runner
WORKDIR /app

ENV HOST 0.0.0.0
ENV PORT 8000

COPY --from=builder /bin/swipe /bin/swipe
COPY --from=builder /app/pkg/migration /app/pkg/migration

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata