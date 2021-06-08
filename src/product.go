package src

import (
	"fmt"
	"math/rand"
	"strconv"
)

// param
const minprice = 1
const maxprice = 100000000

type Product struct {
	ProductId uint64  `xml:"ProductId" json:"ProductId"`
	Info      string  `xml:"Info" json:"Info"`
	Price     float64 `xml:"Price" json:"Price"`
	Vender    *Vender `xml:"Vender" json:"Vender"`
}

func ProductHeaders() []string {
	return []string{
		"ProductId",
		"Info",
		"Price",
		"Vender",
	}
}

func (p *Product) ToSlice() []string {
	return []string{
		strconv.FormatUint(p.ProductId, 10),
		p.Info,
		strconv.FormatFloat(p.Price, 'f', 6, 64),
		fmt.Sprintf("VE%d", p.Vender.VenderId),
	}
}

func product(r *rand.Rand, num uint64, v *Vender) *Product {
	return &Product{
		ProductId: num,
		Info:      GenSentence(r),
		Price:     RangeFloatGen(r, minprice, maxprice),
		Vender:    v,
	}
}

func (p *Product) String() string {
	return fmt.Sprint(
		"ProductId: "+strconv.FormatUint(p.ProductId, 10)+"\n",
		"Info: "+p.Info+"\n",
		"Price: "+strconv.FormatFloat(p.Price, 'f', 6, 64)+"\n",
		"Brand: \n"+p.Vender.String(),
	)
}
