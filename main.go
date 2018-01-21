package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type TestStruct struct {
	ID   int
	Time time.Time
	Rand int64
}

type Accesser interface {
	persist(ts TestStruct) error
	get(id int) (TestStruct, error)
}

func main() {
	ma, err := NewMongoAccesser()
	if err != nil {
		fmt.Println("could not create mongo accesser:", err)
		return
	}

	ua, err := NewUpspinAccesser(
		os.Args[1],
		os.Args[2],
	)
	if err != nil {
		fmt.Println("could not create upspin accesser:", err)
		return
	}

	accessers := []Accesser{ma, ua}
	dataset := generateDataset()

	for i, a := range accessers {
		if i == 0 {
			fmt.Println("Mongo accesser:")
		} else {
			fmt.Println("Upspin accesser:")
		}
		for _, d := range dataset {
			start := time.Now()
			err := a.persist(d)
			pDuration := time.Since(start)
			fmt.Println("persist took", pDuration)
			if err != nil {
				fmt.Println("received error:", err)
				return
			}

			start = time.Now()
			dd, err := a.get(d.ID)
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
