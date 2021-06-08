package main

import (
	"Dbench/src"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// var testdatabase = flag.String("d", "arangodb", "Test which database")
var scale = flag.Int("n", 1000, "Generate scale")
var path = flag.String("s", "data", "Data save to path")
var logfile = flag.String("l", "log/", "Log file path")
var version = flag.String("v", "distributed", "distributed|single")

func main() {
	var f *src.Faker = src.New(1)
	flag.Parse()
	var config = &src.Config{
		Meta:     src.MetaConfigGen(*scale),
		DataPath: *path,
	}
	filename := *logfile + time.Now().Format("2006-01-02") + ".log"
	LogFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(LogFile)
	if *version == "distributed" {
		fmt.Println("Using distributed version")
		f.MapReduceGen(config.Meta, config.DataPath)
	} else {
		fmt.Println("Using single version")
		f.SequentialGen(config.Meta, config.DataPath)
	}
	fmt.Println("Test dataset generate finish")
}
