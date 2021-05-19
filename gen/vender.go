package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Vender struct {
	id      uint64
	country string
	company string
}

func vender(r *rand.Rand, num uint64) *Vender {
	return &Vender{
		id:      num,
		country: GenCountry(r),
		company: GenCompany(r),
	}
}

func (v *Vender) String() string {
	return fmt.Sprint(
		"Vender-Id: "+strconv.FormatUint(v.id, 10)+"\n",
		"Contry: "+v.country+"\n",
		"Company: "+v.company+"\n",
	)
}
