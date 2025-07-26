# Docs: https://docs.docker.com/reference/dockerfile/
FROM golang:1.24.5-alpine3.22 AS base

# Builder stage
# =============================================================================
# Create a builder stage based on the "base" image
FROM base AS builder
# Move to working directory /build
WORKDIR /build
# Update the package list and install make
RUN apk update \
    && apk add --no-cache make \
    && apk add --no-cache upx
# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./
# Install dependencies
RUN go mod download
# Copy the entire source code into the container
COPY . .
# Build the application
# Turn off CGO to ensure static binaries
RUN make prod

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM base AS production
# Move to working directory /prod
WORKDIR /prod
# Update the package list and install curl
RUN apk update && apk add --no-cache curl
# Copy binary from builder stage
COPY --from=builder /build/bin/todo-backend ./
COPY --from=builder /build/.env ./
# Document the port that may need to be published
EXPOSE 8081
# Start the application
CMD ["/prod/todo-backend"]
