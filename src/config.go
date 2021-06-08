package src

// metadata
const bytesPerWordEstimation = 6
const venderperproduct = 100
const perductperperson = 10
const friendperperson = 4
const personintestpro = 10

// order
const beginyear = 2008
const curyear = 2021

// data save path
// const savepath = "../data"

// represent [start, end]
type RangeInt struct {
	start int
	end   int
}

func rangeint(a, b int) *RangeInt {
	return &RangeInt{
		start: a,
		end:   b,
	}
}

type MetaConfig struct {
	rcustomers *RangeInt
	rproducts  *RangeInt
	rvenders   *RangeInt
	ncinp      int
	npknowp    int
}

type Config struct {
	Meta     *MetaConfig
	DataPath string
}

// Sample
func MetaConfigGen(n int) *MetaConfig {
	var numofc int = n
	var numofp int = numofc * perductperperson
	var numofv int = numofp / venderperproduct
	var numofcinp int = numofc * personintestpro
	var numofpkp int = numofc * friendperperson
	return &MetaConfig{
		rcustomers: rangeint(0, numofc),
		rproducts:  rangeint(0, numofp),
		rvenders:   rangeint(0, numofv),
		ncinp:      numofcinp,
		npknowp:    numofpkp,
	}
}
