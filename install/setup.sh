#!/bin/sh
echo "This is an installation for the TSAK: Telemetry Swiss Army Knife"
CWD=`pwd`
cp $CWD/tsak /usr/local/bin
chmod +x /usr/local/bin/tsak
cp $CWD/libclips.so /usr/local/lib
if [ -f /etc/redhat-release ]; then
  echo "Configuring TSAK for RedHat/CentOS"
  echo "Configuring EPEL REPO"
  yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm
  echo "Installing ZeroMQ pre-requirement"
  yum -y install zeromq-devel
  echo "Installing SQLITE"
  yum -y unstall sqlite3-devel
  echo "Configuring/installing crypto library prerequirement"
  wget -qO - https://pkgs-ce.cossacklabs.com/stable/centos/cossacklabs.repo | tee /etc/yum.repos.d/cossacklabs.repo
  yum -y install libthemis-devel
  echo "/usr/local/lib" >> /etc/ld.so.conf.d/local.conf
  ldconfig -v > /dev/null
fi
if [[ `uname -s` == "Darwin" ]]; then
  echo "Configuring TSAK for OSX"
  if [ ! -x /usr/local/bin/brew ]; then
    echo "Installing brew tool"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    echo "Updating brew"
    brew update
  fi
  echo "Installing requirements"
  brew tap cossacklabs/tap
  brew install libthemis
  brew install leveldb
  brew install graphviz
  brew install sqlite3
  brew install zeromq
  echo "Requirements installed"
fi
