FROM gliderlabs/alpine
MAINTAINER Dan Fujimura <dan.fujimura@gmail.com>

ENV GOPATH /go
COPY . /go/src/github.com/wantedly/
WORKDIR /go/src/github.com/wantedly/apig-sample/

RUN apk add --no-cache --virtual=build-deps --update \
  go \
  git \
  gcc \
  musl-dev \
  && go get ./... \
  && go build -o bin/server \
  && cp bin/server /apig-server \
  && cd / \
  && rm -rf /go \
  && apk del build-deps

ENTRYPOINT ["/apig-server"]
