FROM golang:1.14-stretch

RUN set -x; \
  apt-get update -y &&\
  apt-get install -y mysql-client
