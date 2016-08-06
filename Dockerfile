FROM golang:1.6-alpine

ENV SRC_PATH="/go/src/github.com/solher/kit-crud"

ADD https://raw.githubusercontent.com/solher/env2flags/master/env2flags.sh /usr/local/bin/env2flags
RUN chmod u+x /usr/local/bin/env2flags

RUN apk add --update git \
&& mkdir -p $SRC_PATH
COPY . $SRC_PATH
WORKDIR $SRC_PATH

RUN go get -u ./... \
&& go build -o app \
&& cp app /usr/local/bin \
&& apk del git \
&& rm -rf /go/*

WORKDIR /

EXPOSE 8082
ENTRYPOINT ["env2flags", "ZIPKIN_ADDR", "--"]
CMD ["app"]