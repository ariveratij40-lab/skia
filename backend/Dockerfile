# ---------- BUILD ----------
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

# 1) Copiar go.mod + go.sum primero (requisito B)
COPY go.mod go.sum ./

# 2) Descargar deps con cache estable
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 3) Copiar el resto del código
COPY . .

# 4) Build con cache de compilación
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
