FROM alpine:3.6

ADD . /go-etherium
RUN \
  apk add --no-cache git go make gcc musl-dev linux-headers && \
  (cd go-etherium && make geth)                             && \
  cp go-etherium/build/bin/geth /usr/local/bin/             && \
  apk del git go make gcc musl-dev linux-headers            && \
  rm -rf /go-etherium

EXPOSE 8545 30303 30303/udp

ENTRYPOINT ["geth"]
