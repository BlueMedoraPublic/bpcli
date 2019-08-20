# staging environment retrieves dependencies and compiles
#
FROM golang:1.12

WORKDIR /build/src/github.com/BlueMedoraPublic/bpcli

ARG bindplane_api_key
ARG bindplane_live_test
ARG version

ENV BINDPLANE_API_KEY=$bindplane_api_key
ENV BINDPLANE_LIVE_TEST=$bindplane_live_test
ENV GOPATH=/build

RUN \
    apt-get update >> /dev/null && \
    apt-get install -y golint zip

ADD . /bpcli
WORKDIR /bpcli

# Disable CGO to avoid pulling in C dependencies, and compile for
# MACOS, Linux, and Windows
RUN go get github.com/mitchellh/gox
RUN \
    env CGO_ENABLED=0 \
    $GOPATH/bin/gox \
        -arch=amd64 \
        -os='!netbsd !openbsd !freebsd'  \
        ./...

# rename each binary and then zip them
RUN mv bpcli_linux_amd64 bpcli && zip bpcli_linux_amd64.zip bpcli
RUN mv bpcli_darwin_amd64 bpcli && zip bpcli_darwin_amd64.zip bpcli
RUN mv bpcli_windows_amd64.exe bpcli.exe && zip bpcli_windows_amd64.zip bpcli.exe

# build the sha256sum file
RUN ls | grep 'bpcli_' | xargs -n1 sha256sum >> SHA256SUMS
