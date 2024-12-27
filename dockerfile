FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /root
COPY . .
RUN go build -o /usr/local/bin/geo main.go

FROM alpine:3.21
COPY --from=builder /usr/local/bin/geo /usr/local/bin/geo
WORKDIR /
ENTRYPOINT ["geo"]