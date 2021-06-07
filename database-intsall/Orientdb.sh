#!/bin/sh
# install orientdb

# update
sudo apt update && sudo apt upgrade
sudo apt install apt-transport-https default-jdk-headless vim bash-completion

# download
export RELEASE="2.0.4"
wget https://s3.us-east-2.amazonaws.com/orientdb3/releases/${RELEASE}/orientdb-${RELEASE}.tar.gz

# extract
tar xvf orientdb-${RELEASE}.tar.gz
sudo mv orientdb-${RELEASE} /opt/orientdb

# add user
sudo groupadd -r orientdb
sudo useradd --system -g orientdb orientdb
sudo chown -R orientdb:orientdb /opt/orientdb

# create root & active orientdb
sudo /opt/orientdb/bin/server.sh

# create service check this website
# https://www.yiibai.com/orientdb/orientdb_installation.html 第3步 - 将OrientDB服务器配置为服务