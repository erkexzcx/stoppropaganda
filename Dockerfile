## Build stage
FROM golang:1.17-alpine AS build-env
ADD . /app
WORKDIR /app
RUN env CGO_ENABLED=0 go build -ldflags="-s -w" -o stoppropaganda ./cmd/stoppropaganda/main.go

## Create image
FROM alpine
COPY --from=build-env /app/stoppropaganda /app/
ENTRYPOINT ["/app/stoppropaganda"]
