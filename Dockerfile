FROM alpine:edge
MAINTAINER jspc <james@zero-internet.org.uk>

ADD linux/fizz-buzz-interview /fizz-buzz-interview
ENTRYPOINT ["/fizz-buzz-interview"]
