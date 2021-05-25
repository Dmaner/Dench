# Dench

D-bench

## Environment

> Ubuntu 20.04
> python 3.8
> go 1.16.4 linux/amd64
> go get github.com/arangodb/go-driver
> go get github.com/chrislusf/glow
> go get github.com/chrislusf/glow/flow
> pip install pyarango

## Usage

### Step 1. 数据生成

- 第一阶段
  - 采用[LDBC-SNB](https://github.com/ldbc/ldbc_snb_datagen/tree/stable)生成图数据和兴趣矩阵
  - 随机生成产品数据
- 第二阶段
  - 随机生成用户和供应商
  - 随机生成兴趣用户和非兴趣用户
  - 遍历兴趣用户调用Purchase(product, person), 按兴趣概率设定购买量
  - 返回社交网络(graph), 供应商和消费者(relation), 订单和产品(json), 发票(XML), 产品评价(KV)
- 第三阶段
  - 对非兴趣用户对周围认识度较高的人用贝叶斯模型计算购买大于阈值的产品量和购买数
  - re-purchase()

### Step 2. 工作负载生成

- 主要采用Unibench的工作负载

## Paper read

- [x] UniBench: A Benchmark for Multi-model Database Management Systems
- [x] The LDBC Social Network Benchmark: Interactive Workload
- [ ] TPC-H

## Plans

- [ ] Arangodb install & go api 
- [x] Unibench paper read
- [ ] gofakeit read & save data
- [ ] mapreduce
- [ ] raft
- [ ] ldbc generator read

## Blogs & github

- [Benchmark for arangodb (Golang version)](https://github.com/arangodb/gobench)
- [Unibench](https://github.com/HY-UDBMS/UniBench)
- [NoSQL Performance Tests](https://github.com/HY-UDBMS/UniBench)
- [Data generator for golang](https://github.com/brianvoe/gofakeit)