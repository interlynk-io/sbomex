FROM golang:1.20-alpine AS builder
LABEL org.opencontainers.image.source="https://github.com/interlynk-io/sbomex"

RUN apk add --no-cache make git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make ; make build

FROM scratch
LABEL org.opencontainers.image.source="https://github.com/interlynk-io/sbomex"
LABEL org.opencontainers.image.description="SBOM Explorer - Find and pull public SBOMs"
LABEL org.opencontainers.image.licenses=Apache-2.0

COPY --from=builder /app/build/sbomex /app/sbomex
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/app/sbomex"]