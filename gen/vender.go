package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

// ex: VE01, CHINA, OWLER
type Vender struct {
	VenderId uint64 `xml:"VenderId" json:"VenderId"`
	Country  string `xml:"Country" json:"Country"`
	Company  string `xml:"Company" json:"Company"`
}

func vender(r *rand.Rand, num uint64) *Vender {
	return &Vender{
		VenderId: num,
		Country:  GenCountry(r),
		Company:  GenCompany(r),
	}
}

// Get csv headers
func VenderGetHeader() []string {
	return []string{
		"VenderId",
		"Country",
		"Company",
	}
}

// to slice
func (v *Vender) ToSlice() []string {
	return []string{
		fmt.Sprint(v.VenderId),
		v.Country,
		v.Company,
	}
}

func (v *Vender) String() string {
	return fmt.Sprint(
		"Vender-Id: "+strconv.FormatUint(v.VenderId, 10)+"\n",
		"Contry: "+v.Country+"\n",
		"Company: "+v.Company+"\n",
	)
}
