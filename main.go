package main

import (
	"Dbench/gen"
)

func main() {
	var f *gen.Faker = gen.New(2)
	var config *gen.MetaConfig = gen.MetaConfigGen(1000)
	f.InitMetaData(config)
}
