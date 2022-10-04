FROM ubuntu:jammy

ARG DEBIAN_FRONTEND=noninteractive
ARG TARGETARCH
#ARG TZ="Asia/Shanghai"

RUN  mkdir /tdlib/

COPY conf/*.yaml  /tdlib/

COPY ./dist/docker_linux_$TARGETARCH*/tdlib /tdlib/tdlib

RUN  chmod -R 777 /tdlib/tdlib && mv /tdlib/default_config.yml

EXPOSE 8080

VOLUME /tdlib

CMD cd /tdlib && ./tdlib