package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	_1 := randomNum1(100)
	_2 := randomNum2(100)

	fmt.Println("random number #one : ", _1, "\n"+"random number #two : ", _2)
}

// #1 way
func randomNum1(rangeNum int) int {
	// create initial seed value. default = 1.
	rand.Seed(time.Now().UnixNano())

	// generate random number with selected range (0-n)
	return rand.Intn(rangeNum)
}

// #2 way
func randomNum2(rangeNum int) int {
	// create new source
	// NewSource returns a new pseudo-random Source seeded with the given value.
	s := rand.NewSource(time.Now().UnixNano())
	// create new rand with new source seeded (s)
	r := rand.New(s)
	// generate random number with selected range (0-n)
	i := r.Intn(rangeNum)

	return i
}

// references

// Rand
// func NewSource(i int64) Source  -> https://golang.org/pkg/math/rand/#NewSource
// func New(s source) Rand         -> https://golang.org/pkg/math/rand/#New
// func Intn(i int) int
// func (r *Rand) Seed(s int64)    -> https://golang.org/pkg/math/rand/#Seed

// Time
// func Now() Time                 -> https://golang.org/pkg/time/#Now
// func (t Time) UnixNano() int64  -> https://golang.org/pkg/time/#Time.UnixNano

// Other
// Difficulty with Go Rand package - https://stackoverflow.com/questions/8288679/difficulty-with-go-rand-package
// Go rand.Intn always return same number/value - https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
