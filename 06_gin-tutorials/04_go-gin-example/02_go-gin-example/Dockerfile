FROM golang:lastest

ENV GOPROXY  https://goproxy.cn,direct
WORKDIR  $GOPATH/code/go-gin-example
COPY .  $GOPATH/code/go-gin-example

RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]