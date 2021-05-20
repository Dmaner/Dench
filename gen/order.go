package gen

import "time"

// Single product order
type Order struct {
	id           uint64
	customerId   uint64
	creationdate time.Time
	totalprice   float64
	product      *Product
	feedback     *FeedBack
}

func order(pid, oid uint64, count int, p *Product, f *FeedBack, t time.Time) *Order {
	return &Order{
		id:           oid,
		customerId:   pid,
		creationdate: t,
		totalprice:   p.price * float64(count),
		product:      p,
		feedback:     f,
	}
}

// func (o *Order) AddProduct(p *Product) {
// 	o.orderline = append(o.orderline, p)
// }
