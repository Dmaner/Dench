package main

import (
	"Dbench/gen"
	"fmt"
)

func main() {
	var f *gen.Faker = gen.New(2)
	var c *gen.Customer = f.GenCustomer(1)
	var v *gen.Vender = f.GenVender(1)
	var p *gen.Product = f.GenProduct(1, v)
	fmt.Println(c)
	fmt.Println(v)
	fmt.Println(p)
}
