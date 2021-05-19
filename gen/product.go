package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

// param
const minprice = 1
const maxprice = 100000000

type Product struct {
	id    uint64
	info  string
	price float64
	brand *Vender
}

func product(r *rand.Rand, num uint64, v *Vender) *Product {
	return &Product{
		id:    num,
		info:  GenSentence(r),
		price: RangeFloatGen(r, minprice, maxprice),
		brand: v,
	}
}

func (p *Product) String() string {
	return fmt.Sprint(
		"ProductId: "+strconv.FormatUint(p.id, 10)+"\n",
		"Info: "+p.info+"\n",
		"Price: "+strconv.FormatFloat(p.price, 'f', 6, 64)+"\n",
		"Brand: \n"+p.brand.String(),
	)
}
