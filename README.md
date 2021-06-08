# Dench

D-bench

## Environment

> Ubuntu 20.04  
> python 3.8  
> docker 20.10.7  
> go 1.16.4 linux/amd64  
> go get github.com/arangodb/go-driver   
> go get github.com/chrislusf/gleam/flow   
> go get github.com/chrislusf/gleam/distributed/gleam   
> pip install pyarango   
> pip install pyorient
> pip install pyagens

Database version

- Arangodb 3.7.11   
> 创建dman用户密码test   
> 创建mydb数据库   

- Mongodb 4.4.6   
> 创建dman用户密码test   
> 创建mydb数据库   

- OrientDB 3.0.4
> 创建dman用户密码test   
> 创建mydb数据库   

- AgensGraph
> 创建dman用户密码test   
> 创建mydb数据库   

## Usage

1. 清除历史数据和日志
> make clean


2. 生成数据
> make 

3. 导入数据
> make import (default arangodb)

4. 测试数据库

> python database/[Arangodb.py|AgensGraph.py|Oriendb.py] [all|Q[1-7]]


## Paper read

- [x] UniBench: A Benchmark for Multi-model Database Management Systems
- [x] The LDBC Social Network Benchmark: Interactive Workload
- [x] TPC-*

## Plans

- [x] Arangodb install & go api 
- [x] Unibench paper read
- [x] gofakeit read & save data
- [x] ldbc generator read
- [x] mapreduce
- [ ] raft

## Blogs & github

- [Benchmark for arangodb (Golang version)](https://github.com/arangodb/gobench)
- [Unibench](https://github.com/HY-UDBMS/UniBench)
- [NoSQL Performance Tests](https://github.com/HY-UDBMS/UniBench)
- [Data generator for golang](https://github.com/brianvoe/gofakeit)