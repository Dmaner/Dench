#!/bin/sh
# install AgensGraph

git clone https://github.com/bitnine-oss/agensgraph.git

sudo apt-get install build-essential libreadline-dev zlib1g-dev flex bison

cd agensgraph
./configure --prefix=$(pwd)

make install
. ag-env.sh
echo "export PATH=/path/to/agensgraph/bin:\$PATH" >> ~/.bashrc
echo "export LD_LIBRARY_PATH=/path/to/agensgraph/lib:\$LD_LIBRARY_PATH" >> ~/.bashrc
source ~/.bashrc
