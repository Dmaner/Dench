package gen

import (
	"math/rand"
	"time"
)

// Single product order
type Order struct {
	id           uint64
	creationdate time.Time
	totalprice   float64
	product      *Product
	feedback     *FeedBack
}

// all orders of a customer
type CtrOrders struct {
	id        uint64  // order id
	PersonId  uint64  // person id
	cost      float64 // all cost
	orders    []*Order
	ordersLen int
}

func order(oid uint64, count int, p *Product, f *FeedBack, t time.Time) *Order {
	return &Order{
		id:           oid,
		creationdate: t,
		totalprice:   p.price * float64(count),
		product:      p,
		feedback:     f,
	}
}

// oid: order id
// pid: customer id
func ctrorders(oid, pid uint64) *CtrOrders {
	return &CtrOrders{
		id:        oid,
		PersonId:  pid,
		cost:      0,
		orders:    make([]*Order, 0),
		ordersLen: 0,
	}
}

func (cos *CtrOrders) Apppend(o *Order) {
	cos.orders = append(cos.orders, o)
	cos.ordersLen++
}

func (cos *CtrOrders) randrecommand(r *rand.Rand) *Order {
	return cos.orders[r.Intn(cos.ordersLen)]
}

func (cos *CtrOrders) calcost() {
	for _, o := range cos.orders {
		cos.cost += o.totalprice
	}
}
