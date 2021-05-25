package main

import (
	"Dbench/gen"
	"flag"
	"fmt"
)

var scale = flag.Int("n", 1000, "Generate scale")
var path = flag.String("s", "data", "Data save to path")
var testdatabase = flag.String("d", "arangodb", "Test which database")

func main() {
	var f *gen.Faker = gen.New(1)
	flag.Parse()
	var config = &gen.Config{
		Meta:         gen.MetaConfigGen(*scale),
		DataPath:     *path,
		DataBaseTest: *testdatabase,
	}
	switch config.DataBaseTest {
	case "arangodb":
		f.SequentialGen(config.Meta, config.DataPath)
	default:
		fmt.Printf("No implemention for %s benchmark\n", config.DataBaseTest)
	}
}
