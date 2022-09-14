FROM ubuntu:22.04
COPY docker/.dockertimestamp /


## deps

RUN ( \
    export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y apt-utils && \
\
    apt-get install -y \
\
        build-essential \
        dumb-init \
        gcc \
        gdb \
        git \
        less \
        make \
        moreutils \
        pkg-config \
        silversearcher-ag \
        socat \
        software-properties-common \
        strace \
        sudo \
        telnet \
        tmux \
        unzip \
        vim \
        wget \
        zip \
\
)


## versions

COPY .versions /root/


## go

RUN ( \
    GO_VERSION=$(grep '^GO=' /root/.versions | cut -d= -f2) && \
    \
    mkdir /build && cd /build && \
    if [ "$(uname -m)" = "aarch64" ] ; then GO_ARCH="arm64" ; else GO_ARCH="amd64" ; fi && \
    GO_FILE="go$GO_VERSION.linux-$GO_ARCH.tar.gz" && \
    wget -q "https://go.dev/dl/$GO_FILE" && \
    tar xzf "$GO_FILE" && \
    chown -R root:root ./go && \
    mv -v go /usr/local && \
    cd / && rm -rf /build \
)

ENV GOPATH=/root/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin


## deps

RUN mkdir /build
COPY go.mod go.sum /build/
RUN ( \
    cd /build && \
    go mod download \
)
RUN rm -rf /build


## entrypoint

ENV BANE_CI=1

WORKDIR /bane

COPY . /bane
