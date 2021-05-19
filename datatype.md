# 数据类型

声明数据生成类型

## 目录

- [数据类型](#数据类型)
	- [目录](#目录)
	- [关系型数据](#关系型数据)
		- [Customer](#customer)
		- [Vender](#vender)
	- [KV型数据](#kv型数据)
		- [Feedback](#feedback)
	- [XML数据](#xml数据)
		- [Invoice](#invoice)
	- [Json类型数据](#json类型数据)
		- [Order](#order)
		- [Product](#product)
	- [图数据](#图数据)
		- [PersonInterestProduct](#personinterestproduct)
		- [PersonKnowPerson](#personknowperson)

## 关系型数据

### Customer

```golang
type Customer struct {
	id          uint64
	fristname   string
	lastname    string
	gender      string
	birthday    time.Time
	locationIP  string
	browserUsed string
	place       string
	job         string
}
```

### Vender

```golang
type Vender struct {
	id      uint64
	country string
	company string
}
```

## KV型数据

### Feedback

```golang
type FeedBack struct {
	productId uint64
	personId  uint64
	star      float64 	`评分`
	comment   string	`用户评价`
}
```

## XML数据

### Invoice 

```xml
<OrderId>
<PersonId>
<OrderDate>
<TotalPrice>
<Orderline>
  <productId>
  <asin>
  <title>
  <price>
  <brand>
</Orderline>
```
## Json类型数据

### Order

```golang
type Order struct {
	id           uint64
	creationdate time.Time
	totalprice   uint64
	orderline    []*Product
}
```

### Product 

```golang
type Product struct {
	id    uint64
	info  string
	price float64
	brand *Vender
}
```

## 图数据

```
Customer -- PersonKnowPerson -- > Customer
```

### PersonInterestProduct

```
Person.id,
Product.id
IntersetValue
```

### PersonKnowPerson

```
from,
to,
creationDate
```