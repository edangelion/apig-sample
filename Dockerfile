FROM gliderlabs/alpine
MAINTAINER Dan Fujimura <dan.fujimura@gmail.com>

ENV GOPATH /go
COPY . /go/src/github.com/wantedly/
WORKDIR /go/src/github.com/wantedly/apig-sample/

RUN apk add --no-cache --update \
  go \
  git \
  gcc \
  musl-dev \
  && go get ./...

RUN go build -o bin/server

ENTRYPOINT ["bin/server"]
