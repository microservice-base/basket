FROM golang:alpine as builder
RUN go version

RUN apk update && apk upgrade && \
    apk add --no-cache bash git build-base

#
# WORKDIR $GOPATH/src/basket
WORKDIR $GOPATH/src
RUN git clone https://github.com/microservice-base/basket.git 
# COPY  . .
#

WORKDIR $GOPATH/src/basket/basket

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN go build -o app 

FROM alpine:latest
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Europe/Istanbul /etc/localtime
COPY --from=builder /go/src/basket/basket/app /app
ENTRYPOINT ["/app"]
