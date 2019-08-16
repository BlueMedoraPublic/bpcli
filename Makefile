ifeq (, $(shell which docker))
    $(error "No docker in $(PATH)")
endif

ifndef BINDPLANE_API_KEY
    $(error BINDPLANE_API_KEY is not set)
endif

ifndef BINDPLANE_LIVE_TEST
    $(info BINDPLANE_LIVE_TEST is not set, tests will be performed locally only.)
endif

VERSION := $(shell cat bindplane/version.go | grep "const VERSION" | cut -c 17- | tr -d '"')

$(shell mkdir -p artifacts)

build: clean
	$(info building bpcli ${VERSION})

	@docker build \
	    --no-cache \
	    --build-arg bindplane_api_key=${BINDPLANE_API_KEY} \
	    --build-arg bindplane_live_test=${BINDPLANE_LIVE_TEST} \
	    --build-arg version=${VERSION} \
	    -t bpcli:${VERSION} .

	@docker create -ti --name bpcliartifacts bpcli:${VERSION} bash && \
	    docker cp bpcliartifacts:/bpcli/bpcli_linux_amd64.zip artifacts/bpcli_linux_amd64.zip && \
	    docker cp bpcliartifacts:/bpcli/bpcli_darwin_amd64.zip artifacts/bpcli_darwin_amd64.zip && \
	    docker cp bpcliartifacts:/bpcli/bpcli_windows_amd64.zip artifacts/bpcli_windows_amd64.zip && \
	    docker cp bpcliartifacts:/bpcli/SHA256SUMS artifacts/SHA256SUMS

	# cleanup
	@docker rm -fv bpcliartifacts &> /dev/null

test:
	go test ./...

lint:
	golint ./...

fmt:
	go fmt ./...

clean:
	$(shell rm -rf artifacts/*)

quick:
	$(shell env CGO_ENABLED=0 go build)
