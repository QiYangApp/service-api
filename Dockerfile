FROM golang:1.20-bullseye

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1

RUN apt-get -y install git iproute2 procps lsb-release

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v golang.org/x/tools/gopls@latest  \
    && go install -v github.com/stamblerre/gocode@latest \
    && go install -v golang.org/x/tools/cmd/goimports@latest    \
    && go install -v github.com/ramya-rao-a/go-outline@latest  \
    && go install -v github.com/rogpeppe/godef@latest \
    && go install -v honnef.co/go/tools/cmd/staticcheck@latest


        # Clean up.
RUN  apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*


# Revert workaround at top layer.
ENV DEBIAN_FRONTEND=dialog

# Expose service ports.
EXPOSE 8000
