FROM golang:1.22.4

RUN set -x; \
  apt-get update -y &&\
  apt-get install -y default-mysql-client
