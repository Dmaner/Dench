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
	VenderId uint64 `xml:"VenderId" json:"VenderId"`
	Country  string `xml:"Country" json:"Country"`
	Company  string `xml:"Company" json:"Company"`
}
```

## KV型数据

### Feedback

```golang
type FeedBack struct {
	ProductId  uint64  `xml:"ProductId" json:"ProductId"`
	CustomerId uint64  `xml:"CustomerId" json:"CustomerId"`
	Star       float64 `xml:"Star" json:"Star"`
	Comment    string  `xml:"Comment" json:"Comment"`
}
```

## XML数据

### Invoice 

```golang
// Single product order
type Order struct {
	SubOrderId   uint64    `xml:"SubOrderId" json:"SubOrderId"`
	CreationDate time.Time `xml:"CreationDate" json:"CreationDate"`
	TotalPrice   float64   `xml:"TotalPrice" json:"TotalPrice"`
	Product      *Product  `xml:"Product" json:"Product"`
	Feedback     *FeedBack `xml:"Feedback" json:"Feedback"`
}

// all orders of a customer
type CtrOrders struct {
	OrderId    uint64   `xml:"OrderId" json:"OrderId"`
	CustomerId uint64   `xml:"CustomerId" json:"CustomerId"`
	Cost       float64  `xml:"Cost" json:"Cost"`
	Orders     []*Order `xml:"Suborders" json:"Suborders"`
	OrdersLen  int      `xml:"OrdersLen" json:"OrdersLen"`
}
```
## Json类型数据

### Order

```golang
// Single product order
type Order struct {
	SubOrderId   uint64    `xml:"SubOrderId" json:"SubOrderId"`
	CreationDate time.Time `xml:"CreationDate" json:"CreationDate"`
	TotalPrice   float64   `xml:"TotalPrice" json:"TotalPrice"`
	Product      *Product  `xml:"Product" json:"Product"`
	Feedback     *FeedBack `xml:"Feedback" json:"Feedback"`
}

// all orders of a customer
type CtrOrders struct {
	OrderId    uint64   `xml:"OrderId" json:"OrderId"`
	CustomerId uint64   `xml:"CustomerId" json:"CustomerId"`
	Cost       float64  `xml:"Cost" json:"Cost"`
	Orders     []*Order `xml:"Suborders" json:"Suborders"`
	OrdersLen  int      `xml:"OrdersLen" json:"OrdersLen"`
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