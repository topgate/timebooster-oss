###########################################
##             swagger-generator         ##
###########################################
FROM ubuntu:16.04
MAINTAINER @eaglesakura

ENV LANG="ja_JP.UTF-8"

RUN  dpkg --add-architecture i386 \
  && apt-get update \
  && mkdir $HOME/tools \
  && apt-get install -y --fix-missing language-pack-ja-base language-pack-ja \
  && update-locale LANG=ja_JP.UTF-8

# install tools(support 32bit binary)
RUN apt-get install -y --fix-missing libc6:i386 libstdc++6:i386 curl wget zip unzip git-core maven \
 && curl -s "https://get.sdkman.io" | bash \
 && echo y | /bin/bash --login -i -c "sdk install java > /dev/null" \
 && echo y | /bin/bash --login -i -c "sdk install groovy 2.4.9 > /dev/null"

# install go-tools
ENV PATH=/root/tools/go/bin:/root/tools/gopath/bin:/root/tools/idea/bin:$PATH \
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
ENV _JAVA_OPTIONS="-Dfile.encoding=UTF-8" \
    JAVA_HOME=/root/.sdkman/candidates/java/current/ \
    PATH=/root/.sdkman/candidates/java/current:$PATH \
    PATH=/root/.sdkman/candidates/groovy/current/bin:$PATH
# build maven cache
RUN  echo '@GrabResolver(name = "eaglesakura", root = "https://dl.bintray.com/eaglesakura/maven/")' > /root/sync.groovy \
  && echo '@Grab("com.eaglesakura:lightweight-swagger-codegen:1.0.52")' >> /root/sync.groovy \
  && echo 'import java.lang.Object' >> /root/sync.groovy \
  && echo 'println "sync maven"' >> /root/sync.groovy \
  && groovy /root/sync.groovy \
  && rm /root/sync.groovy
