package main

import (
	"Dbench/gen"
	log "Dbench/util"
)

func main() {
	var f *gen.Faker = gen.New(2)
	var config *gen.MetaConfig = gen.MetaConfigGen(1000)
	f.InitMetaData(config)
	log.WriteLog("1.log", "test")
}
