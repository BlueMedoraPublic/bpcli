# bpcli

Lightweight command line utility for interacting with the BindPlane public API

[bindplane.bluemedora.com](https://bindplane.bluemedora.com)
[Bindplane Getting Started](https://docs.bindplane.bluemedora.com/docs/getting-started)
[Bindplane API Documentation](https://docs.bindplane.bluemedora.com/reference#introduction)

## Install
Download the release zip for your platform, unzip, and place the
binary in your path:
```
unzip bpcli_darwin_amd64.zip
chmod +x bpcli
mv bpcli /usr/local/bin
```

## Usage
bpcli uses [cobra](https://github.com/spf13/cobra) for managing
commands and flags.

See `/doc` for usage examples.

#### Help
All commands have a built in help flag. `--help` can be passed at any time.
```
bpcli --help
```

#### Sources
```
bpcli source
bpcli source create
bpcli source delete
bpcli source get
bpcli source list
bpcli source type
bpcli source type get
bpcli source type list
bpcli source type template
```
#### Credentials
```
bpcli credential
bpcli credential create
bpcli credential delete
bpcli credential get
bpcli credential list
bpcli credential type
bpcli credential type get
bpcli credential type template
```
#### Collectors
```
bpcli collector get
bpcli collector list
bpcli collector delete
bpcli collector group get
bpcli collector group list
```
#### Jobs
````
bpcli job list
bpcli job get
````


## Developing

#### Lint

When making changes please write tests and lint your code with `golint`:
```
make lint
```

#### Test

Tests require an api key to be set, even if it is invalid
```
export BINDPLANE_API_KEY='somefakekey'
make test
```

Tests can be run against the live Bindplane api, however, a valid
Bindplane environment and api key must be present
```
export BINDPLANE_API_KEY='somerealkey'
export BINDPLANE_LIVE_TEST=1
make test
```


#### Build

The `Makefile` in this repository will use Docker to build
`bpcli`. This is to ensure a consistent build environment.
Compiling with Docker is optional.

Build with Docker, and check the artifacts directory when finished
- linux zip archive
- darwin zip archive
- windows zip archive
- SHA256 sum file
```
make
```

To build on your own system, without Docker:
```
go get ./...
env CGO_ENABLED=0 go build
```

To cross compile on your own system, without Docker, set
`GOOS`, and `GOARCH`:
```
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get ./...
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## Third party packages

bpcli relies on:
- [cobra](https://github.com/spf13/cobra) for command line flags
