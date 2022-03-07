## Build stage
FROM --platform=$BUILDPLATFORM golang:1.17-alpine AS build-env
ADD . /app
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o stoppropaganda ./cmd/stoppropaganda/main.go

## Create image
FROM scratch
COPY --from=build-env /app/stoppropaganda /
ENTRYPOINT ["/stoppropaganda"]
