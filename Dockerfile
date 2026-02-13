# Stage 1: Builder (aligned with Dockerfile.bootstrap; CGO_CFLAGS preserved via Makefile CGO_CFLAGS ?=)
FROM golang:1.24.0-bookworm AS builder

LABEL stage=builder

# Portable blst build: avoids SIGILL in blst_cgo_init when image runs on different CPU (e.g. build on Mac, run on Ubuntu 24)
ENV CGO_CFLAGS="-O -D__BLST_PORTABLE__ -std=gnu11"
ENV CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"
ENV CGO_ENABLED=1

# Install necessary build dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends make git build-essential

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,id=callisto_gomod,target=/go/pkg/mod \
    go mod download -x
COPY . .
# Makefile uses CGO_CFLAGS ?= so this ENV is not overwritten; use fresh cache so old blst objects aren't reused
RUN --mount=type=cache,id=callisto_gobuild_portable,target=/root/.cache/go-build \
    make build

# Stage 2: Final Runtime Image
FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup --system appgroup --gid 1001 && \
    adduser --system appuser --uid 1001 --ingroup appgroup

WORKDIR /home/appuser
COPY --from=builder /app/build/callisto /usr/local/bin/callisto
RUN chmod +x /usr/local/bin/callisto
USER appuser
ENTRYPOINT ["/usr/local/bin/callisto"]
CMD ["parse", "bootstrap", "start", "--home", "/callisto/.callisto"]