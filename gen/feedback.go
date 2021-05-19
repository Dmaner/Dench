package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

type FeedBack struct {
	productId uint64
	personId  uint64
	star      float64
	comment   string
}

func feedback(r *rand.Rand, pr uint64, pe uint64) *FeedBack {
	return &FeedBack{
		productId: pr,
		personId:  pe,
		star:      RangeFloatGen(r, 0, 10),
		comment:   GenSentence(r),
	}
}

func (f *FeedBack) String() string {
	return fmt.Sprint(
		"ProductId: "+strconv.FormatUint(f.productId, 10)+"\n",
		"PersonId: "+strconv.FormatUint(f.personId, 10)+"\n",
		"Star"+strconv.FormatFloat(f.star, 'f', 1, 32)+"\n",
		"Comment: "+f.comment,
	)
}
