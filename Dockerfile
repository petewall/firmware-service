FROM golang:1.19-alpine AS builder

WORKDIR /firmware-service

COPY . /firmware-service/

ARG GOOS=linux
ARG GOARCH=amd64

RUN apk add bash curl make && \
    make build

FROM alpine

WORKDIR /

COPY --from=builder /firmware-service/build/firmware-service /firmware-service

# ENV PORT=5050

ENTRYPOINT ["/firmware-service"]
