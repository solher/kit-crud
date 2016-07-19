FROM golang:1.6-alpine

ENV APP_NAME="kit-crud"
ENV SRC_PATH="/go/src/github.com/solher/kit-crud"

RUN apk add --update git \
&& mkdir -p $SRC_PATH
COPY . $SRC_PATH
WORKDIR $SRC_PATH

RUN go get github.com/Masterminds/glide \
&& glide install \
&& go build -v \
&& cp $APP_NAME /usr/bin \
&& apk del git \
&& rm -rf /go/* \
&& adduser -D app

WORKDIR /

USER app
EXPOSE 3000
CMD $APP_NAME