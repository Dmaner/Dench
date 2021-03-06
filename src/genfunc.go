package src

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
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

func RangeFloatGen(r *rand.Rand, min, max float64) float64 {
	if min >= max {
		return min
	}
	return min + r.Float64()*(max-min)
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

func GenCountry(r *rand.Rand) string {
	return countries[r.Intn(len(countries))]
}

func GenCompany(r *rand.Rand) string {
	return companies[r.Intn(len(companies))]
}

func GenWord(r *rand.Rand) string {
	if BoolGen(r) {
		return noun[r.Intn(len(noun))]
	}
	return verb[r.Intn(len(verb))]
}

func GenSentence(r *rand.Rand) string {
	count := r.Intn(10) + 2
	if count == 0 {
		return ""
	}
	separator := ' '
	sentence := bytes.Buffer{}
	sentence.Grow(count * bytesPerWordEstimation)

	for i := 0; i < count; i++ {
		word := GenWord(r)
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			word = string(runes)
		}
		sentence.WriteString(word)
		if i < count-1 {
			sentence.WriteRune(separator)
		}
	}
	sentence.WriteRune('.')
	return sentence.String()
}

func ConnectStringKey(nums ...int) string {
	var s []string
	for _, num := range nums {
		s = append(s, strconv.Itoa(num))
	}
	return strings.Join(s, "-")
}

// generate random index
func (f *Faker) GenRangeIdx(n int) int {
	return f.Rand.Intn(n)
}

// Customer
func (f *Faker) GenCustomer(id uint64) *Customer {
	return customer(f.Rand, id)
}

// Vender
func (f *Faker) GenVender(id uint64) *Vender {
	return vender(f.Rand, id)
}

// Product
func (f *Faker) GenProduct(id uint64, v *Vender) *Product {
	return product(f.Rand, id, v)
}

// FeedBack
// proid : productId
// perid : customerId
func (f *Faker) GenFeedBack(proid, perid uint64) *FeedBack {
	return feedback(f.Rand, proid, perid)
}

// Order
func (f *Faker) GenOrder(oid, pid uint64, p *Product) *Order {
	count := int(RangeIntGen(f.Rand, 1, 100))
	fb := f.GenFeedBack(p.ProductId, pid)
	t := DateRangeYear(f.Rand, beginyear, curyear)
	return order(oid, count, p, fb, t)
}

// Person interested in product
func (f *Faker) GenCinP(pe, pr uint64) *CinP {
	return cinp(f.Rand, pe, pr)
}

// Person know person
func (f *Faker) GenPKnowP(pf, pt uint64) *PKonwP {
	return pknowp(f.Rand, pf, pt)
}

// set worker
func setworker() {
	log.SetPrefix("[Worker]")
}

// merge worker
func mergeworker() {
	log.SetPrefix("")
}
