FROM ubuntu:jammy

ARG DEBIAN_FRONTEND=noninteractive
ARG TARGETARCH
#ARG TZ="Asia/Shanghai"

RUN  mkdir /tdlib/

COPY conf/task.yaml  /tdlib/task.yaml

COPY ./dist/docker_linux_$TARGETARCH*/tdlib /tdlib/tdlib

RUN  chmod -R 777 /tdlib/tdlib

EXPOSE 8080

VOLUME /tdlib

CMD cd /tdlib && ./tdlib