# ---------- BUILD STAGE ----------
# Matching Go version with explicit version in go.mod
FROM golang:1.16-alpine AS builder

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /build/reader_service

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o reader .

# ---------- RUNTIME STAGE ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary and set ownership to root
COPY --from=builder /build/reader_service/reader /app/reader
RUN chown root:root /app/reader

# Use non-root runtime user
USER nonroot:nonroot

EXPOSE 5002

ENTRYPOINT ["/app/reader"]