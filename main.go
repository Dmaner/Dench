package main

import (
	"Dbench/gen"
	log "Dbench/util"
)

func main() {
	var f *gen.Faker = gen.New(2)
	// var config *gen.MetaConfig = gen.MetaConfigGen(1000)
	// f.InitMetaData(config)
	vs, err := f.GenVenders(0, 100)
	if err != nil {
		log.ErrorLog(err)
	}
	cs, _ := f.GenCustomers(0, 100)
	ps, _ := f.GenProducts(0, 1000, vs)
	gen.SaveVenders("data", vs)
	gen.SaveCustomers("data", cs)
	gen.SaveProducts("data", ps)
}
