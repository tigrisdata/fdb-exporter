#!/usr/bin/env bash

OS=$(uname -s)

if [ "${OS}" == "Linux" ]
then
  sudo apt update && sudo apt install -y wget && sudo apt-get purge -y --auto-remove
  wget https://github.com/apple/foundationdb/releases/download/7.1.7/foundationdb-clients_7.1.7-1_amd64.deb
  sudo dpkg -i foundationdb*.deb
fi