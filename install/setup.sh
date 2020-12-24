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
  echo "Configuring/installing crypto library prerequirement"
  wget -qO - https://pkgs-ce.cossacklabs.com/stable/centos/cossacklabs.repo | tee /etc/yum.repos.d/cossacklabs.repo
  yum -y install libthemis-devel
  echo "/usr/local/lib" >> /etc/ld.so.conf.d/local.conf
  ldconfig -v > /dev/null
fi
