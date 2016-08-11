FROM gliderlabs/alpine:3.2

COPY . /go/src/github.com/wantedly/

RUN apk update \
  && apk add go \
  && apk add git \
  && apk add musl-dev

ENV GOPATH /go

WORKDIR /go/src/github.com/wantedly/apig-sample/

RUN go get ./...

RUN go build -o bin/server

ENTRYPOINT ["AUTOMIGRATE=1", "bin/server"]
