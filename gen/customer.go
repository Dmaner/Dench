package gen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
	id          uint64
	fristname   string
	lastname    string
	gender      string
	birthday    time.Time
	locationIP  string
	browserUsed string
	place       string
	job         string
}

func genFirstName(r *rand.Rand) string {
	return firstnames[r.Intn(len(firstnames))]
}

func genLastName(r *rand.Rand) string {
	return lastnames[r.Intn(len(lastnames))]
}

func genGender(r *rand.Rand) string {
	if BoolGen(r) {
		return "Male"
	}
	return "Female"
}

func genIPv4(r *rand.Rand) string {
	num := func() int { return r.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

func genBrowerUsed(r *rand.Rand) string {
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

func genJob(r *rand.Rand) string {
	return jobs[r.Intn(len(jobs))]
}

func customer(r *rand.Rand, num uint64) *Customer {
	return &Customer{
		id:          num,
		fristname:   genFirstName(r),
		lastname:    genLastName(r),
		gender:      genGender(r),
		birthday:    GenDate(r),
		locationIP:  genIPv4(r),
		browserUsed: genBrowerUsed(r),
		place:       GenCountry(r),
		job:         genJob(r),
	}
}

func (c *Customer) String() string {
	return fmt.Sprint(
		"Id : "+strconv.FormatUint(c.id, 10)+"\n",
		"Name: "+c.fristname+" "+c.lastname+"\n",
		"Gender: "+c.gender+"\n",
		"Birthday: "+c.birthday.String()+"\n",
		"IPaddr: "+c.locationIP+"\n",
		"Browser: "+c.browserUsed+"\n",
		"Place: "+c.place+"\n",
		"Job: "+c.job,
	)
}
