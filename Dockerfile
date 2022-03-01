## Build stage
FROM golang:1.17-alpine AS build-env
ADD . /app
WORKDIR /app
RUN env CGO_ENABLED=0 go build -ldflags="-s -w" -o stoppropaganda ./cmd/stoppropaganda/main.go

# prepare rootfs for the runtime
WORKDIR /tmp/rootfs
RUN set -x \
    && mkdir -p ./etc \
    && echo 'stoppropaganda:x:10001:10001::/nonexistent:/sbin/nologin' > ./etc/passwd \
    && echo 'stoppropaganda:x:10001:' > ./etc/group \
    && mv /app/stoppropaganda ./stoppropaganda

## Create image
FROM scratch

LABEL \
    # Docs: <https://github.com/opencontainers/image-spec/blob/master/annotations.md>
    org.opencontainers.image.title="stoppropaganda" \
    org.opencontainers.image.url="https://github.com/erkexzcx/stoppropaganda" \
    org.opencontainers.image.source="https://github.com/erkexzcx/stoppropaganda" \
    org.opencontainers.image.vendor="erkexzcx"

# Import from builder
COPY --from=build-env /tmp/rootfs /

# Use an unprivileged user
USER stoppropaganda:stoppropaganda

ENTRYPOINT ["/stoppropaganda"]
