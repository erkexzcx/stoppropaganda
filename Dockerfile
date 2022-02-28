## Build stage
FROM golang:1.17-alpine AS build-env
RUN go get github.com/peterbourgon/ff/v3
RUN go get github.com/miekg/dns
ADD ./* /go/src/github.com/erkexzcx/stoppropaganda/
WORKDIR /go/src/github.com/erkexzcx/stoppropaganda
RUN env CGO_ENABLED=0 go build -ldflags="-s -w" -o stoppropaganda cmd/stoppropaganda/main.go

## Create image
FROM alpine
COPY --from=build-env /go/src/github.com/erkexzcx/stoppropaganda/stoppropaganda /app/
ENTRYPOINT ["/app/stoppropaganda"]
