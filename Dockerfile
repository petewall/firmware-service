FROM golang:1.18-alpine AS builder

WORKDIR /firmware-service

COPY . /firmware-service/

ARG GOOS=linux
ARG GOARCH=amd64

RUN apk add git make && make build

FROM alpine

WORKDIR /

COPY --from=builder /firmware-service/build/firmware-service /firmware-service

# ENV PORT=5050
# ENV DB_HOST=
# ENV DB_PORT=6379

ENTRYPOINT ["/firmware-service"]
