BUILD_PARAMS="-tags=release"

all: build

build:
	go build ${BUILD_PARAMS} .

clean:
	rm -f ./fdb-exporter

fmt:
	find . -name \*.go -not -path bin/ -exec goimports -w {} \;
