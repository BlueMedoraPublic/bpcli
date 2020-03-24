FROM golang:1.14

ARG bindplane_api_key
ARG bindplane_live_test
ARG version

ENV BINDPLANE_API_KEY=$bindplane_api_key
ENV BINDPLANE_LIVE_TEST=$bindplane_live_test

RUN \
    apt-get update >> /dev/null && \
    apt-get install -y zip

ADD . /bpcli
WORKDIR /bpcli

RUN go test ./...

# Disable CGO to avoid pulling in C dependencies, and compile for
# MACOS, Linux, and Windows
RUN go get github.com/mitchellh/gox
RUN \
    env CGO_ENABLED=0 \
    $GOPATH/bin/gox \
        -arch=amd64 \
        -os='!netbsd !openbsd !freebsd'  \
        -output "artifacts/bpcli_{{.OS}}_{{.Arch}}" \
        ./...

# zip and checksum the output
WORKDIR /bpcli/artifacts
RUN ls | xargs -I{} zip {}.zip {}
# checksum the binaries and zip files, even though we remove
# the binaries at the end
RUN ls | grep 'bpcli' | xargs -n1 sha256sum >> SHA256SUMS
# keep only the zip files
RUN ls | grep -Ev 'zip|SUM' | xargs -n1 rm -f
