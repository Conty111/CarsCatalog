PROJECT_PKG = github.com/Conty111/CarsCatalog
BUILD_DIR = build

VERSION ?=$(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)

# remove debug info from the binary & make it smaller
LDFLAGS += -s -w

# inject build info
LDFLAGS += -X ${PROJECT_PKG}/internal/app/build.Version=${VERSION} -X ${PROJECT_PKG}/internal/app/build.CommitHash=${COMMIT_HASH} -X ${PROJECT_PKG}/internal/app/build.BuildDate=${BUILD_DATE}

run-external-API:
	go run ./test/externalAPIserver.go

run:
	go run ./cmd/app/main.go serve

test-unit:
	go test -v -cover ./...

test-all:
	#$(MAKE) start-docker-compose-test
	go test -v ./...
	#${MAKE} stop-docker-compose-test

.PHONY: build
build:
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/app ./cmd/app

gen:
	go generate ./...

deps:
	wire ./...

swagger:
	swag init --parseDependency -g cmd/app/main.go --output=./docs/api/web

install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/onsi/ginkgo/ginkgo