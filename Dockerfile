## Build stage
FROM golang:1.17-alpine AS build-env
ADD ./* /go/src/github.com/erkexzcx/stoppropaganda/
WORKDIR /go/src/github.com/erkexzcx/stoppropaganda
RUN env CGO_ENABLED=0 go build -ldflags="-s -w" -o stoppropaganda

## Create image
FROM alpine
COPY --from=build-env /go/src/github.com/erkexzcx/stoppropaganda/stoppropaganda /app/
ENTRYPOINT ["/app/stoppropaganda"]
