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

# install GAE/Go
RUN  apt-get install -y google-cloud-sdk-app-engine-go \
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

# setup build env
RUN for py in `find /usr/lib/google-cloud-sdk/platform/google_appengine -name "*.py"`; do chmod +x $py; done
ENV PATH="/usr/lib/google-cloud-sdk/platform/google_appengine:/usr/lib/google-cloud-sdk/platform/google_appengine/goroot-1.8/bin:$PATH" \
    GAE_GO_HOME=/usr/lib/google-cloud-sdk/platform/google_appengine \
    GOROOT="/usr/lib/google-cloud-sdk/platform/google_appengine/goroot-1.8"
