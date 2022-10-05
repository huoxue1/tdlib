FROM ubuntu:jammy

ARG DEBIAN_FRONTEND=noninteractive
ARG TARGETARCH
#ARG TZ="Asia/Shanghai"

RUN  mkdir /tdlib/

COPY conf/task.yaml  /tdlib/task.yaml
COPY conf/default_config.yaml  /tdlib/config.yaml

COPY ./dist/docker_linux_$TARGETARCH*/tdlib /tdlib/tdlib

RUN  chmod -R 777 /tdlib/tdlib


VOLUME /tdlib

CMD cd /tdlib && ls && ./tdlib