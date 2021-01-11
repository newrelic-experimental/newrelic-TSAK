#!/bin/sh
echo "This is an de-installation for the TSAK: Telemetry Swiss Army Knife"
CWD=`pwd`
cp $CWD/tsak /usr/local/bin
chmod +x /usr/local/bin/tsak
cp $CWD/libclips.so /usr/local/lib
if [ -f /etc/redhat-release ]; then
  yum -y remove zeromq-devel
  yum -y remove sqlite3-devel
  echo "UnConfiguring crypto library prerequirement"
  rm -f /etc/yum.repos.d/cossacklabs.repo
  yum -y remove libthemis-devel
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
  brew untap cossacklabs/tap
  brew uninstall libthemis
  brew uninstall leveldb
  brew uninstall graphviz
  brew uninstall sqlite3
  brew uninstall zeromq
  echo "Requirements installed"
fi
