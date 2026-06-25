FROM golang:1.23-alpine AS build-download
WORKDIR /build
COPY download.go .
RUN go mod init osah && go build -o download .

FROM golang:1.23-alpine AS build-parse
WORKDIR /build
COPY parse.go .
RUN go mod init osah && go build -o parse .

FROM alpine:3.20 AS download
COPY --from=build-download /build/download /usr/local/bin/download
WORKDIR /data
ENTRYPOINT ["download"]

FROM alpine:3.20 AS parse
COPY --from=build-parse /build/parse /usr/local/bin/parse
WORKDIR /data
ENTRYPOINT ["parse"]
