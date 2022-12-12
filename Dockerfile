FROM ubuntu:jammy

ENV GOPATH /go

RUN apt update && apt install -y golang wget && apt-get purge -y --auto-remove \
 && wget https://github.com/apple/foundationdb/releases/download/7.1.7/foundationdb-clients_7.1.7-1_amd64.deb \
 && dpkg -i foundationdb*.deb

RUN mkdir -p /go/src/github.com/tigrisdata/fdb-exporter
WORKDIR /go/src/github.com/tigrisdata/fdb-exporter

COPY go.mod .
COPY go.sum .

RUN GO11MODULE=on go mod download

COPY . .

RUN GO11MODULE=make build

FROM ubuntu:jammy

COPY --from=0 /go/src/github.com/tigrisdata/fdb-exporter /fdb-exporter

RUN apt update && apt install -y wget && apt-get purge -y --auto-remove \
 && wget https://github.com/apple/foundationdb/releases/download/7.1.7/foundationdb-clients_7.1.7-1_amd64.deb \
 && dpkg -i foundationdb*.deb

ENTRYPOINT [ "/fdb-exporter" ]