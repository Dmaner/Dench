package gen

import (
	"math/rand"
	"time"
)

func Year(r *rand.Rand) int {
	return int(RangeIntGen(r, 1980, 2021))
}

func Month(r *rand.Rand) time.Month {
	return time.Month(int(RangeIntGen(r, 1, 12)))
}

func Day(r *rand.Rand) int {
	return int(RangeIntGen(r, 1, 28))
}

func Hour(r *rand.Rand) int {
	return int(RangeIntGen(r, 0, 23))
}

func Min(r *rand.Rand) int {
	return int(RangeIntGen(r, 0, 59))
}

func Sec(r *rand.Rand) int {
	return int(RangeIntGen(r, 0, 59))
}

func Date(r *rand.Rand) time.Time {
	return time.Date(Year(r), Month(r), Day(r), Hour(r), Min(r), Sec(r), 0, time.UTC)
}
