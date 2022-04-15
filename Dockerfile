FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# RUN sed -E -i -e 's/(archive|ports).ubuntu.com/mirrors.aliyun.com/g' -e '/security.ubuntu.com/d' /etc/apt/sources.list
# apt
RUN apt-get update && \
    apt-get -y install git cmake gcc golang-go libseccomp-dev 

# sandbox
RUN git clone https://github.com/isther/sandbox.git /tmp/sandbox && \
    cd /tmp/sandbox && mkdir build && cd build && cmake .. && make && cp sandbox /bin 

# judger
RUN git clone https://github.com/SXUOJ/old-judger.git /tmp/judger && \
    cd /tmp/judger && git checkout old && go build -o /judger 

# clear
RUN    rm -rf /tmp/sandbox /tmp/judger && \ 
    apt-get purge -y --auto-remove cmake git && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# init
RUN useradd -u 11001 compiler && useradd -u 11002 runner && \
    mkdir /sxu-judger && \
    mkdir /sxu-judger/code /sxu-judger/sample && \
    mkdir /sxu-judger/run && mkdir /sxu-judger/output && \
    chown compiler /sxu-judger/run 

EXPOSE 9000

ENTRYPOINT ["/judger"]
