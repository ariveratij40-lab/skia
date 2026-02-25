# ---------- BUILD ----------
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

# Copy go.mod + go.sum first
COPY go.mod go.sum ./

# Download deps with stable cache
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest
COPY . .

# Build
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o skia-api ./cmd/api

# ---------- RUNTIME ----------
FROM gcr.io/distroless/base-debian12
WORKDIR /app

ENV GIN_MODE=release

COPY --from=builder /app/skia-api /app/skia-api

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/skia-api"]
