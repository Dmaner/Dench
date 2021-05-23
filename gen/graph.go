package gen

import (
	"math/rand"
	"strconv"
	"time"
)

// Customer insterest product
type CinP struct {
	PersonId      uint64
	ProductId     uint64
	IntersetValue uint8
}

// Person know Person
type PKonwP struct {
	Personfrom   uint64
	Personto     uint64
	Creationdate time.Time
}

func CinpHeaders() []string {
	return []string{
		"CustomerId",
		"ProductId",
		"IntersetValue",
	}
}

func PKnowPHeaders() []string {
	return []string{
		"CustomerId",
		"CustomerId",
		"Creationdate",
	}
}

func (cp *CinP) ToSlice() []string {
	return []string{
		strconv.FormatUint(cp.PersonId, 10),
		strconv.FormatUint(cp.ProductId, 10),
		strconv.FormatUint(uint64(cp.IntersetValue), 10),
	}
}

func (pp *PKonwP) ToSlice() []string {
	return []string{
		strconv.FormatUint(pp.Personfrom, 10),
		strconv.FormatUint(pp.Personto, 10),
		pp.Creationdate.Format("2006-01-02"),
	}
}

func cinp(r *rand.Rand, pe, pr uint64) *CinP {
	return &CinP{
		PersonId:      pe,
		ProductId:     pr,
		IntersetValue: uint8(RangeIntGen(r, 0, 100)),
	}
}

func pknowp(r *rand.Rand, pf, pt uint64) *PKonwP {
	return &PKonwP{
		Personfrom:   pf,
		Personto:     pt,
		Creationdate: GenDate(r),
	}
}
