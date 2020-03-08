# Building server
FROM golang:1.14 AS builder

WORKDIR /srv/app
RUN make deps
RUN make


# Production-ready image of server
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /srv/app
COPY --from=builder /srv/app/server .

CMD ["./server"]