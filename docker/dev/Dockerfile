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
        aria2 \
        bison \
        build-essential \
        ca-certificates \
        clang \
        cmake \
        curl \
        cvs \
        dnsutils \
        dumb-init \
        flex \
        gcc \
        gdb \
        git \
        gnupg \
        graphviz \
        htop \
        iftop \
        iotop \
        iputils-ping \
        iputils-tracepath \
        irssi \
        jq \
        less \
        llvm \
        lsb-release \
        lsof \
        make \
        man-db \
        mercurial \
        moreutils \
        mosh \
        nano \
        ncdu \
        net-tools \
        netcat \
        nodejs \
        npm \
        perl \
        pkg-config \
        postgresql-client \
        protobuf-compiler \
        python3 \
        silversearcher-ag \
        socat \
        software-properties-common \
        strace \
        subversion \
        sudo \
        tcpdump \
        telnet \
        tmux \
        unzip \
        vim \
        wget \
        xsltproc  \
        zip \
        zsh \
\
        build-essential \
        gdb \
        lcov \
        libbz2-dev \
        libffi-dev \
        libgdbm-compat-dev \
        libgdbm-dev \
        liblzma-dev \
        libncurses5-dev \
        libreadline6-dev \
        libsqlite3-dev \
        libssl-dev \
        lzma \
        lzma-dev \
        pkg-config \
        tk-dev \
        uuid-dev \
        zlib1g-dev \
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
    apt-get install -y docker-ce-cli \
)


## jdk

RUN ( \
    add-apt-repository ppa:openjdk-r/ppa && \
    apt-get update && \
    apt-get install -y \
    \
        openjdk-18-jdk \
    \
)

RUN ( \
    # https://stackoverflow.com/questions/6784463/error-trustanchors-parameter-must-be-non-empty \
    /usr/bin/printf '\xfe\xed\xfe\xed\x00\x00\x00\x02\x00\x00\x00\x00\xe2\x68\x6e\x45\xfb\x43\xdf\xa4\xd9\x92\xdd\x41\xce\xb6\xb2\x1c\x63\x30\xd7\x92' > /etc/ssl/certs/java/cacerts && \
    /var/lib/dpkg/info/ca-certificates-java.postinst configure \
)


## versions

COPY .versions /root/


## python

RUN ( \
    PYTHON_VERSION=$(grep '^PYTHON=' /root/.versions | cut -d= -f2) && \
    \
    git clone 'https://github.com/pyenv/pyenv' "$HOME/.pyenv" && \
    "$HOME/.pyenv/bin/pyenv" install -s "$PYTHON_VERSION" \
)


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


## config

COPY \
\
    docker/dev/.gdbinit \
    docker/dev/.tmux.conf \
    docker/dev/.vimrc \
\
    /root/


## entrypoint

WORKDIR /bane

ENTRYPOINT ["dumb-init", "--"]
CMD ["sh", "-c", "echo 'Ready' && sleep infinity"]
