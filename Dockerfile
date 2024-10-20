FROM --platform=$BUILDPLATFORM golang:1.23.2-alpine3.20 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY main.go .
COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETPLATFORM go build -o /usr/local/bin/ekdo .

FROM gcr.io/distroless/static

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

COPY --from=builder /usr/local/bin/ekdo /ekdo

ENTRYPOINT ["/ekdo"]
