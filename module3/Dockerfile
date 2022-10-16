FROM golang:1.16 as builder
RUN mkdir -p /go/src/test
WORKDIR /go/src/test
COPY httpServer.go go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -o http_server .

FROM alpine:latest
LABEL version="1.0.0"
WORKDIR /root
COPY --from=builder /go/src/test/http_server /root
EXPOSE 80
CMD ["./http_server"]
