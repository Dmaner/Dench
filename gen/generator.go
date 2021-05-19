package gen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// random generator
type Faker struct {
	Rand *rand.Rand
}

func New(seed int64) *Faker {
	return &Faker{Rand: rand.New(rand.NewSource(seed))}
}

// Common

func RangeIntGen(r *rand.Rand, min int64, max int64) int64 {
	if min >= max {
		return min
	}
	return r.Int63n((max+1)-min) + min
}

func BoolGen(r *rand.Rand) bool {
	return RangeIntGen(r, 0, 1) == 1
}

func GenDate(r *rand.Rand) time.Time {
	return Date(r)
}

func GenId(r *rand.Rand, a int64, b int64) string {
	return strconv.Itoa(int(RangeIntGen(r, a, b)))
}

// Customer

func GenFirstName(r *rand.Rand) string {
	return firstnames[r.Intn(len(firstnames))]
}

func GenLastName(r *rand.Rand) string {
	return lastnames[r.Intn(len(lastnames))]
}

func GenGender(r *rand.Rand) string {
	if BoolGen(r) {
		return "Male"
	}
	return "Female"
}

func GenIPv4(r *rand.Rand) string {
	num := func() int { return r.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

func GenBrowerUsed(r *rand.Rand) string {
	num := RangeIntGen(r, 0, 4)
	switch num {
	case 0:
		return "Chrome"
	case 1:
		return "FireFox"
	case 2:
		return "Safari"
	case 3:
		return "Opera"
	default:
		return "Chrome"
	}
}

func GenCountry(r *rand.Rand) string {
	return countries[r.Intn(len(countries))]
}

func GenJob(r *rand.Rand) string {
	return jobs[r.Intn(len(jobs))]
}

func (f *Faker) GenCustomer(num uint64) *Customer {
	return customer(f.Rand, num)
}
