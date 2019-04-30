FROM golang:1.9

LABEL maintainer="Kerami Ozsoy  <keramiozsoy@hotmail.com>"

WORKDIR $GOPATH/src/basket

COPY ./ ./

WORKDIR $GOPATH/src/basket/basket

# RUN go get -d -v github.com/golang/dep/cmd/dep
# RUN dep ensure

# RUN go get -d -v ./...

RUN go install -v ./...

WORKDIR $GOPATH/src/basket/container/docker
RUN chmod +x app.sh
CMD ["./app.sh"]

# CMD ["go", "run", "basketApplication.go"]