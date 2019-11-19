#!/usr/bin/env sh

# ================================================================================================
#  INSTALL GOLANG
# ================================================================================================
export DEBIAN_FRONTEND=noninteractive

apt-get update
apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" install --no-install-recommends \
	g++ \
	gcc \
	libc6-dev \
	make \
	pkg-config

export GOLANG_VERSION=1.13
export GOLANG_INSTALL_URL="https://raw.githubusercontent.com/zeroc0d3/golang-tools-install-script/golang-v${GOLANG_VERSION}/goinstall.sh"
export GOLANG_DEP_URL="https://raw.githubusercontent.com/golang/dep/master/install.sh"
export GOPATH=$HOME/go
export GOROOT=/usr/local/go

if ! [ "${GO_VERSION}" = "" ]
then
  GOLANG_VERSION=${GO_VERSION}
fi

set -eux;
curl ${GOLANG_INSTALL_URL} | bash

##### CUSTOMIZE ~/.profile #####
echo '' >> ~/.profile
echo '### GO-Lang $GOPATH ###
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.profile

# reload source ~/.profile
/bin/bash -c "source ~/.profile"

### Install Dep (Golang Package Manager) ###
curl ${GOLANG_DEP_URL} | bash
