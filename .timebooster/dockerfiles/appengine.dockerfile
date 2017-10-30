FROM ubuntu:16.04
MAINTAINER @eaglesakura

RUN  dpkg --add-architecture i386 \
  && apt-get update \
  && mkdir $HOME/tools

# install tools
RUN apt-get install -y apt-transport-https curl wget zip unzip git-core gcc jq \
  && echo "deb https://packages.cloud.google.com/apt cloud-sdk-xenial main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list \
  && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
  && apt-get update \
  && apt-get install -y google-cloud-sdk

# setup GAE/Go
ENV PATH="/root/tools/go_appengine:$PATH"
RUN  wget "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.48.zip" -O $HOME/tools/sdk.zip \
  && unzip -d $HOME/tools/ $HOME/tools/sdk.zip > /dev/null \
  && rm $HOME/tools/sdk.zip \
  && wget https://bootstrap.pypa.io/get-pip.py -O $HOME/get-pip.py \
  && python $HOME/get-pip.py \
  && rm $HOME/get-pip.py

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

# install python
ENV PATH="/opt/python/2.7.9/bin:$PATH"
RUN  apt-get install -y make gcc \
  && apt-get install -y libssl-dev zlib1g-dev libreadline6-dev sqlite3 libsqlite3-dev \
  && wget https://www.python.org/ftp/python/2.7.9/Python-2.7.9.tgz -O $HOME/tools/python.zip \
  && tar xovfz $HOME/tools/python.zip -C "$HOME/tools/" > /dev/null \
  && cd $HOME/tools/Python-2.7.9 \
  && ./configure --prefix=/opt/python/2.7.9 > /dev/null \
  && make  > /dev/null \
  && make install \
  && rm $HOME/tools/python.zip \
  && cd /root \
  && rm -rf $HOME/tools/Python-2.7.9

# setup build env
ENV GAE_GO_HOME=/root/tools/go_appengine \
    GOROOT="/root/tools/go_appengine/goroot" \
    GOBIN="/root/tools/go_appengine/bin"
