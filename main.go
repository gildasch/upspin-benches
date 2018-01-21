package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TestStruct struct {
	ID   int
	Time time.Time
	Rand int64
}

func main() {
	ma, err := NewMongoAccesser()
	if err != nil {
		fmt.Println("could not create mongo accesser:", err)
		return
	}

	dataset := generateDataset()

	for _, d := range dataset {
		start := time.Now()
		err := ma.persist(d)
		pDuration := time.Since(start)
		fmt.Println("persist took", pDuration)
		if err != nil {
			fmt.Println("received error:", err)
			return
		}

		start = time.Now()
		dd, err := ma.get(d.ID)
		gDuration := time.Since(start)
		fmt.Println("get took", gDuration)
		if err != nil {
			fmt.Println("received error:", err)
			return
		}

		if d != dd {
			fmt.Println("written != read:", d, dd)
		}
	}
}

func generateDataset() []TestStruct {
	rand.Seed(time.Now().Unix())

	var ret []TestStruct
	for i := 0; i < 10; i++ {
		ret = append(ret, TestStruct{
			ID:   i,
			Time: time.Now().Truncate(time.Millisecond),
			Rand: rand.Int63(),
		})
	}

	return ret
}
