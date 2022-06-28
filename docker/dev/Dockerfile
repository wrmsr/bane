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
        apt-transport-https \
        build-essential \
        ca-certificates \
        clang \
        curl \
        dnsutils \
        dumb-init \
        gcc \
        gdb \
        git \
        gnupg \
        htop \
        iputils-ping \
        iputils-tracepath \
        jq \
        less \
        lsb-release \
        lsof \
        make \
        man-db \
        moreutils \
        nano \
        ncdu \
        net-tools \
        netcat \
        pkg-config \
        protobuf-compiler \
        python3 \
        silversearcher-ag \
        socat \
        software-properties-common \
        strace \
        tcpdump \
        telnet \
        tmux \
        unzip \
        vim \
        wget \
        zip \
\
)


## docker

RUN ( \
    mkdir -p /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
    apt-get update && \
    apt-get install docker-ce-cli \
)


# jdk

RUN ( \
    add-apt-repository ppa:openjdk-r/ppa && \
    apt-get update && \
    apt-get install -y \
    \
        openjdk-17-jdk \
    \
)

RUN ( \
    # https://stackoverflow.com/questions/6784463/error-trustanchors-parameter-must-be-non-empty \
    /usr/bin/printf '\xfe\xed\xfe\xed\x00\x00\x00\x02\x00\x00\x00\x00\xe2\x68\x6e\x45\xfb\x43\xdf\xa4\xd9\x92\xdd\x41\xce\xb6\xb2\x1c\x63\x30\xd7\x92' > /etc/ssl/certs/java/cacerts && \
    /var/lib/dpkg/info/ca-certificates-java.postinst configure \
)


## go

RUN ( \
    GO_VERSION="1.18.3" && \
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


## config

COPY \
\
    docker/dev/.gdbinit \
    docker/dev/.tmux.conf \
    docker/dev/.vimrc \
\
    /root/

RUN sed -i '/^\[ -z \"\$PS1\" \]/d' /root/.bashrc


## entrypoint

WORKDIR /bane

ENTRYPOINT ["dumb-init", "--"]
CMD ["sh", "-c", "echo 'Ready' && sleep infinity"]
