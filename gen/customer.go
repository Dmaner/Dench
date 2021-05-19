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

func customer(r *rand.Rand, num uint64) *Customer {
	return &Customer{
		id:          num,
		fristname:   GenFirstName(r),
		lastname:    GenLastName(r),
		gender:      GenGender(r),
		birthday:    GenDate(r),
		locationIP:  GenIPv4(r),
		browserUsed: GenBrowerUsed(r),
		place:       GenCountry(r),
		job:         GenJob(r),
	}
}

func (c *Customer) Show() {
	fmt.Println("Id : " + strconv.Itoa(int(c.id)))
	fmt.Printf("Name: %s %s\n", c.fristname, c.lastname)
	fmt.Printf("Gender: %s\n", c.gender)
	fmt.Println("Birthday: " + c.birthday.String())
	fmt.Printf("IPaddr: %s\n", c.locationIP)
	fmt.Printf("Browser: %s\n", c.browserUsed)
	fmt.Printf("Place: %s\n", c.place)
	fmt.Printf("Job: %s\n", c.job)
}
