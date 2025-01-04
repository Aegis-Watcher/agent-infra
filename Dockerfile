# Stage 1: Builder
FROM golang:1.23.4-alpine3.21 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app:
#   - Output binary: 'agent-infra'
#   - -trimpath removes filesystem paths from the binary
#   - -ldflags="-s -w" strips debugging info to reduce size
RUN go build -o agent-infra -trimpath -ldflags="-s -w" ./cmd/main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/agent-infra .

# Run the binary
CMD ["./agent-infra"]

