#!/usr/bin/env bash

OS=$(uname -s)

if [ ${OS} == "Linux" ]
then
  apt update && apt install -y wget && apt-get purge -y --auto-remove
  wget https://github.com/apple/foundationdb/releases/download/7.1.7/foundationdb-clients_7.1.7-1_amd64.deb
  dpkg -i foundationdb*.deb
fi
