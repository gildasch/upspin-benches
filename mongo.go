package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
