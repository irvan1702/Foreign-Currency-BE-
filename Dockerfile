FROM golang:1.10.1 AS builder
RUN go version

COPY . /go/src/currency-exchange/
WORKDIR /go/src/currency-exchange/
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/currency-exchange/main .

EXPOSE 8081
ENTRYPOINT ["./main"]