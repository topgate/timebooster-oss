FROM ubuntu:16.04
MAINTAINER @eaglesakura

RUN  dpkg --add-architecture i386 \
  && apt-get update \
  && mkdir $HOME/tools

# install tools
RUN apt-get install -y apt-transport-https curl wget unzip git-core jq

# setup golang
ENV PATH="/root/tools/go/bin:/root/tools/gopath/bin:$PATH" \
    GOROOT="/root/tools/go" \
    GOPATH="/root/tools/gopath" \
    GOBIN="/root/tools/gopath/bin"
RUN  mkdir $HOME/tools/gopath \
  && mkdir $HOME/tools/gopath/bin \
  && wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz -O $HOME/golang.temp.tar.gz \
  && tar xovfz "$HOME/golang.temp.tar.gz" -C "$HOME/tools/" > /dev/null \
  && go get -f -u github.com/eaglesakura/prjdep \
  && rm $HOME/golang.temp.tar.gz
