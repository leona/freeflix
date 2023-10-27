FROM alpine

WORKDIR /app

RUN apk update && apk add go alpine-sdk

ENV GOPATH="/root/go"
ENV PATH="$PATH:$GOPATH/bin"
RUN go install github.com/mitranim/gow@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v golang.org/x/tools/cmd/goimports@latest
RUN go install -v github.com/rogpeppe/godef@latest
RUN go install -v github.com/stamblerre/gocode@latest
RUN mkdir /config && mkdir /data

ADD go.mod /app/
ADD go.sum /app/
RUN go mod download