BUILD_PARAMS="-tags=release"

all: build

build:
	go build ${BUILD_PARAMS} .

clean:
	rm -f ./fdb-exporter
