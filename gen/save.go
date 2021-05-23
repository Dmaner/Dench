package gen

import (
	log "Dbench/util"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
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
	log.WriteLog(infolog, "Sucessully save to", path)
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
	log.WriteLog(infolog, "Sucessully save to", path)
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
	log.WriteLog(infolog, "Sucessully save to", path)
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
	log.WriteLog(infolog, "Sucessully save to", path)
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
	log.WriteLog(infolog, "Sucessully save to", path)
}

func SaveFeedBacks(savepath string, fbs []*FeedBack) {
	path := savepath + "/feedbacks.csv"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	// write head
	writeCsvSting(csvwriter, FeedBackHeader())

	// write slice
	for _, data := range fbs {
		writeCsvSting(csvwriter, data.ToSlice())
	}
	log.WriteLog(infolog, "Sucessully save to", path)
}

func SaveCtrOrderJson(savepath string, cs []*CtrOrders) {
	path := savepath + "/orders.json"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	js, err := json.Marshal(cs)
	if err != nil {
		log.ErrorLog(err)
	}
	ioutil.WriteFile(path, js, 0644)
	log.WriteLogf(infolog, "Sucessfully save to", path)
}

func SaveCtrOderXML(savepath string, cs []*CtrOrders) {
	path := savepath + "/invoice.xml"
	file, err := os.Create(path)
	if err != nil {
		log.ErrorLog(err)
	}
	defer file.Close()

	js, err := xml.Marshal(cs)
	if err != nil {
		log.ErrorLog(err)
	}
	ioutil.WriteFile(path, js, 0644)
	log.WriteLogf(infolog, "Sucessfully save to", path)
}
