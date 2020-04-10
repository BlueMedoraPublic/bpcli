ifeq (, $(shell which docker))
    $(error "No docker in $(PATH)")
endif

ifndef BINDPLANE_API_KEY
    $(error BINDPLANE_API_KEY is not set)
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
		docker cp bpcliartifacts:/bpcli/artifacts/. artifacts/

	# cleanup
	@docker rm -fv bpcliartifacts &> /dev/null

test: clean-test quick
	go test --tags=integration ./...
	$(shell cp -f bpcli scripts/test/)
	./scripts/test/test_all.sh

test-local:
	go test ./...

lint:
	golint ./...

fmt:
	go fmt ./...

clean-test:
	$(shell rm -f scripts/test/bpcli)

clean: clean-tests
	$(shell rm -rf artifacts/*)

quick:
	$(shell env CGO_ENABLED=0 go build)
