## Build stage
FROM golang:1.17-alpine AS build-env
RUN apk update && apk add build-base && rm -rf /var/cache/apk/*
ADD ./* /go/src/github.com/erkexzcx/stoppropaganda/
WORKDIR /go/src/github.com/erkexzcx/stoppropaganda
RUN go build -ldflags="-s -w" -o stoppropaganda

## Create image
FROM alpine
RUN apk update && apk add bash ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /go/src/github.com/erkexzcx/stoppropaganda/stoppropaganda /app/
ENTRYPOINT ["/app/stoppropaganda"]
