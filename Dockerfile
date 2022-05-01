FROM golang:1.17-alpine AS builder

WORKDIR /firmware-service

COPY . /firmware-service/

ARG GOOS=linux
ARG GOARCH=amd64

RUN go mod download
RUN go build -o build/firmware-service github.com/petewall/firmware-service/v2

FROM alpine

WORKDIR /

COPY --from=builder /firmware-service/build/firmware-service /firmware-service

# ENV PORT=5050
# ENV DB_HOST=
# ENV DB_PORT=6379

ENTRYPOINT ["/firmware-service"]
