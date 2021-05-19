package main

import (
	"Dbench/gen"
)

func main() {
	var f *gen.Faker = gen.New(2)
	var c *gen.Customer = f.GenCustomer(1)
	c.Show()
}
