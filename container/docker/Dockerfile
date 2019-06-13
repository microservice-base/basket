FROM golang:alpine as builder
RUN go version

RUN apk update && apk upgrade && \
    apk add --no-cache bash git build-base

WORKDIR $GOPATH/src/basket

COPY  . .

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


# FROM golang:1.12.4
# WORKDIR /root/
# COPY --from=builder /go/src/basket/basket/ .
# COPY --from=builder /go/src/basket/container/docker/ .
# RUN pwd
# RUN ls -al 
# ENTRYPOINT ["sh","/root/app.sh"]

# FROM golang:1.12.4 as builder
# LABEL maintainer="Kerami Ozsoy <mail>"
# WORKDIR $GOPATH/src/basket/
# COPY . $GOPATH/src/basket
# RUN go install -v ./...
# FROM golang:1.12.4-alpine
# COPY --from=builder /go/bin/basket /basket
# CMD ["/basket"]
