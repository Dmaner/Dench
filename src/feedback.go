package src

import (
	"fmt"
	"math/rand"
	"strconv"
)

type FeedBack struct {
	ProductId  uint64  `xml:"ProductId" json:"ProductId"`
	CustomerId uint64  `xml:"CustomerId" json:"CustomerId"`
	Star       float64 `xml:"Star" json:"Star"`
	Comment    string  `xml:"Comment" json:"Comment"`
}

func FeedBackHeader() []string {
	return []string{
		"ProductId",
		"CustomerId",
		"Star",
		"Comment",
	}
}

func (f *FeedBack) ToSlice() []string {
	return []string{
		strconv.FormatUint(f.ProductId, 10),
		strconv.FormatUint(f.CustomerId, 10),
		strconv.FormatFloat(f.Star, 'f', 1, 32),
		f.Comment,
	}
}

func feedback(r *rand.Rand, pr uint64, pe uint64) *FeedBack {
	return &FeedBack{
		ProductId:  pr,
		CustomerId: pe,
		Star:       RangeFloatGen(r, 0, 10),
		Comment:    GenSentence(r),
	}
}

func (f *FeedBack) String() string {
	return fmt.Sprint(
		"ProductId: "+strconv.FormatUint(f.ProductId, 10)+"\n",
		"PersonId: "+strconv.FormatUint(f.CustomerId, 10)+"\n",
		"Star"+strconv.FormatFloat(f.Star, 'f', 1, 32)+"\n",
		"Comment: "+f.Comment,
	)
}
