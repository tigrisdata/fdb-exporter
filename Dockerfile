FROM golang:1.22 AS download-env

ENV FDB_DOWNLOAD_URL="https://github.com/apple/foundationdb/releases/download"
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV MULTVERSIONS="6.3.23 7.1.31"
ENV VERSION="7.3.59"

RUN for v in $MULTVERSIONS; do \
        wget -O /tmp/libfdb_c.$v.x86_64.so $FDB_DOWNLOAD_URL/$v/libfdb_c.x86_64.so; \
    done && \
    wget -P /tmp/ ${FDB_DOWNLOAD_URL}/${VERSION}/foundationdb-clients_${VERSION}-1_amd64.deb && \
    dpkg -i /tmp/foundationdb-clients_${VERSION}-1_amd64.deb && \
    rm /tmp/foundationdb-clients_${VERSION}-1_amd64.deb

FROM golang:1.22 AS build-env

COPY --from=download-env /usr/include/foundationdb/fdb* /usr/include/foundationdb/
COPY --from=download-env /usr/lib/libfdb_c.so /lib/

RUN mkdir -p /go/src/github.com/tigrisdata/fdb-exporter
WORKDIR /go/src/github.com/tigrisdata/fdb-exporter

COPY go.mod .
COPY go.sum .

RUN GO11MODULE=on go mod download

COPY . .

RUN go build -tags=release -buildvcs=false -o /tmp/fdb-exporter

FROM debian:trixie-slim
ENV FDB_NETWORK_OPTION_IGNORE_EXTERNAL_CLIENT_FAILURES=""
ENV FDB_NETWORK_OPTION_EXTERNAL_CLIENT_DIRECTORY=/usr/lib/
COPY --from=download-env /tmp/libfdb*.so /usr/lib/
COPY --from=download-env /usr/bin/fdbcli /usr/bin/
COPY --from=download-env /usr/lib/libfdb_c.so /lib/
COPY --from=build-env /tmp/fdb-exporter /usr/local/bin
ENTRYPOINT ["/usr/local/bin/fdb-exporter"]
