# ---------- BUILD STAGE ----------
# Matching Go version with explicit version in go.mod
FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /build/writer_service

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o writer ./cmd/main.go

# ---------- RUNTIME STAGE ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary with non-root ownership for runtime
COPY --from=builder --chown=nonroot:nonroot /build/writer_service/writer /app/writer

# Use non-root runtime user
USER nonroot:nonroot

EXPOSE 5003

ENTRYPOINT ["/app/writer"]