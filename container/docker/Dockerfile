FROM golang:1.12.4

LABEL maintainer="Kerami Ozsoy <keramiozsoy@hotmail.com>"

WORKDIR $GOPATH/src/basket

COPY ./ ./

RUN go install -v ./...

RUN chmod +x container/docker/app.sh
CMD ["./container/docker/app.sh"]
