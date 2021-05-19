package gen

import "time"

type Order struct {
	id           uint64
	creationdate time.Time
	totalprice   uint64
	orderline    []*Product
}

func order(oid uint64, t time.Time, pc uint64) *Order {
	return &Order{
		id:           oid,
		creationdate: t,
		totalprice:   0,
		orderline:    make([]*Product, pc),
	}
}

func (o *Order) AddProduct(p *Product) {
	o.orderline = append(o.orderline, p)
}
