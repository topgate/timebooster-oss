FROM ubuntu:16.04
MAINTAINER @eaglesakura

ADD timebooster /usr/local/bin
RUN    apt-get update \
    && apt-get install -y --fix-missing unzip ca-certificates git-core \
    && apt-get autoremove \
    && apt-get clean \
    && chmod +x -f /usr/local/bin/timebooster \
    && echo "timebooster -> `timebooster -h`"
