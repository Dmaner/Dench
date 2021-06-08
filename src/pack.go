package src

import (
	"hash/fnv"
	"log"
)

type Transaction struct {
	customer *Customer
	product  *Product
	cinp     *CinP
}

type Pack struct {
	transactions []*Transaction
}

func transaction(t *CinP, c *Customer, p *Product) *Transaction {
	return &Transaction{
		customer: c,
		product:  p,
		cinp:     t,
	}
}

func newPack() *Pack {
	return &Pack{
		make([]*Transaction, 0),
	}
}

// add transaction
func (p *Pack) addTransaction(t *Transaction) {
	p.transactions = append(p.transactions, t)
}

func GetBukect(t *CinP, k int) int {
	s := ConnectStringKey(int(t.PersonId), int(t.ProductId))
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32()) % k
}

// split data into k packages
func SplitData(ps []*Product, cs []*Customer, ts []*CinP, k int) []*Pack {
	packs := make([]*Pack, k)
	for _, t := range ts {
		key := GetBukect(t, k)
		trans := transaction(t, cs[t.PersonId], ps[t.ProductId])
		if packs[key] == nil {
			packs[key] = newPack()
		}
		packs[key].addTransaction(trans)
	}
	log.Printf("Split input into %d package\n", k)
	return packs
}
