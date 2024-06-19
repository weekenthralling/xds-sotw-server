FROM golang:1.22 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o xds-server

FROM ubuntu:22.04

WORKDIR /app/

COPY --from=builder /build/xds-server .

CMD ["./xds-server"]