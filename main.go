package main

import (
	"fmt"
	"math/rand"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type MongoAccesser struct {
	*mgo.Collection
}

func NewMongoAccesser() (MongoAccesser, error) {
	session, err := mgo.Dial("localhost:30017")
	if err != nil {
		return MongoAccesser{nil}, err
	}

	coll := session.DB("benches").C("upspin-benches")

	err = coll.EnsureIndex(mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	if err != nil {
		return MongoAccesser{nil}, err
	}

	err = coll.DropCollection()
	if err != nil {
		return MongoAccesser{nil}, err
	}

	return MongoAccesser{coll}, nil
}

func (ma MongoAccesser) persist(ts TestStruct) error {
	return ma.Insert(ts)
}

func (ma MongoAccesser) get(id int) (TestStruct, error) {
	var tss []TestStruct
	err := ma.Find(bson.M{
		"id": id,
	}).All(&tss)
	if err != nil {
		return TestStruct{}, err
	}
	return tss[0], nil
}
