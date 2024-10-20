FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY main.go /app/main.go
COPY internal/ /app/internal/

RUN CGO_ENABLED=0 go build -o /usr/local/bin/ekdo .

FROM gcr.io/distroless/static AS static

LABEL maintainer="omissis"
# LABEL org.opencontainers.image.created
LABEL org.opencontainers.image.authors="omissis"
LABEL org.opencontainers.image.url="https://github.com/mecha-hq/ekdo"
LABEL org.opencontainers.image.documentation="https://github.com/mecha-hq/ekdo"
LABEL org.opencontainers.image.source="https://github.com/mecha-hq/ekdo"
# LABEL org.opencontainers.image.version
# LABEL org.opencontainers.image.revision
LABEL org.opencontainers.image.vendor="Mecha CI"
# LABEL org.opencontainers.image.licenses
# LABEL org.opencontainers.image.ref.name
LABEL org.opencontainers.image.title="ekdo"
LABEL org.opencontainers.image.description="A simple CLI tool to render image scan results to HTML."
# LABEL org.opencontainers.image.base.digest
# LABEL org.opencontainers.image.base.name

USER nonroot:nonroot

ENTRYPOINT ["/ekdo"]

FROM static AS dockerbuild

COPY --from=builder /usr/local/bin/ekdo /ekdo

FROM static AS goreleaser

COPY ekdo /ekdo
