package gen

import (
	"math/rand"
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
