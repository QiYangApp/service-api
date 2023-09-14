ARG WORKSPACE_GO_VERSION=latest
FROM golang:${WORKSPACE_GO_VERSION}

# Set Environment Variables
ENV DEBIAN_FRONTEND noninteractive

# Start as root
USER root

###########################################################################
# If you're in China, or you need to change sources, will be set CHANGE_SOURCE to true in .env.
###########################################################################

ARG CHANGE_SOURCE=false
RUN if [ ${CHANGE_SOURCE} = true ]; then \
    # Change application source from deb.debian.org to aliyun source
    sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list; \
  fi;

###########################################################################
# Laradock non-root user:
###########################################################################


# always run apt update when start and after add new source list, then clean up at end.
RUN set -xe; \
    apt-get update -y && \
    apt-get install -y \
    libzip-dev \
     zip \
      unzip \
      apt-utils  \
      iproute2 \
      procps \
      lsb-release \
       inetutils-ping \
       net-tools \
       sshpass \
       openssh-server \
        pkg-config \
        libcurl4-openssl-dev \
        libedit-dev \
        libssl-dev \
        libxml2-dev \
        xz-utils \
        libsqlite3-dev \
        rsync \
        sqlite3 \
        git \
        curl \
        vim \
        nano \
        tree \
        cmake \
        gcc \
        make 

ARG WORKSPACE_TZ=UTC
ENV WORKSPACE_TZ ${WORKSPACE_TZ}

RUN ln -snf /usr/share/zoneinfo/$WORKSPACE_TZ /etc/localtime && echo $WORKSPACE_TZ > /etc/timezone

# Add a non-root user to prevent files being created with root permissions on host machine.
ARG PUID=1000
ENV PUID ${PUID}
ARG PGID=1000
ENV PGID ${PGID}

RUN groupadd -g ${PGID} work && \
    useradd -l -u ${PUID} -g work -m work && \
    usermod -p "*" work -s /bin/bash

###########################################################################
# User Aliases
###########################################################################

COPY ./aliases.sh /root/aliases.sh
COPY ./aliases.sh /home/work/aliases.sh

RUN sed -i 's/\r//' /root/aliases.sh && \
    sed -i 's/\r//' /home/work/aliases.sh && \
    chown work:work /home/work/aliases.sh && \
    echo "" >> ~/.bashrc && \
    echo "# Load Custom Aliases" >> ~/.bashrc && \
    echo "source ~/aliases.sh" >> ~/.bashrc && \
	  echo "" >> ~/.bashrc

RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

COPY insecure_id_rsa /tmp/id_rsa
COPY insecure_id_rsa.pub /tmp/id_rsa.pub

RUN rm -f /etc/service/sshd/down && \
    mkdir /root/.ssh && \
    cat /tmp/id_rsa.pub >> /root/.ssh/authorized_keys \
        && cat /tmp/id_rsa.pub >> /root/.ssh/id_rsa.pub \
        && cat /tmp/id_rsa >> /root/.ssh/id_rsa \
        && rm -f /tmp/id_rsa* \
        && chmod 644 /root/.ssh/authorized_keys /root/.ssh/id_rsa.pub \
    && chmod 400 /root/.ssh/id_rsa \
    && cp -rf /root/.ssh /home/work \
    && chown -R work:work /home/work/.ssh \
    && service ssh start \
    && chown work:work /go

USER work

RUN echo "" >> ~/.bashrc && \
    echo "# Load Custom Aliases" >> ~/.bashrc && \
    echo "source ~/aliases.sh" >> ~/.bashrc && \
	echo "" >> ~/.bashrc

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct && \
    echo 'export GO111MODULE = on' >> /home/work/.profile && \
    echo 'export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct' >> /home/work/.profile

#RUN go install github.com/go-delve/delve/cmd/dlv@latest \
#    && go install  golang.org/x/tools/gopls@latest  \
#    && go install  github.com/stamblerre/gocode@latest \
#    && go install  golang.org/x/tools/cmd/goimports@latest    \
#    && go install  github.com/ramya-rao-a/go-outline@latest  \
#    && go install  github.com/rogpeppe/godef@latest \
#    && go install github.com/codegangsta/gin@latest \
#    && go install github.com/stamblerre/gocode@latest

EXPOSE 3000 22

ENV APP_CODE_PATH_CONTAINER=/work/api
WORKDIR ${APP_CODE_PATH_CONTAINER}