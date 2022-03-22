FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN sed -E -i -e 's/(archive|ports).ubuntu.com/mirrors.aliyun.com/g' -e '/security.ubuntu.com/d' /etc/apt/sources.list
RUN apt-get update && \
    apt-get -y install git cmake gcc golang-go libseccomp-dev && \
    git clone https://github.com/isther/sandbox.git /tmp/sandbox && \
    git clone https://github.com/isther/judger.git /tmp/judger &&\
    cd /tmp/sandbox && mkdir build && cd build && cmake .. && make && cp sandbox /bin &&\ 
    cd /tmp/judger && go build -o /judger &&\ 
    # mkdir -p /code && \
    # useradd -u 12001 compiler && useradd -u 12002 code && useradd -u 12003 spj && usermod -a -G code spj
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    apt-get purge -y --auto-remove cmake git 

EXPOSE 8080

ENTRYPOINT ["/judger"]