package gen

import (
	log "Dbench/util"
	"bytes"
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

//////////////////////////////////////////////////////////////////////
//         first step generate product & social network           ////
//////////////////////////////////////////////////////////////////////

func (f *Faker) GenVenders(start int, end int) ([]*Vender, error) {
	ret := make([]*Vender, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = f.GenVender(uint64(i))
	}
	log.WriteLogf(infolog, "Generate %d venders successfully", end-start)
	return ret, nil
}

func (f *Faker) GenProducts(start int, end int, vs []*Vender) ([]*Product, error) {
	vlen := len(vs)
	ret := make([]*Product, end-start)
	for i := start; i < end; i++ {
		randv := vs[f.GenRangeIdx(vlen)]
		ret[i-start] = f.GenProduct(uint64(i), randv)
	}
	log.WriteLogf(infolog, "Generate %d products successfully", end-start)
	return ret, nil
}

func (f *Faker) GenCustomers(start int, end int) ([]*Customer, error) {
	ret := make([]*Customer, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = f.GenCustomer(uint64(i))
	}
	log.WriteLogf(infolog, "Generate %d customers successfully", end-start)
	return ret, nil
}

// warning : will generate nil edge
func (f *Faker) GenCinPs(count int, pers []*Customer, pros []*Product) ([]*CinP, error) {
	ret := make([]*CinP, 0, count)
	set := make(map[string]bool)
	count_pers := len(pers)
	count_pros := len(pros)
	for i := 0; i < count; i++ {
		per_idx := f.GenRangeIdx(count_pers)
		pe := pers[per_idx].id
		pro_idx := f.GenRangeIdx(count_pros)
		pr := pros[pro_idx].ProductId
		key := ConnectStringKey(per_idx, pro_idx)
		// check if generate before
		if _, ok := set[key]; !ok {
			set[key] = true
			ret = append(ret, f.GenCinP(pe, pr))
		}
	}
	log.WriteLogf(infolog, "Generate %d cunstomer insterest products edges successfully", len(ret))
	return ret, nil
}

// warning will generate nil edge
func (f *Faker) GenPKnowPs(count int, pers []*Customer) ([]*PKonwP, error) {
	ret := make([]*PKonwP, 0, count)
	set := make(map[string]bool)
	count_pers := len(pers)
	for i := 0; i < count; i++ {
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
			ret = append(ret, f.GenPKnowP(pers[p1].id, pers[p2].id))
		}
	}
	log.WriteLogf(infolog, "Generate %d cunstomer knows customer edges successfully", len(ret))
	return ret, nil
}

// step 1 : initial metadata
// return producrs, venders, customers, cinp, pkonwp
func (f *Faker) InitMetaData(m *MetaConfig) ([]*Product, []*Vender, []*Customer, []*CinP, []*PKonwP, error) {
	// generate vender & customer
	venders, err := f.GenVenders(m.rvenders.start, m.rcustomers.end)
	if err != nil {
		log.ErrorLog(err)
	}
	customers, err := f.GenCustomers(m.rcustomers.start, m.rcustomers.end)
	if err != nil {
		log.ErrorLog(err)
	}
	// generate products
	products, err := f.GenProducts(m.rproducts.start, m.rproducts.end, venders)
	if err != nil {
		log.ErrorLog(err)
	}

	// add edge
	cinps, err := f.GenCinPs(m.ncinp, customers, products)
	if err != nil {
		log.ErrorLog(err)
	}
	pknowps, err := f.GenPKnowPs(m.npknowp, customers)
	if err != nil {
		log.ErrorLog(err)
	}
	log.WriteLog(infolog, "Metadata initial finish")
	return products, venders, customers, cinps, pknowps, nil
}

//////////////////////////////////////////////////////////////////////////
////         second step   purchase product gengerate order           ////
//////////////////////////////////////////////////////////////////////////

// step 2 purchase
func (f *Faker) Purchase(
	products []*Product,
	customers []*Customer,
	cinps []*CinP,
) (map[uint64]*CtrOrders, uint64, uint64, error) {
	var singleoId uint64 = 0 // single order id
	var csorderId uint64 = 0 // a customer's orders' id
	customermap := map[uint64]*CtrOrders{}

	// for each interest people
	for _, cinp := range cinps {
		product := products[cinp.ProductId]
		order := f.GenOrder(singleoId, cinp.PersonId, product)
		if _, ok := customermap[cinp.PersonId]; !ok {
			csos := ctrorders(csorderId, cinp.PersonId)
			csos.Apppend(order)
			customermap[cinp.PersonId] = csos
			csorderId++
		} else {
			customermap[cinp.PersonId].Apppend(order)
		}
		singleoId++
	}

	return customermap, singleoId, csorderId, nil
}

/////////////////////////////////////////////////////////
////         third step  spread & repurchase         ////
/////////////////////////////////////////////////////////

// Friends recommended to buy good product
// return by or not
func (f *Faker) Expand(fb *FeedBack) bool {
	// TODO: change to poisson distribution
	if fb.Star > 5 {
		return true
	} else {
		return false
	}
}

// step 3
func (f *Faker) SpreadRepurchase(
	products []*Product,
	customers []*Customer,
	pknowps []*PKonwP,
	csmap map[uint64]*CtrOrders,
	sId, csId uint64,
) (map[uint64]*CtrOrders, uint64, uint64, error) {
	// for not interest people
	for _, pp := range pknowps {

		// interest people search not interested people
		if cos, ok := csmap[pp.Personfrom]; ok {
			// choice an product to recommend
			recorder := cos.randrecommand(f.Rand)
			product := recorder.Product

			// if recommend sucessfully
			if f.Expand(recorder.Feedback) {
				order := f.GenOrder(sId, pp.Personto, product)
				if _, ok := csmap[pp.Personto]; !ok {
					csos := ctrorders(csId, pp.Personto)
					csos.Apppend(order)
					csmap[pp.Personto] = csos
					csId++
				} else {
					csmap[pp.Personto].Apppend(order)
				}
				sId++
			}
		}
	}
	return csmap, sId, csId, nil
}

func CustomerMapToArr(csmap map[uint64]*CtrOrders) []*CtrOrders {
	arr := make([]*CtrOrders, 0, len(csmap))
	for _, cs := range csmap {
		cs.calcost()
		arr = append(arr, cs)
	}
	return arr
}

////////////////////////////////////////////////////////////
////         sequential version data generator          ////
////////////////////////////////////////////////////////////

func (f *Faker) SequentialGen(m *MetaConfig, path string) {
	products, venders, customers, cinps, pknowps, err := f.InitMetaData(m)
	if err != nil {
		log.ErrorLog(err)
	}
	SaveVenders(path, venders)
	SaveProducts(path, products)
	SaveCustomers(path, customers)
	SaveCinps(path, cinps)
	SavePknowps(path, pknowps)
	csmap, oldsId, oldcsId, err := f.Purchase(products, customers, cinps)
	if err != nil {
		log.ErrorLog(err)
	}
	log.WriteLogf(infolog, "Generate %d order, %d total order", oldsId, oldcsId)
	csmap, sId, csId, err := f.SpreadRepurchase(products, customers, pknowps, csmap, oldsId, oldcsId)
	if err != nil {
		log.ErrorLog(err)
	}
	log.WriteLogf(infolog, "Spread %d order, %d total order", sId-oldsId, csId-oldcsId)
	csarr := CustomerMapToArr(csmap)
	SaveCtrOrderJson(path, csarr)
	SaveCtrOderXML(path, csarr)
	log.WriteLog(infolog, "Sequential version run sucessfully")
}
