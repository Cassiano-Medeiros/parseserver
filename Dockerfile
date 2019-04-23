FROM golang

ADD ./src /go/src/github.com/cassiano-medeiros/parseserver/src

RUN go get golang.org/x/text
RUN go get -u golang.org/x/tools/...
RUN go get -u golang.org/x/crypto/...

RUN go install -v ./...

ENTRYPOINT /go/src/github.com/cassiano-medeiros/parseserver/src

EXPOSE 8080