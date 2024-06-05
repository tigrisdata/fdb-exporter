BUILD_PARAMS="-tags=release"

.PHONY: all build clean fmt test

all: build

build:
	go build "${BUILD_PARAMS}" .

clean:
	rm -f ./fdb-exporter

deps:
	bash -x scripts/install_deps.sh

fmt:
	find . -name \*.go -not -path bin/ -exec goimports -w {} \;

test: deps
	# Set TEST_JSON_OUTPUT to -json to have a json output from the test
	go test "${TEST_PARAM}" "./..."
