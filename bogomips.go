// Package bogomips is the Go port of Jeff Tranter standalone bogomips implementation.
package bogomips

import (
	"errors"
	"fmt"
	"time"
)

func Print() {
	fmt.Printf("Calibrating delay loop.. ")
	bogomips, err := Bogomips()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ok - %0.2f BogoMips\n",
			bogomips,
		)
	}
}

func Bogomips() (float64, error) {
	var loops uint64 = 1

	for ; loops > 0; loops <<= 1 {
		start := time.Now()
		delay(loops)
		end := time.Now().Sub(start)

		if end > time.Second {
			var bogomips float64 = float64(loops / uint64(end) * uint64(time.Second) / 500000)
			return bogomips, nil
		}
	}

	return 0, errors.New("failed")
}

func delay(loops uint64) {
	var i uint64
	for i = 0; i < loops; i++ {
	}
}
