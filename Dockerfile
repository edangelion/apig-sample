FROM gliderlabs/alpine

COPY . /go/src/github.com/wantedly/

RUN apk update \
  && apk add go \
  && apk add git \
  && apk add gcc \
  && apk add musl-dev

ENV GOPATH /go

WORKDIR /go/src/github.com/wantedly/apig-sample/

RUN go get ./...

RUN go build -o bin/server

#ENTRYPOINT ["AUTOMIGRATE=1", "bin/server"]
