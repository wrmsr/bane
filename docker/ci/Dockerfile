FROM ubuntu:22.04
COPY docker/.dockertimestamp /


# deps

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


## go

RUN ( \
    GO_VERSION="1.19" && \
    \
    mkdir /build && cd /build && \
    if [ "$(uname -m)" = "aarch64" ] ; then GO_ARCH="arm64" ; else GO_ARCH="amd64" ; fi && \
    GO_FILE="go$GO_VERSION.linux-$GO_ARCH.tar.gz" && \
    wget -q "https://go.dev/dl/$GO_FILE" && \
    tar xzf "$GO_FILE" && \
    chown -R root:root ./go && \
    mv -v go /usr/local && \
    echo "\n\
export GOPATH="'$'"HOME/go\n\
export PATH="'$'"PATH:/usr/local/go/bin:"'$'"GOPATH/bin\n\
" >> /root/.bashrc && \
    cd / && rm -rf /build \
)


# entrypoint

RUN ( \
    touch ~/.bashrc && \
    echo 'export BANE_CI=1' >> ~/.bashrc \
)

WORKDIR /bane

COPY . /bane
