FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN sed -i 's/ports.ubuntu.com/mirror.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && \
    apt-get -y install git cmake gcc golang-go libseccomp-dev && \
    cd /tmp && git clone https://github.com/isther/sandbox.git && \
    cd sandbox && mkdir build && cd build && cmake .. && make && cp sandbox /bin 

WORKDIR /build
COPY . .
RUN go build -o judger .

RUN apt-get clean && rm -rf /var/lib/apt/lists/* && \
    apt-get purge -y --auto-remove cmake git 
# mkdir -p /code && \
# useradd -u 12001 compiler && useradd -u 12002 code && useradd -u 12003 spj && usermod -a -G code spj

COPY /build/judger /judger

EXPOSE 8080

ENTRYPOINT ["/judger"]