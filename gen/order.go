package gen

import (
	"math/rand"
	"time"
)

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
	OrderId    uint64   `json:"OrderId"`
	CustomerId uint64   `json:"CustomerId"`
	Cost       float64  `json:"Cost"`
	Orders     []*Order `json:"Suborders"`
	OrdersLen  int      `json:"OrdersLen"`
}

func order(oid uint64, count int, p *Product, f *FeedBack, t time.Time) *Order {
	return &Order{
		SubOrderId:   oid,
		CreationDate: t,
		TotalPrice:   p.price * float64(count),
		Product:      p,
		Feedback:     f,
	}
}

// oid: order id
// pid: customer id
func ctrorders(oid, pid uint64) *CtrOrders {
	return &CtrOrders{
		OrderId:    oid,
		CustomerId: pid,
		Cost:       0,
		Orders:     make([]*Order, 0),
		OrdersLen:  0,
	}
}

func (cos *CtrOrders) Apppend(o *Order) {
	cos.Orders = append(cos.Orders, o)
	cos.OrdersLen++
}

func (cos *CtrOrders) randrecommand(r *rand.Rand) *Order {
	return cos.Orders[r.Intn(cos.OrdersLen)]
}

func (cos *CtrOrders) calcost() {
	for _, o := range cos.Orders {
		cos.Cost += o.TotalPrice
	}
}
