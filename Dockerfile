FROM alpine:3.12
MAINTAINER Jiajun Liu <liujiajun@rightpaddle.com>

WORKDIR /data/service

ADD thunes_http_server_linux /data/service

CMD ./thunes_http_server_linux
