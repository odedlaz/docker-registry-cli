SOURCEDIR=.
BINARY=docker-registry-cli

LDFLAGS=-ldflags "-X github.com/odedlaz/docker-registry-cli/core.build=`git rev-parse HEAD`"

.DEFAULT_GOAL: all

.PHONY: all
all: build build-alpine64 build-osx64

# options: https://stackoverflow.com/questions/20728767/all-possible-goos-value

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${SOURCEDIR}/bin/${BINARY}-linux64

.PHONY: build-alpine
build-alpine64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -a -installsuffix cgo -o ${SOURCEDIR}/bin/${BINARY}-alpine64

.PHONY: build-osx
build-osx64:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -a -installsuffix cgo -o ${SOURCEDIR}/bin/${BINARY}-osx64

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	rm -rf ${SOURCEDIR}/bin/*
