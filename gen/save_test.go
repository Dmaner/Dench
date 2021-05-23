package gen

import (
	"fmt"
	"os"
	"testing"
)

// file exist or not
func checkFileExit(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// delete test file
func deltestfile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("file remove Error!")
		fmt.Printf("%s", err)
	}
}

func TestSaveVenders(t *testing.T) {
	var f *Faker = New(1)
	vs, err := f.GenVenders(0, 100)
	if err != nil {
		t.Fatalf("generate failed")
	}
	SaveVenders("data", vs)
	checkpath := "../data/venders.csv"
	checkFileExit(checkpath)
	deltestfile(checkpath)
}

func TestSaveCustomers(t *testing.T) {
	var f *Faker = New(1)
	vs, err := f.GenCustomers(0, 100)
	if err != nil {
		t.Fatalf("generate failed")
	}
	SaveCustomers("data", vs)
	checkpath := "../data/customers.csv"
	checkFileExit(checkpath)
	deltestfile(checkpath)
}

func TestSaveCtrOderJson(t *testing.T) {
	var f *Faker = New(1)
	v := f.GenVender(1)
	p := f.GenProduct(1, v)
	o := f.GenOrder(1, 1, p)
	os := ctrorders(1, p.id)
	os.Apppend(o)
	SaveCtrOrderJson("../data", []*CtrOrders{os})
}
