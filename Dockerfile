# Use buildx for multi-platform builds
# Build stage
FROM --platform=$BUILDPLATFORM golang:1.22.2-alpine AS builder
LABEL org.opencontainers.image.source="https://github.com/interlynk-io/sbomex"

RUN apk add --no-cache make git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build for multiple architectures
ARG TARGETOS TARGETARCH
RUN make build && chmod +x ./build/sbomex

# Final stage
FROM alpine:3.19
LABEL org.opencontainers.image.source="https://github.com/interlynk-io/sbomex"
LABEL org.opencontainers.image.description="Find & pull public SBOMs"
LABEL org.opencontainers.image.licenses=Apache-2.0

COPY --from=builder /app/build/sbomex /app/sbomex
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/app/sbomex"]
