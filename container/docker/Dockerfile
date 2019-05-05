FROM golang:1.12.4 as builder
RUN go version

COPY . $GOPATH/src/basket/

WORKDIR $GOPATH/src/basket/basket/

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o app *.go

RUN ls -al 
RUN pwd

FROM golang:1.12.4
WORKDIR /root/
COPY --from=builder /go/src/basket/basket/ .
COPY --from=builder /go/src/basket/container/docker/ .

# RUN pwd
# RUN ls -al 

ENTRYPOINT ["sh","/root/app.sh"]



# FROM golang:1.12.4 as builder

# LABEL maintainer="Kerami Ozsoy <mail>"

# WORKDIR $GOPATH/src/basket/

# COPY . $GOPATH/src/basket

# RUN go install -v ./...

# FROM golang:1.12.4-alpine
# COPY --from=builder /go/bin/basket /basket
# CMD ["/basket"]
