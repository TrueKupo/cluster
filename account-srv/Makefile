export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export GOSUMDB=off

BINARY:=account-srv
COMMON_DIR:=$(CURDIR)/../common
LOCAL_BIN:=$(CURDIR)/bin
BUILD_ENVPARMS:=CGO_ENABLED=0

GIT_BRANCH:=$(shell git branch 2> /dev/null | grep '*' | cut -f2 -d' ')
GIT_HASH:=$(shell git log --format="%h" -n 1 2> /dev/null)

GIT_TAG:=$(shell git describe --exact-match --abbrev=0 --tags 2> /dev/null)
GIT_LOG:=$(shell git log --decorate --oneline -n1 2> /dev/null | base64 | tr -d '\n')

APP_VERSION:=$(GIT_BRANCH)/$(GIT_HASH)

LDFLAGS:=-X 'github.com/truekupo/cluster/account-srv/service.Name=account-srv'\
		 -X 'github.com/truekupo/cluster/account-srv/service.Version=$(APP_VERSION)'

export PATH:=$(LOCAL_BIN):$(PATH)

.PHONY: build
build:
	$(info Building...)
	$(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o ./bin/$(BINARY)

proto:
	$(info Generating protobuf... $(COMMON_DIR))
	$(MAKE) -C $(COMMON_DIR)

clean:
	rm -rf bin

test:
	$(info Running tests...)
	go test ./...

docker-image:
	$(info Building docker image...)
	DOCKER_BUILDKIT=1 docker build --no-cache --ssh default --platform=linux/amd64 --tag account-srv -f ./Dockerfile .

docker-run:
	$(info Running docker image...)
	docker run -p 11223 account-srv
