#!/bin/bash

# Before using this script, perform "docker login" :)

VERSION=0.0.31
PLATFORM=linux/amd64,linux/arm64,linux/ppc64le,linux/386,linux/arm/v7

docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker buildx create --use --name multiarch --platform ${PLATFORM} || true
docker buildx build --push --platform ${PLATFORM} --tag erikmnkl/stoppropaganda:${VERSION} --tag erikmnkl/stoppropaganda:latest .
