package src

import "log"

//////////////////////////////////////////////////////////////////////
//         first step generate product & social network           ////
//////////////////////////////////////////////////////////////////////

func (f *Faker) GenVenders(start int, end int) ([]*Vender, error) {
	ret := make([]*Vender, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = f.GenVender(uint64(i))
	}
	log.Printf("Generate %d venders successfully", end-start)
	return ret, nil
}

func (f *Faker) GenProducts(start int, end int, vs []*Vender) ([]*Product, error) {
	vlen := len(vs)
	ret := make([]*Product, end-start)
	for i := start; i < end; i++ {
		randv := vs[f.GenRangeIdx(vlen)]
		ret[i-start] = f.GenProduct(uint64(i), randv)
	}
	log.Printf("Generate %d products successfully", end-start)
	return ret, nil
}

func (f *Faker) GenCustomers(start int, end int) ([]*Customer, error) {
	ret := make([]*Customer, end-start)
	for i := start; i < end; i++ {
		ret[i-start] = f.GenCustomer(uint64(i))
	}
	log.Printf("Generate %d customers successfully", end-start)
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
	log.Printf("Generate %d cunstomer insterest products edges successfully", len(ret))
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
	log.Printf("Generate %d cunstomer knows customer edges successfully", len(ret))
	return ret, nil
}

// step 1 : initial metadata
// return producrs, venders, customers, cinp, pkonwp
func (f *Faker) InitMetaData(m *MetaConfig) ([]*Product, []*Vender, []*Customer, []*CinP, []*PKonwP, error) {
	// generate vender & customer
	venders, err := f.GenVenders(m.rvenders.start, m.rcustomers.end)
	if err != nil {
		log.Fatal(err)
	}
	customers, err := f.GenCustomers(m.rcustomers.start, m.rcustomers.end)
	if err != nil {
		log.Fatal(err)
	}
	// generate products
	products, err := f.GenProducts(m.rproducts.start, m.rproducts.end, venders)
	if err != nil {
		log.Fatal(err)
	}

	// add edge
	cinps, err := f.GenCinPs(m.ncinp, customers, products)
	if err != nil {
		log.Fatal(err)
	}
	pknowps, err := f.GenPKnowPs(m.npknowp, customers)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Metadata initial finish")
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
) (map[uint64]*CtrOrders, []*FeedBack, uint64, uint64, error) {
	var singleoId uint64 = 0 // single order id
	var csorderId uint64 = 0 // a customer's orders' id
	customermap := map[uint64]*CtrOrders{}
	feedbacks := make([]*FeedBack, 0, len(cinps))

	// for each interest people
	for _, cinp := range cinps {
		product := products[cinp.ProductId]
		order := f.GenOrder(singleoId, cinp.PersonId, product)
		feedbacks = append(feedbacks, order.Feedback)
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

	return customermap, feedbacks, singleoId, csorderId, nil
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
) (map[uint64]*CtrOrders, []*FeedBack, uint64, uint64, error) {
	feedbacks := make([]*FeedBack, 0, len(pknowps))
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
				feedbacks = append(feedbacks, order.Feedback)
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
	return csmap, feedbacks, sId, csId, nil
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
		log.Fatal(err)
	}
	SaveVenders(path, venders)
	SaveProducts(path, products)
	SaveCustomers(path, customers)
	SaveCinps(path, cinps)
	SavePknowps(path, pknowps)
	csmap, feedbacks1, oldsId, oldcsId, err := f.Purchase(products, customers, cinps)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generate %d order, %d total order\n", oldsId, oldcsId)
	csmap, feedbacks2, sId, csId, err := f.SpreadRepurchase(products, customers, pknowps, csmap, oldsId, oldcsId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Spread %d order, %d total order\n", sId-oldsId, csId-oldcsId)
	csarr := CustomerMapToArr(csmap)
	feedbacks := append(feedbacks1, feedbacks2...)
	SaveFeedBacks(path, feedbacks)
	SaveCtrOrderJson(path, csarr)
	SaveCtrOderXML(path, csarr)
	log.Println("Sequential version run sucessfully")
}

////////////////////////////////////////////////////////////
////          mapreduce version data generator          ////
////////////////////////////////////////////////////////////

func (f *Faker) MapReduceGen(m *MetaConfig, path string) {
	products, venders, customers, cinps, pknowps, err := f.InitMetaData(m)
	if err != nil {
		log.Fatal(err)
	}
	SaveVenders(path, venders)
	SaveProducts(path, products)
	SaveCustomers(path, customers)
	SaveCinps(path, cinps)
	SavePknowps(path, pknowps)

	// split data

}
