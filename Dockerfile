FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod tidy \
    go build -o judger .

FROM ubuntu:20.04

COPY --from=builder /build/judger /

ENV DEBIAN_FRONTEND=noninteractive
RUN sed -i 's/ports.ubuntu.com/mirror.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && \
    apt-get -y install git cmake gcc golang-go libseccomp-dev && \
    cd /tmp && git clone https://github.com/isther/sandbox.git && \
    cd sandbox && mkdir build && cd build && cmake .. && make && cp sandbox /bin && \ 
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    apt-get purge -y --auto-remove cmake git 
# mkdir -p /code && \
# useradd -u 12001 compiler && useradd -u 12002 code && useradd -u 12003 spj && usermod -a -G code spj

EXPOSE 8080

ENTRYPOINT ["/judger"]