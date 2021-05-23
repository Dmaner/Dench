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
// Single product order
type Order struct {
	SubOrderId   uint64    `json:"SubOrderId"`
	CreationDate time.Time `json:"CreationDate"`
	TotalPrice   float64   `json:"TotalPrice"`
	Product      *Product  `json:"Product"`
	Feedback     *FeedBack `json:"Feedback"`
}

// all orders of a customer
type CtrOrders struct {
	OrderId   uint64   `json:"OrderId"`
	PersonId  uint64   `json:"PersonId"`
	Cost      float64  `json:"Cost"`
	Orders    []*Order `json:"Suborders"`
	OrdersLen int      `json:"OrdersLen"`
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

```golang
// Customer insterest product
type CinP struct {
	PersonId      uint64
	ProductId     uint64
	IntersetValue uint8
}
```

### PersonKnowPerson

```golang
// Person know Person
type PKonwP struct {
	Personfrom   uint64
	Personto     uint64
	Creationdate time.Time
}
```