package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

// ex: VE01, CHINA, OWLER
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

// Get csv headers
func VenderGetHeader() []string {
	return []string{
		"VendID",
		"Country",
		"Company",
	}
}

// to slice
func (v *Vender) ToSlice() []string {
	return []string{
		fmt.Sprintf("VE%d", v.id),
		v.country,
		v.company,
	}
}

func (v *Vender) String() string {
	return fmt.Sprint(
		"Vender-Id: "+strconv.FormatUint(v.id, 10)+"\n",
		"Contry: "+v.country+"\n",
		"Company: "+v.company+"\n",
	)
}
