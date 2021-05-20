package gen

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// params
const bytesPerWordEstimation = 6
const venderperproduct = 10

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
	count := r.Intn(10)
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
func (f *Faker) GenFeedBack(proid uint64, perid uint64) *FeedBack {
	return feedback(f.Rand, proid, perid)
}

// Order
func GenOrder(id uint64, t time.Time, count uint64) *Order {
	return order(id, t, count)
}

// Person interested in product
func (f *Faker) GenCinP(pe, pr uint64) *CinP {
	return cinp(f.Rand, pe, pr)
}

// Person know person
func (f *Faker) GenPKnowP(pf, pt uint64) *PKonwP {
	return pknowp(f.Rand, pf, pt)
}

//////////////////////////////////////////////////////////////////////
//         first step generate product & social network           ////
//////////////////////////////////////////////////////////////////////

func (f *Faker) GenVenders(start int, count int) ([]*Vender, error) {
	ret := make([]*Vender, count)
	for i := start; i < count; i++ {
		ret[i] = f.GenVender(uint64(i))
	}
	fmt.Printf("Generate %d venders successfully\n", count)
	return ret, nil
}

func (f *Faker) GenProducts(start int, count int, vs []*Vender) ([]*Product, error) {
	vlen := len(vs)
	ret := make([]*Product, count)
	for i := start; i < count; i++ {
		randv := vs[f.GenRangeIdx(vlen)]
		ret[i] = f.GenProduct(uint64(i), randv)
	}
	fmt.Printf("Generate %d products successfully\n", count)
	return ret, nil
}

func (f *Faker) GenCustomers(start int, count int) ([]*Customer, error) {
	ret := make([]*Customer, count)
	for i := start; i < count; i++ {
		ret[i] = f.GenCustomer(uint64(i))
	}
	fmt.Printf("Generate %d customers successfully\n", count)
	return ret, nil
}

func (f *Faker) GenCinPs(start int, count int, pers []*Customer, pros []*Product) ([]*CinP, error) {
	ret := make([]*CinP, count)
	set := make(map[string]bool)
	count_pers := len(pers)
	count_pros := len(pros)
	for i := start; i < count; i++ {
		per_idx := f.GenRangeIdx(count_pers)
		pe := pers[per_idx].id
		pro_idx := f.GenRangeIdx(count_pros)
		pr := pros[pro_idx].id
		key := ConnectStringKey(per_idx, pro_idx)
		// check if generate before
		if _, ok := set[key]; !ok {
			set[key] = true
			ret[i] = f.GenCinP(pe, pr)
		} else {
			ret[i] = nil
		}
	}
	fmt.Printf("Generate %d cunstomer insterest products edges successfully\n", count)
	return ret, nil
}

func (f *Faker) GenPKnowPs(start int, count int, pers []*Customer) ([]*PKonwP, error) {
	ret := make([]*PKonwP, count)
	set := make(map[string]bool)
	count_pers := len(pers)
	for i := start; i < count; i++ {
		p1 := f.GenRangeIdx(count_pers)
		p2 := f.GenRangeIdx(count_pers)
		// check if the same customer
		if p1 == p2 {
			continue
		}
		key := ConnectStringKey(p1, p2)
		// check if generate before
		if _, ok := set[key]; !ok {
			set[key] = true
			ret[i] = f.GenPKnowP(pers[p1].id, pers[p2].id)
		} else {
			ret[i] = nil
		}
	}
	fmt.Printf("Generate %d cunstomer knows customer edges successfully\n", count)
	return ret, nil
}

// step 1
// return producrs, venders, customers, cinp, pkonwp
func (f *Faker) StepOne(start int, count int) ([]*Product, []*Vender, []*Customer, []*CinP, []*PKonwP, error) {
	// generate vender & customer
	venders, _ := f.GenVenders(start, count/venderperproduct)
	customers, _ := f.GenCustomers(start, count)

	// generate products
	products, _ := f.GenProducts(start, count, venders)

	// add edge
	cinps, _ := f.GenCinPs(start, count, customers, products)
	pknowsp, _ := f.GenPKnowPs(start, count, customers)

	return products, venders, customers, cinps, pknowsp, nil
}
