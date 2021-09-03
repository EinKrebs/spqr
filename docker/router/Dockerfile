FROM golang:latest as base


RUN mkdir /router
COPY ./ /router

RUN cd /router && make init && make build

ENTRYPOINT /router/spqr run -s /router/docker/router/init.sql || sleep 2332233
