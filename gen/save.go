package gen

import (
	log "Dbench/util"
	"encoding/csv"
	"os"
)

func writeCsvSting(writer *csv.Writer, header []string) {
	if err := writer.Write(header); err != nil {
		log.ErrorLog(err)
	}
}

func SaveVenders(savepath string, vs []*Vender) {
	path := savepath + "/venders.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()

	// check if empty
	if len(vs) == 0 {
		log.ErrorLog("Empty venders!")
	}

	// write head
	writeCsvSting(csvwriter, VenderGetHeader())

	// write slice
	for _, data := range vs {
		writeCsvSting(csvwriter, data.ToSlice())
	}
}

func SaveCustomers(savepath string, cs []*Customer) {
	path := savepath + "/customers.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	// write head
	writeCsvSting(csvwriter, CustomerHeaders())

	// write slice
	for _, data := range cs {
		writeCsvSting(csvwriter, data.ToSlice())
	}
}

func SaveProducts(savepath string, ps []*Product) {
	path := savepath + "/products.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	// write head
	writeCsvSting(csvwriter, ProductHeaders())

	// write slice
	for _, data := range ps {
		writeCsvSting(csvwriter, data.ToSlice())
	}
}

func SaveCinps(savepath string, cps []*CinP) {
	path := savepath + "/cinps.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	// write head
	writeCsvSting(csvwriter, CinpHeaders())

	// write slice
	for _, data := range cps {
		writeCsvSting(csvwriter, data.ToSlice())
	}
}

func SavePknowps(savepath string, pps []*PKonwP) {
	path := savepath + "/pknowps.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	// write head
	writeCsvSting(csvwriter, PKnowPHeaders())

	// write slice
	for _, data := range pps {
		writeCsvSting(csvwriter, data.ToSlice())
	}
}
